package sqls

import "fmt"

//BodyRes SQL 写入人体传感器的数据
func BodyRes(resFlag int) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`insert into bodysensor (status) 
		values (?)`, resFlag)
	if err != nil {
		fmt.Println("人体传感器，写入出错", err.Error())
		return false
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("人体传感器，commit出错", err.Error())
		return false
	}
	return true
}
