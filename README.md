# jptime [![build status](https://travis-ci.org/jackpopper/jptime.svg?branch=master)](https://travis-ci.org/jackpopper/jptime) [![Coverage Status](https://coveralls.io/repos/github/jackpopper/jptime/badge.svg?branch=master)](https://coveralls.io/github/jackpopper/jptime?branch=master) [![GoDoc](https://godoc.org/github.com/jackpopper/jptime?status.svg)](https://godoc.org/github.com/jackpopper/jptime)
go言語の日本の祝日・暦・季節・時刻表示のためのパッケージ

# Example
```Go
package main

import (
       	"fmt"
       	"jptime"
       	"time"
)

func main() {
       	t := jptime.NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
       	isHoliday, holidayName := t.Holiday()
       	if isHoliday {
       		fmt.Println(holidayName) // 元日
       	}
}
```