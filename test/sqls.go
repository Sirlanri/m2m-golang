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
	days := map[int]map[int]int{}
	daymap := make(map[int]int)
	data2 := make([]int, 0, 169)
	flagNext := true
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
				if flagNext {
					data2 = append(data2, times)
				} else {
					data2 = append(data2, 0)
					flagNext = true
				}

				//清除hour相关的标记
				times = 1
				currentHour = hour
			}
		} else {
			if currentDay+1 != day {
				flagNext = false
			}
			days[currentDay] = daymap
			daymap = nil
			times = 1
			currentHour = 0
			currentDay = day
		}
	}
	delete(days, 0)
	data2 = data2[1:]
	convert2(data2)

}

func convert2(res []int) {
	var data [][]int
	index := 0
	for x := 0; x < 7; x++ {
		for y := 0; y < 24; y++ {
			var block []int
			if index >= len(res) {
				block = []int{x, y, 0}
			} else {
				block = []int{x, y, res[index]}
			}

			data = append(data, block)
			index++
		}
	}
	print()
}

func convernJson(res *map[int]map[int]int) {
	//单个数据由单个列表、列表里有3个int
	for x := 0; x < 7; x++ {
		for y := 0; y < 24; y++ {

		}
	}
}

func main() {
	GetAllHour()

}
