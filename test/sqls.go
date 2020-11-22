package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"ri-co.cn/m2m/configs"
)

//Db 创建的唯一指针
var Db *sql.DB

//初始化，自动创建db指针
func init() {
	Db = ConnectDB()
}

//ConnectDB 初始化时，连接数据库
func ConnectDB() *sql.DB {
	database := configs.SQLConfg()
	Db, err := sql.Open("mysql", database)
	if err != nil {
		fmt.Println("数据库初始化链接失败", err.Error())
	}

	if Db.Ping() != nil {
		fmt.Println("初始化-数据库-用户/密码/库验证失败", Db.Ping().Error())
		return nil
	}

	return Db
}

//创建日期的假数据
func GenerateDate() string {
	rand.NewSource(time.Now().Unix())
	hour := strconv.Itoa(rand.Intn(24))
	minute := strconv.Itoa(rand.Intn(60))
	second := strconv.Itoa(rand.Intn(60))

	timedate := strings.Join([]string{hour, minute, second}, ":")
	day1 := rand.Intn(23)
	if day1 == 0 {
		day1++
	}
	day := strconv.Itoa(day1)
	daydate := "2020-11-" + day + " " + timedate
	return daydate
}

func InsertDate() {
	tx, _ := Db.Begin()
	for i := 0; i < 1000; i++ {
		flag := rand.Intn(2)
		date := GenerateDate()
		_, err := tx.Exec(`insert into bodysensor values(?,?)`, flag, date)
		if err != nil {
			fmt.Println("插入出错", err.Error())
		}
	}
	tx.Commit()
}

func GetAllHour() {
	tx, _ := Db.Begin()
	rows1, err := tx.Query(`SELECT day(itime), hour(itime)
		FROM bodysensor WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY itime ORDER BY itime;`)
	if err != nil {
		fmt.Println("查询出错", err.Error())
		return
	}

	var currentDay = 0
	var currentHour = 0
	var times = 1
	days := map[int]interface{}{}
	daymap := make(map[int]int)
	for rows1.Next() {
		if daymap == nil {
			daymap = make(map[int]int)
		}
		var day, hour int
		err := rows1.Scan(&day, &hour)
		if err != nil {
			panic("写入数据出错" + err.Error())
		}
		if currentDay == day {
			if currentHour == hour {
				times++
			} else {
				//hour变化，记录hour的值
				daymap[currentHour] = times
				//清除hour相关的标记
				times = 1
				currentHour = hour
			}
		} else {
			//day变化，清除flag
			days[currentDay] = daymap
			daymap = nil
			times = 1
			currentHour = 0
			currentDay = day
		}
	}

	fmt.Println("")

}

func main() {
	GetAllHour()

}
