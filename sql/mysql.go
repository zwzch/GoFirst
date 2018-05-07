package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	//"fmt"
	"fmt"
)


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func MysqlTest(){
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err := sql.Open("mysql", "root:root@/mysql?charset=utf8")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		err = rows.Scan(&id, &username, &password, &email)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(email)
	}
	}
