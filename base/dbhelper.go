package base

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() *sql.DB {
	//構建連接："用戶名:密碼@tcp(IP:端口)/數據庫?charset=utf8"
	fmt.Println("開始建立連線")
	userName := "root"
	password := "123456"
	ip := "192.168.36.13"
	port := "3106"
	dbName := "BNJ_DB"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println("path :", path)
	// 打開數據庫,前者是驅動名，所以要導入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//設置數據庫最大連接數
	DB.SetConnMaxLifetime(100)
	//設置上數據庫最大閒置連接數
	DB.SetMaxIdleConns(10)
	//驗證連接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
	}
	// defer DB.Close()
	fmt.Println("connnect success")
	return DB
}
