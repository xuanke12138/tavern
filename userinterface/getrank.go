package userinterface

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func getrank(id int) (int ,int){
	db, err := sql.Open("mysql", "root:20020112a@tcp(127.0.0.1:3306)/tavern?charset=utf8")
	defer  db.Close()
	fmt.Println(err)
	rows, _ := db.Query(" select u.alcohol,u.rowno from(select id,alcohol,(@rowNum:=@rowNum+1)as rowno from users,(select (@rowNum := 0))b order by users.alcohol desc) u where u.id=?",id)
	//rows, _ := db.Query("SELECT id FROM users where id =4")

	for rows.Next(){
		var alcohol,rank int
		err=rows.Scan(&alcohol,&rank)
		return alcohol,rank
	}
	return -1,-1

}
