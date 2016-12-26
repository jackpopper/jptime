package jptime

import (
	"fmt"
	"time"
)

func ExampleFmtInt() {
	fmt.Println(FmtInt(2016))
	// Output:
	// 2016
}

func ExampleFmtIntKanji() {
	fmt.Println(FmtIntKanji(2016))
	// Output:
	// 二〇一六
}

func ExampleFmtIntKanjiMeisuu() {
	fmt.Println(FmtIntKanjiMeisuu(99))
	// Output:
	// 九十九
}

func ExampleJpYear_String() {
	t := NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.JpYear().String())
	// Output:
	// 二〇一六年
}

func ExampleJpMonth_String() {
	t := NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.JpMonth().String())
	// Output:
	// 一月
}

func ExampleJpWeekday_String() {
	t := NewJpTime(time.Date(2016, time.December, 31, 0, 0, 0, 0, time.Local))
	fmt.Println(t.JpWeekday().String())
	// Output:
	// 土
}

func ExampleJpDay_String() {
	t := NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.JpDay().String())
	// Output:
	// 一日
}

func ExampleJpHour_String() {
	t := NewJpTime(time.Date(2016, time.January, 1, 1, 0, 0, 0, time.Local))
	fmt.Println(t.JpHour().String())
	// Output:
	// 一時
}

func ExampleJpMinute_String() {
	t := NewJpTime(time.Date(2016, time.January, 1, 1, 1, 0, 0, time.Local))
	fmt.Println(t.JpMinute().String())
	// Output:
	// 一分
}

func ExampleJpSecond_String() {
	t := NewJpTime(time.Date(2016, time.January, 1, 1, 1, 1, 0, time.Local))
	fmt.Println(t.JpSecond().String())
	// Output:
	// 一秒
}

func ExampleJpTime_KyuurekiMonth() {
	t := NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.KyuurekiMonth())
	// Output:
	// 睦月
}

func ExampleJpTime_Wareki() {
	t := NewJpTime(time.Date(0, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Wareki())
	t = NewJpTime(time.Date(1872, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Wareki())
	t = NewJpTime(time.Date(1989, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Wareki())
	t = NewJpTime(time.Date(1989, time.December, 31, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Wareki())
	t = NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Wareki())
	// Output:
	// 紀元前 B.C. 1
	// 西暦 A.D. 1872
	// 昭和 S 64
	// 平成 H 1
	// 平成 H 28
}

func ExampleJpTime_Eto() {
	t := NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Eto())
	// Output:
	// 申
}

func ExampleJpTime_Sekku() {
	t := NewJpTime(time.Date(2016, time.May, 5, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Sekku())
	// Output:
	// true 端午 菖蒲の節句
}

func ExampleJpTime_Sekki24() {
	t := NewJpTime(time.Date(2016, time.March, 20, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Sekki24())
	// Output:
	// true 春分
}

func ExampleJpTime_Holiday() {
	t := NewJpTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(t.Holiday())
	// Output:
	// true 元日
}

func ExampleJpTime_JpFormat() {
	t := NewJpTime(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local))
	fmt.Println(t.JpFormat(ISO8601))
	fmt.Println(t.JpFormat(JISX0301))
	fmt.Println(t.JpFormat(JISX0301JP))
	fmt.Println(t.JpFormat(KanjiDate))
	fmt.Println(t.JpFormat(KanjiTime))
	fmt.Println(t.JpFormat(WarekiDate))
	fmt.Println(t.JpFormat(WarekiKanjiDate))
	fmt.Println(t.JpFormat(JpWeekdayBrackets))
	fmt.Println(t.JpFormat(JpWeekdayString))
	// Output:
	// 2006-01-02T15:04:05+09:00
	// H18.01.02
	// 平18.01.02
	// 二〇〇六年一月二日
	// 十五時四分五秒
	// 平成18年1月2日
	// 平成十八年一月二日
	// （月）
	// 月曜日
}
