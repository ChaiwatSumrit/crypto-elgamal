package main

import (
	"fmt"
	"strconv"
	"time"
)

type datetime struct {
	TIME time.Time `json:"time" valid:"optional`
	TEXT string    `json:"text" valid:"optional`
}

func getServerTime() (datetime, error) {
	var Datetime datetime

	t := time.Now().Format("2006-01-02") //03:04:05
	fmt.Println("TIME input :")
	fmt.Println(t)

	UpdateDate, errUpdateDate := time.Parse(time.RFC3339, t)
	if errUpdateDate != nil {
		// fmt.Println("ERROR")
		// fmt.Println(errUpdateDate)

		return Datetime, errUpdateDate
		// logger.Errorf(methodName+" Failed in parsing Date: %s", errUpdateDate.Error())
	}

	Datetime = datetime{
		TIME: UpdateDate,
		TEXT: "test",
	}
	return Datetime, nil
}

func gettypetime(a string) string { //string ตัวแรก คือประกาศ type ของ a ตัวสองคือ type ของค่าที่ return กลับ
	var yy, mm, dd, aa = 0.0, 0.0, 0.0, 0.0
	aa, err := strconv.ParseFloat(a, 32)
	if err == nil {
		fmt.Println(aa)
	}

	y := time.Now().Format("2006")
	yy, err = strconv.ParseFloat(y, 32)
	if err == nil {
		fmt.Println(yy)
	}

	m := time.Now().Format("01")
	mm, err = strconv.ParseFloat(m, 32)
	if err == nil {
		fmt.Println(mm)
	}

	d := time.Now().Format("02")
	dd, err = strconv.ParseFloat(d, 32)
	if err == nil {
		fmt.Println(dd)
	}
	sum := yy + mm + dd
	suma := sum / aa
	r := fmt.Sprintf("%.4f", suma)

	return r
}
func 
