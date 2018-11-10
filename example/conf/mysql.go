package conf

// mysql配置
//@author Zhiqiang Guo
var UserDb = map[string]string{
	"type":   "mysql",
	"name":   "userdb",
	"writer": "root:123456@tcp(127.0.0.1:3306)/userdb?charset=utf8mb4&parseTime=true",
	"read":  "root:123456@tcp(127.0.0.1:3306)/userdb?charset=utf8mb4&parseTime=true",
}
