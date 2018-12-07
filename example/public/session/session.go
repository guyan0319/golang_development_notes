package session

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

//session存储方式接口
type Provider interface {
	//初始化一个session，sid根据需要生成后传入
	SessionInit(sid string) (Session, error)
	//根据sid,获取session
	SessionRead(sid string) (Session, error)
	//销毁session
	SessionDestroy(sid string) error
	//回收
	SessionGC(maxLifeTime int64)
}

//Session操作接口
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(ket interface{}) error
	SessionID() string
}

type Manager struct {
	cookieName  string
	lock        sync.Mutex //互斥锁
	provider    Provider   //存储session方式
	maxLifeTime int64      //有效期
}

//实例化一个session管理器
func NewSessionManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provide, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q ", provideName)
	}
	return &Manager{cookieName: cookieName, provider: provide, maxLifeTime: maxLifeTime}, nil
}

//注册 由实现Provider接口的结构体调用
func Register(name string, provide Provider) {
	if provide == nil {
		panic("session: Register provide is nil")
	}
	if _, ok := provides[name]; ok {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provide
}

var provides = make(map[string]Provider)

//生成sessionId
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	//加密
	return base64.URLEncoding.EncodeToString(b)
}

//判断当前请求的cookie中是否存在有效的session，存在返回，否则创建
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock() //加锁
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		//创建一个
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid), //转义特殊符号@#￥%+*-等
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxLifeTime),
			Expires:  time.Now().Add(time.Duration(manager.maxLifeTime)),
			//MaxAge和Expires都可以设置cookie持久化时的过期时长，Expires是老式的过期方法，
			// 如果可以，应该使用MaxAge设置过期时间，但有些老版本的浏览器不支持MaxAge。
			// 如果要支持所有浏览器，要么使用Expires，要么同时使用MaxAge和Expires。
		}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value) //反转义特殊符号
		session, _ = manager.provider.SessionRead(sid)
	}
	return session
}

//销毁session 同时删除cookie
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		sid, _ := url.QueryUnescape(cookie.Value)
		manager.provider.SessionDestroy(sid)
		expiration := time.Now()
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Path:     "/",
			HttpOnly: true,
			Expires:  expiration,
			MaxAge:   -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}
