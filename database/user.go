package database

import (
	"fmt"
)

// 用户表结构体
type User struct {
	Id int64 `db:"id"`
	Name string  `db:"name"`
	Age int `db:"age"`
}

// 查询数据，指定字段名
func StructQueryField(index int) *User{

	user := new(User)
	row := MysqlDb.QueryRow("select id, name, age from users where id = ?", index)
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v",err)
		return nil
	}
	fmt.Println(user)
	return user
}

// 查询数据，取所有字段
func StructQueryAllField() *[]User {
	users := make([]User, 0)
	rows, _ := MysqlDb.Query("SELECT * FROM `users` limit ?",100)
	// 遍历
	var user User
	for rows.Next(){
		rows.Scan(&user.Id, &user.Name, &user.Age)
		users = append(users, user)
	}
	fmt.Println(users)
	return &users
}

// 插入数据
func StructInsert() int64 {
	ret, _ := MysqlDb.Exec("insert INTO users(name, age) values(?, ?)","小红", 23)

	//插入数据的主键id
	lastInsertID,_ := ret.LastInsertId()
	fmt.Println("LastInsertID:", lastInsertID)

	//影响行数
	row,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:", row)
	return row
}

// 更新数据
func StructUpdate() int64{
	ret,_ := MysqlDb.Exec("UPDATE users set age=? where id=?","100", 1)
	updNums,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",updNums)
	return updNums
}

// 删除数据
func StructDel() int64{
	ret,_ := MysqlDb.Exec("delete from users where id=?",4)
	delNums,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",delNums)
	return delNums
}


// 事务处理,结合预处理
func StructTx() {
	//事务处理
	tx, _ := MysqlDb.Begin()

	userAddPre, _ := MysqlDb.Prepare("insert into users(name, age) values(?, ?)")
	addRet, _ := userAddPre.Exec("zhaoliu", 15)
	insNums, _ := addRet.RowsAffected()

	userUpdatePre1, _ := tx.Exec("update users set name = 'zhansan'  where name=?", "张三")
	updNums1, _ := userUpdatePre1.RowsAffected()
	userUpdatePre2, _ := tx.Exec("update users set name = 'lisi'  where name=?", "李四")
	updNums2, _ := userUpdatePre2.RowsAffected()

	fmt.Println(insNums)
	fmt.Println(updNums1)
	fmt.Println(updNums2)

	if insNums > 0 && updNums1 > 0 && updNums2 > 0 {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}

// 查询数据，指定字段名,不采用结构体
func RawQueryField() {
	rows, _ := MysqlDb.Query("select id,name from users")
	if rows == nil {
		return
	}
	id := 0
	name := ""
	fmt.Println(rows)
	fmt.Println(rows)
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

// 查询数据,取所有字段,不采用结构体
func RawQueryAllField() {
	rows2, _ := MysqlDb.Query("select * from users")
	cols, _ := rows2.Columns()

	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	fmt.Println(result)
}
