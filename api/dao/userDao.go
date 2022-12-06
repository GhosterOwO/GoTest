package dao

import (
	"fmt"
	"go_server/global"
	"log"
	"reflect"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type userbox struct {
	Items []user
}

func GetUser() (userbox, error) {
	var u user
	var data = userbox{}
	rows, error := global.Mysql.Query("select id,name from BNJ_USER")
	if error != nil {
		return data, nil
	}
	fmt.Println("rows :", rows)
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		u.ID = id
		u.Name = name
		data.Items = append(data.Items, u)
	}
	fmt.Println("end data.Items :", data.Items)
	fmt.Println(reflect.TypeOf(data.Items))
	return data, nil
}
