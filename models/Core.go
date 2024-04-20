package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	//读取配置文件,相对路径  ./conf/app.ini
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	config, err := ini.Load(dir + "/conf/app.ini")
	if err != nil {
		fmt.Println("读取配置文件失败", err)
	}
	//用户名
	username := config.Section("mysql").Key("username").String()
	//密码
	password := config.Section("mysql").Key("password").String()
	//主机
	host := config.Section("mysql").Key("host").String()
	//端口
	port := config.Section("mysql").Key("port").String()
	//数据库
	dbname := config.Section("mysql").Key("db_name").String()

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:root@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//打印sql
		QueryFields: true,
	})

	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

}
