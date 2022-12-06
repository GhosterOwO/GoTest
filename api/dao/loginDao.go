package dao

import (
	"fmt"
	"go_server/global"
)

type login struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Psassword string `json:"password"`
}

func LoginUser(name, password string) (login, error) {
	row := global.Mysql.QueryRow("select * from BNJ_USER where name = ? and password = ? ", name, password)
	var u login
	err := row.Scan(&u.ID, &u.Name, &u.Psassword)
	if err != nil {
		return login{}, err
	}
	return u, nil
}

func AddUser(name, password string) (login, error) {
	// var u user
	result, err := global.Mysql.Exec("insert INTO BNJ_USER(name,password) values(?,?)", name, password)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return login{}, err
	}
	//sql.Result 的LastInsertId()可取得AUTO_INCREMENT的值
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return login{}, err
	}

	fmt.Println("Insert data id:", lastInsertID)

	//RowsAffected() 影響的資料筆數，如果很嚴謹的寫法會判斷RowsAffected()是否與新增的資料筆數一致
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return login{}, err
	}
	fmt.Println("Affected rows:", rowsaffected)
	return login{ID: int(lastInsertID)}, err
}
