package main

import (
	"example/example/public/goinject"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DBEngine struct {
	Key string
}

func (d *DBEngine) Name() string {
	return "DbName"
}

func NewDBEngine() *DBEngine {
	return &DBEngine{}
}

type CacheEngine struct {
	Key string
	T   map[int]string
}

func (c *CacheEngine) Name() string {
	return "CacheEngine"
}

func NewCacheEngine() *CacheEngine {
	return &CacheEngine{}
}

type UserDB struct {
	Db    *DBEngine    `inject:"foo"`
	Db1   *DBEngine    `inject:""`
	Cache *CacheEngine `inject:""`
}

type ItemDB struct {
	DBEngine `inject:"inline"`
	Cache    *CacheEngine `inject:""`
}

type UserService struct {
	Db *UserDB `inject:""`
}

type ItemService struct {
	Db *ItemDB `inject:""`
}

type App struct {
	User *UserService `inject:""`
	Item *ItemService `inject:""`
}

func (a *App) Render() string {
	return fmt.Sprintf(
		"db name is %s ,cache name is %s.",
		a.User.Db.Db.Key,
		a.Item.Db.Cache.Name(),
	)
}
func main() {

	db := NewDBEngine()
	//cache := NewCacheEngine()
	var g goinject.Graph

	var app App
	err := g.Provider(
		&goinject.Object{Value: &app},
		&goinject.Object{Value: db, Name: "foo"},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Ensure(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(app.Render())

}
