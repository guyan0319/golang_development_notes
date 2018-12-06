package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//-------------session_implements-----------------
//Session操作接口，不同存储方式的Sesion操作不同，实现也不同
type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Remove(key interface{}) error
	GetId() string
}

//session实现
type SessionFromMemory struct {
	//唯一标示
	sid              string
	lock             sync.Mutex                  //一把互斥锁
	lastAccessedTime time.Time                   //最后访问时间
	maxAge           int64                       //超时时间
	data             map[interface{}]interface{} //主数据
}

//实例化
func newSessionFromMemory() *SessionFromMemory {
	return &SessionFromMemory{
		data:   make(map[interface{}]interface{}),
		maxAge: 60 * 30, //默认30分钟
	}
}

//同一个会话均可调用，进行设置，改操作必须拥有排斥锁
func (si *SessionFromMemory) Set(key, value interface{}) {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.data[key] = value
}

func (si *SessionFromMemory) Get(key interface{}) interface{} {
	if value := si.data[key]; value != nil {
		return value
	}
	return nil
}
func (si *SessionFromMemory) Remove(key interface{}) error {
	if value := si.data[key]; value != nil {
		delete(si.data, key)
	}
	return nil
}
func (si *SessionFromMemory) GetId() string {
	return si.sid
}

//--------------session_from---------------------------
//session存储方式接口，可以存储在内存，数据库或者文件
//分别实现该接口即可
//如存入数据库的CRUD操作
type Storage interface {
	//初始化一个session，id根据需要生成后传入
	InitSession(sid string, maxAge int64) (Session, error)
	//根据sid，获得当前session
	SetSession(session Session) error
	//销毁session
	DestroySession(sid string) error
	//回收
	GCSession()
}

//session来自内存
type FromMemory struct {
	//由于session包含所有的请求
	//并行时，保证数据独立、一致、安全
	lock     sync.Mutex //互斥锁
	sessions map[string]Session
}

//实例化一个内存实现
func newFromMemory() *FromMemory {
	return &FromMemory{
		sessions: make(map[string]Session, 0),
	}
}

//初始换会话session，这个结构体操作实现Session接口
func (fm *FromMemory) InitSession(sid string, maxAge int64) (Session, error) {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	newSession := newSessionFromMemory()
	newSession.sid = sid
	if maxAge != 0 {
		newSession.maxAge = maxAge
	}
	newSession.lastAccessedTime = time.Now()

	fm.sessions[sid] = newSession //内存管理map
	return newSession, nil
}

//设置
func (fm *FromMemory) SetSession(session Session) error {
	fm.sessions[session.GetId()] = session
	return nil
}

//销毁session
func (fm *FromMemory) DestroySession(sid string) error {
	if _, ok := fm.sessions[sid]; ok {
		delete(fm.sessions, sid)
		return nil
	}
	return nil
}

//监判超时
func (fm *FromMemory) GCSession() {

	sessions := fm.sessions

	//fmt.Println("gc session")

	if len(sessions) < 1 {
		return
	}

	//fmt.Println("current active sessions ", sessions)

	for k, v := range sessions {
		t := (v.(*SessionFromMemory).lastAccessedTime.Unix()) + (v.(*SessionFromMemory).maxAge)

		if t < time.Now().Unix() { //超时了

			fmt.Println("timeout-------->", v)
			delete(fm.sessions, k)
		}
	}

}

//--------------session_manager----------------------
//管理Session,实际操作cookie，Storage
//由于该结构体是整个应用级别的，写、修改都需要枷锁
type SessionManager struct {
	//session数据最终需要在客户端（浏览器）和服务器各存一份
	//客户端时，存放在cookie中
	cookieName string
	//存放方式，如内存，数据库，文件
	storage Storage
	//超时时间
	maxAge int64
	//由于session包含所有的请求
	//并行时，保证数据独立、一致、安全
	lock sync.Mutex
}

//实例化一个session管理器
func NewSessionManager() *SessionManager {
	sessionManager := &SessionManager{
		cookieName: "lzy-cookie",
		storage:    newFromMemory(), //默认以内存实现
		maxAge:     60 * 30,         //默认30分钟
	}
	go sessionManager.GC()

	return sessionManager
}

func (m *SessionManager) GetCookieN() string {
	return m.cookieName
}

//先判断当前请求的cookie中是否存在有效的session,存在返回，不存在创建
func (m *SessionManager) BeginSession(w http.ResponseWriter, r *http.Request) Session {
	//防止处理时，进入另外的请求
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" { //如果当前请求没有改cookie名字对应的cookie
		fmt.Println("-----------> current session not exists")
		//创建一个
		sid := m.randomId()
		//根据保存session方式，如内存，数据库中创建
		session, _ := m.storage.InitSession(sid, m.maxAge) //该方法有自己的锁，多处调用到

		maxAge := m.maxAge

		if maxAge == 0 {
			maxAge = session.(*SessionFromMemory).maxAge
		}
		//用session的ID于cookie关联
		//cookie名字和失效时间由session管理器维护
		cookie := http.Cookie{
			Name: m.cookieName,
			//这里是并发不安全的，但是这个方法已上锁
			Value:    url.QueryEscape(sid), //转义特殊符号@#￥%+*-等
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(maxAge),
			Expires:  time.Now().Add(time.Duration(maxAge)),
		}
		http.SetCookie(w, &cookie) //设置到响应中
		return session
	} else { //如果存在

		sid, _ := url.QueryUnescape(cookie.Value)        //反转义特殊符号
		session := m.storage.(*FromMemory).sessions[sid] //从保存session介质中获取
		fmt.Println("session --------->", session)
		if session == nil {
			fmt.Println("-----------> current session is nil")
			//创建一个
			//sid := m.randomId()
			//根据保存session方式，如内存，数据库中创建
			newSession, _ := m.storage.InitSession(sid, m.maxAge) //该方法有自己的锁，多处调用到

			maxAge := m.maxAge

			if maxAge == 0 {
				maxAge = newSession.(*SessionFromMemory).maxAge
			}
			//用session的ID于cookie关联
			//cookie名字和失效时间由session管理器维护
			newCookie := http.Cookie{
				Name: m.cookieName,
				//这里是并发不安全的，但是这个方法已上锁
				Value:    url.QueryEscape(sid), //转义特殊符号@#￥%+*-等
				Path:     "/",
				HttpOnly: true,
				MaxAge:   int(maxAge),
				Expires:  time.Now().Add(time.Duration(maxAge)),
			}
			http.SetCookie(w, &newCookie) //设置到响应中
			return newSession
		}
		fmt.Println("-----------> current session exists")
		return session
	}

}

//更新超时
func (m *SessionManager) Update(w http.ResponseWriter, r *http.Request) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil {
		return
	}
	t := time.Now()
	sid, _ := url.QueryUnescape(cookie.Value)

	sessions := m.storage.(*FromMemory).sessions
	session := sessions[sid].(*SessionFromMemory)
	session.lastAccessedTime = t
	sessions[sid] = session

	if m.maxAge != 0 {
		cookie.MaxAge = int(m.maxAge)
	} else {
		cookie.MaxAge = int(session.maxAge)
	}
	http.SetCookie(w, cookie)
}

//通过ID获取session
func (m *SessionManager) GetSessionById(sid string) Session {
	session := m.storage.(*FromMemory).sessions[sid]
	return session
}

//是否内存中存在
func (m *SessionManager) MemoryIsExists(sid string) bool {
	_, ok := m.storage.(*FromMemory).sessions[sid]
	if ok {
		return true
	}
	return false
}

//手动销毁session，同时删除cookie
func (m *SessionManager) Destroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()

		sid, _ := url.QueryUnescape(cookie.Value)
		m.storage.DestroySession(sid)

		cookie2 := http.Cookie{
			MaxAge:  0,
			Name:    m.cookieName,
			Value:   "",
			Path:    "/",
			Expires: time.Now().Add(time.Duration(0)),
		}

		http.SetCookie(w, &cookie2)
	}
}

func (m *SessionManager) CookieIsExists(r *http.Request) bool {
	_, err := r.Cookie(m.cookieName)
	if err != nil {
		return false
	}
	return true
}

//开启每个会话，同时定时调用该方法
//到达session最大生命时，且超时时。回收它
func (m *SessionManager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.storage.GCSession()
	//在多长时间后执行匿名函数，这里指在某个时间后执行GC
	time.AfterFunc(time.Duration(m.maxAge*10), func() {
		m.GC()
	})
}

//是否将session放入内存（操作内存）默认是操作内存
func (m *SessionManager) IsFromMemory() {
	m.storage = newFromMemory()
}

//是否将session放入数据库（操作数据库）
func (m *SessionManager) IsFromDB() {
	//TODO
	//关于存数据库暂未实现
}

func (m *SessionManager) SetMaxAge(t int64) {
	m.maxAge = t
}

//如果你自己实现保存session的方式，可以调该函数进行定义
func (m *SessionManager) SetSessionFrom(storage Storage) {
	m.storage = storage
}

//生成一定长度的随机数
func (m *SessionManager) randomId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	//加密
	return base64.URLEncoding.EncodeToString(b)
}
