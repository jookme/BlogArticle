package Database

import (
	"database/sql"
	"log"

	//导入mysql驱动包
	_ "github.com/go-sql-driver/mysql"
)

//Init 数据库池初始化
func Init() *sql.DB {
	//初始化一个sql.DB对象,接口为mysql中的blog数据库，账号为root，密码为1996
	SqlDb, err := sql.Open("mysql", "root:1996@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer SqlDb.Close()

	//测试与数据库的连接是否可用
	err = SqlDb.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return SqlDb
}
