package jptime

import (
	"fmt"
	"testing"
	"time"
)

type JpTimeFmtTest struct {
	num int
	str string
}

var fmtinttests = []JpTimeFmtTest{
	{1, "1"},
	{10, "10"},
	{2000, "2000"},
	{2016, "2016"},
}

func TestFmtInt(t *testing.T) {
	for _, test := range fmtinttests {
		newI := FmtInt(test.num)
		if newI != test.str {
			t.Errorf("FmtInt %v = %v", test.str, newI)
		}
	}
}

var fmtintkanjitests = []JpTimeFmtTest{
	{1, "一"},
	{10, "一〇"},
	{2000, "二〇〇〇"},
	{2016, "二〇一六"},
}

func TestFmtIntKanji(t *testing.T) {
	for _, test := range fmtintkanjitests {
		newK := FmtIntKanji(test.num)
		if newK != test.str {
			t.Errorf("FmtIntKanji %v = %v", test.str, newK)
		}
	}
}

var fmtintkanjimeisuutests = []JpTimeFmtTest{
	{0, "零"},
	{1, "一"},
	{10, "十"},
	{99, "九十九"},
	{99999, "九万九千九百九十九"},
	{100000, ""},
}

func TestFmtIntKanjiMeisuu(t *testing.T) {
	for _, test := range fmtintkanjimeisuutests {
		newK := FmtIntKanjiMeisuu(test.num)
		if newK != test.str {
			t.Errorf("FmtIntKanjiMeisuu %v = %v", test.str, newK)
		}
	}
}

type JpTimeTest struct {
	time time.Time
	str  string
}

var jpyeartestts = []JpTimeTest{
	{time.Date(0, time.January, 1, 0, 0, 0, 0, time.Local), "紀元前一年"},
	{time.Date(12, time.January, 1, 0, 0, 0, 0, time.Local), "一二年"},
	{time.Date(123, time.January, 1, 0, 0, 0, 0, time.Local), "一二三年"},
	{time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local), "二〇〇〇年"},
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "二〇一六年"},
}

func TestJpYear_String(t *testing.T) {
	for _, test := range jpyeartestts {
		jpt := NewJpTime(test.time)
		newJpYear := jpt.JpYear().String()
		if newJpYear != test.str {
			t.Errorf("JpTime_JpYear_String %v = %v", test.str, newJpYear)
		}
	}
}

var jpmonthtests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "一月"},
	{time.Date(2016, time.October, 1, 0, 0, 0, 0, time.Local), "十月"},
	{time.Date(2016, time.December, 1, 0, 0, 0, 0, time.Local), "十二月"},
}

func TestJpMonth_String(t *testing.T) {
	for _, test := range jpmonthtests {
		jpt := NewJpTime(test.time)
		newJpMonth := jpt.JpMonth().String()
		if newJpMonth != test.str {
			t.Errorf("JpTime_JpMonth_String %v = %v", test.str, newJpMonth)
		}
	}
}

var jpweekdaytests = []JpTimeTest{
	{time.Date(2016, time.December, 25, 0, 0, 0, 0, time.Local), "日"},
	{time.Date(2016, time.December, 26, 0, 0, 0, 0, time.Local), "月"},
	{time.Date(2016, time.December, 27, 0, 0, 0, 0, time.Local), "火"},
	{time.Date(2016, time.December, 28, 0, 0, 0, 0, time.Local), "水"},
	{time.Date(2016, time.December, 29, 0, 0, 0, 0, time.Local), "木"},
	{time.Date(2016, time.December, 30, 0, 0, 0, 0, time.Local), "金"},
	{time.Date(2016, time.December, 31, 0, 0, 0, 0, time.Local), "土"},
}

func TestJpWeekday_String(t *testing.T) {
	for _, test := range jpweekdaytests {
		jpt := NewJpTime(test.time)
		newJpWeekday := jpt.JpWeekday().String()
		if newJpWeekday != test.str {
			t.Errorf("JpTime_JpWeekday_String %v = %v", test.str, newJpWeekday)
		}
	}
}

var jpdaytests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "一日"},
	{time.Date(2016, time.January, 10, 0, 0, 0, 0, time.Local), "十日"},
	{time.Date(2016, time.January, 11, 0, 0, 0, 0, time.Local), "十一日"},
}

func TestJpDay_String(t *testing.T) {
	for _, test := range jpdaytests {
		jpt := NewJpTime(test.time)
		newJpDay := jpt.JpDay().String()
		if newJpDay != test.str {
			t.Errorf("JpTIme_JpDay_String %v = %v", test.str, newJpDay)
		}
	}
}

var jphourtests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "零時"},
	{time.Date(2016, time.January, 1, 1, 0, 0, 0, time.Local), "一時"},
	{time.Date(2016, time.January, 1, 10, 0, 0, 0, time.Local), "十時"},
	{time.Date(2016, time.January, 1, 11, 0, 0, 0, time.Local), "十一時"},
}

func TestJpHour_String(t *testing.T) {
	for _, test := range jphourtests {
		jpt := NewJpTime(test.time)
		newJpHour := jpt.JpHour().String()
		if newJpHour != test.str {
			t.Errorf("JpTime_JpHour_String %v = %v", test.str, newJpHour)
		}
	}
}

var jpminutetests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "零分"},
	{time.Date(2016, time.January, 1, 0, 1, 0, 0, time.Local), "一分"},
	{time.Date(2016, time.January, 1, 0, 10, 0, 0, time.Local), "十分"},
	{time.Date(2016, time.January, 1, 0, 11, 0, 0, time.Local), "十一分"},
}

func TestJpMinute_String(t *testing.T) {
	for _, test := range jpminutetests {
		jpt := NewJpTime(test.time)
		newJpMinute := jpt.JpMinute().String()
		if newJpMinute != test.str {
			t.Errorf("JpTime_JpMinute_String %v = %v", test.str, newJpMinute)
		}
	}
}

var jpsecondtests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "零秒"},
	{time.Date(2016, time.January, 1, 0, 0, 1, 0, time.Local), "一秒"},
	{time.Date(2016, time.January, 1, 0, 0, 10, 0, time.Local), "十秒"},
	{time.Date(2016, time.January, 1, 0, 0, 11, 0, time.Local), "十一秒"},
}

func TestJpSecond_String(t *testing.T) {
	for _, test := range jpsecondtests {
		jpt := NewJpTime(test.time)
		newJpSecond := jpt.JpSecond().String()
		if newJpSecond != test.str {
			t.Errorf("JpTime_JpSecond_String %v = %v", test.str, newJpSecond)
		}
	}
}

var kyuurekimonthtests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "睦月"},
	{time.Date(2016, time.February, 1, 0, 0, 0, 0, time.Local), "如月"},
	{time.Date(2016, time.March, 1, 0, 0, 0, 0, time.Local), "弥生"},
	{time.Date(2016, time.April, 1, 0, 0, 0, 0, time.Local), "卯月"},
	{time.Date(2016, time.May, 1, 0, 0, 0, 0, time.Local), "皐月"},
	{time.Date(2016, time.June, 1, 0, 0, 0, 0, time.Local), "水無月"},
	{time.Date(2016, time.July, 1, 0, 0, 0, 0, time.Local), "文月"},
	{time.Date(2016, time.August, 1, 0, 0, 0, 0, time.Local), "葉月"},
	{time.Date(2016, time.September, 1, 0, 0, 0, 0, time.Local), "長月"},
	{time.Date(2016, time.October, 1, 0, 0, 0, 0, time.Local), "神無月"},
	{time.Date(2016, time.November, 1, 0, 0, 0, 0, time.Local), "霜月"},
	{time.Date(2016, time.December, 1, 0, 0, 0, 0, time.Local), "師走"},
}

func TestJpTime_KyuurekiMonth(t *testing.T) {
	for _, test := range kyuurekimonthtests {
		jpt := NewJpTime(test.time)
		newKyuurekiMonth := jpt.KyuurekiMonth()
		if newKyuurekiMonth != test.str {
			t.Errorf("JpTime_KyuurekiMonth %v = %v", test.str, newKyuurekiMonth)
		}
	}
}

type JpTimeWarekiTest struct {
	time    time.Time
	name    string
	initial string
	year    int
}

var warekitests = []JpTimeWarekiTest{
	{time.Date(0, time.December, 31, 0, 0, 0, 0, time.Local), "紀元前", "B.C.", 1},
	{time.Date(1, time.January, 1, 0, 0, 0, 0, time.Local), "西暦", "A.D.", 1},
	{time.Date(1873, time.January, 1, 0, 0, 0, 0, time.Local), "明治", "M", 6},
	{time.Date(1912, time.July, 29, 0, 0, 0, 0, time.Local), "明治", "M", 45},
	{time.Date(1912, time.July, 30, 0, 0, 0, 0, time.Local), "大正", "T", 1},
	{time.Date(1926, time.December, 24, 0, 0, 0, 0, time.Local), "大正", "T", 15},
	{time.Date(1926, time.December, 25, 0, 0, 0, 0, time.Local), "昭和", "S", 1},
	{time.Date(1989, time.January, 7, 0, 0, 0, 0, time.Local), "昭和", "S", 64},
	{time.Date(1989, time.January, 8, 0, 0, 0, 0, time.Local), "平成", "H", 1},
	{time.Date(2016, time.January, 8, 0, 0, 0, 0, time.Local), "平成", "H", 28},
}

func TestJpTime_Wareki(t *testing.T) {
	for _, test := range warekitests {
		jpt := NewJpTime(test.time)
		newName, newInitial, newYear := jpt.Wareki()
		if newName != test.name {
			t.Errorf("JpTime_Wareki %v = %v", test.name, newName)
		}
		if newInitial != test.initial {
			t.Errorf("JpTime_Wareki %v = %v", test.initial, newInitial)
		}
		if newYear != test.year {
			t.Errorf("JpTime_Wareki %v = %v", test.year, newYear)
		}
	}
}

var etotests = []JpTimeTest{
	{time.Date(2008, time.January, 1, 0, 0, 0, 0, time.Local), "子"},
	{time.Date(2009, time.January, 1, 0, 0, 0, 0, time.Local), "丑"},
	{time.Date(2010, time.January, 1, 0, 0, 0, 0, time.Local), "寅"},
	{time.Date(2011, time.January, 1, 0, 0, 0, 0, time.Local), "卯"},
	{time.Date(2012, time.January, 1, 0, 0, 0, 0, time.Local), "辰"},
	{time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local), "巳"},
	{time.Date(2014, time.January, 1, 0, 0, 0, 0, time.Local), "午"},
	{time.Date(2015, time.January, 1, 0, 0, 0, 0, time.Local), "未"},
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), "申"},
	{time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local), "酉"},
	{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local), "戌"},
	{time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local), "亥"},
}

func TestJpTime_Eto(t *testing.T) {
	for _, test := range etotests {
		jpt := NewJpTime(test.time)
		newEto := jpt.Eto()
		if newEto != test.str {
			t.Errorf("JpTime_Eto %v = %v", test.str, newEto)
		}
	}
}

var sekkutests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), ""},
	{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.Local), "人日"},
	{time.Date(2016, time.March, 3, 0, 0, 0, 0, time.Local), "上巳"},
	{time.Date(2016, time.May, 5, 0, 0, 0, 0, time.Local), "端午"},
	{time.Date(2016, time.July, 7, 0, 0, 0, 0, time.Local), "七夕"},
	{time.Date(2016, time.September, 9, 0, 0, 0, 0, time.Local), "重陽"},
}

func TestJpTime_Sekku(t *testing.T) {
	for _, test := range sekkutests {
		jpt := NewJpTime(test.time)
		_, newSekku, _ := jpt.Sekku()
		if newSekku != test.str {
			t.Errorf("JpTime_Sekku %v = %v", test.str, newSekku)
		}
	}
}

var sekki24tests = []JpTimeTest{
	{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local), ""},
	{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.Local), "小寒"},
	{time.Date(2016, time.January, 21, 0, 0, 0, 0, time.Local), "大寒"},
	{time.Date(2016, time.February, 4, 0, 0, 0, 0, time.Local), "立春"},
	{time.Date(2016, time.February, 19, 0, 0, 0, 0, time.Local), "雨水"},
	{time.Date(2016, time.March, 5, 0, 0, 0, 0, time.Local), "啓蟄"},
	{time.Date(2016, time.March, 20, 0, 0, 0, 0, time.Local), "春分"},
	{time.Date(2016, time.April, 4, 0, 0, 0, 0, time.Local), "清明"},
	{time.Date(2016, time.April, 20, 0, 0, 0, 0, time.Local), "穀雨"},
	{time.Date(2016, time.May, 5, 0, 0, 0, 0, time.Local), "立夏"},
	{time.Date(2016, time.May, 20, 0, 0, 0, 0, time.Local), "小満"},
	{time.Date(2016, time.June, 5, 0, 0, 0, 0, time.Local), "芒種"},
	{time.Date(2016, time.June, 21, 0, 0, 0, 0, time.Local), "夏至"},
	{time.Date(2016, time.July, 7, 0, 0, 0, 0, time.Local), "小暑"},
	{time.Date(2016, time.July, 22, 0, 0, 0, 0, time.Local), "大暑"},
	{time.Date(2016, time.August, 7, 0, 0, 0, 0, time.Local), "立秋"},
	{time.Date(2016, time.August, 23, 0, 0, 0, 0, time.Local), "処暑"},
	{time.Date(2016, time.September, 7, 0, 0, 0, 0, time.Local), "白露"},
	{time.Date(2016, time.September, 22, 0, 0, 0, 0, time.Local), "秋分"},
	{time.Date(2016, time.October, 8, 0, 0, 0, 0, time.Local), "寒露"},
	{time.Date(2016, time.October, 23, 0, 0, 0, 0, time.Local), "霜降"},
	{time.Date(2016, time.November, 7, 0, 0, 0, 0, time.Local), "立冬"},
	{time.Date(2016, time.November, 22, 0, 0, 0, 0, time.Local), "小雪"},
	{time.Date(2016, time.December, 7, 0, 0, 0, 0, time.Local), "大雪"},
	{time.Date(2016, time.December, 21, 0, 0, 0, 0, time.Local), "冬至"},
	{time.Date(2100, time.January, 1, 0, 0, 0, 0, time.Local), ""},
}

func TestJpTime_Sekki24(t *testing.T) {
	for _, test := range sekki24tests {
		jpt := NewJpTime(test.time)
		_, newSekki24 := jpt.Sekki24()
		if newSekki24 != test.str {
			t.Errorf("JpTime_Sekki24 %v = %v", test.str, newSekki24)
		}
	}
}

var holidaytests = []JpTimeTest{
	{time.Date(2015, time.January, 1, 0, 0, 0, 0, time.Local), "元日"},
	{time.Date(2015, time.January, 12, 0, 0, 0, 0, time.Local), "成人の日"},
	{time.Date(2015, time.February, 11, 0, 0, 0, 0, time.Local), "建国記念の日"},
	{time.Date(2015, time.March, 21, 0, 0, 0, 0, time.Local), "春分の日"},
	{time.Date(2015, time.April, 29, 0, 0, 0, 0, time.Local), "昭和の日"},
	{time.Date(2015, time.May, 3, 0, 0, 0, 0, time.Local), "憲法記念日"},
	{time.Date(2015, time.May, 4, 0, 0, 0, 0, time.Local), "みどりの日"},
	{time.Date(2015, time.May, 5, 0, 0, 0, 0, time.Local), "こどもの日"},
	{time.Date(2015, time.May, 6, 0, 0, 0, 0, time.Local), "振替休日"},
	{time.Date(2015, time.July, 20, 0, 0, 0, 0, time.Local), "海の日"},
	{time.Date(2015, time.September, 21, 0, 0, 0, 0, time.Local), "敬老の日"},
	{time.Date(2015, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2015, time.September, 23, 0, 0, 0, 0, time.Local), "秋分の日"},
	{time.Date(2015, time.October, 12, 0, 0, 0, 0, time.Local), "体育の日"},
	{time.Date(2015, time.November, 3, 0, 0, 0, 0, time.Local), "文化の日"},
	{time.Date(2015, time.November, 23, 0, 0, 0, 0, time.Local), "勤労感謝の日"},
	{time.Date(2015, time.December, 23, 0, 0, 0, 0, time.Local), "天皇誕生日"},
}

func TestJpTime_Holiday(t *testing.T) {
	tm := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.Local)
	k := 0
	for i := 0; i < 365; i++ {
		jpt := NewJpTime(tm.AddDate(0, 0, i))
		isHoliday, holidayName := jpt.Holiday()
		if isHoliday {
			if holidaytests[k].str != holidayName {
				t.Errorf("JpTime_Holiday %v = %v", holidayName, holidaytests[k].str)
			}
			k++
		}
	}
}

var holidayothertests = []JpTimeTest{
	{time.Date(1947, time.January, 1, 0, 0, 0, 0, time.Local), ""},
	{time.Date(1972, time.January, 2, 0, 0, 0, 0, time.Local), ""},
	{time.Date(1973, time.January, 2, 0, 0, 0, 0, time.Local), ""},
	{time.Date(1988, time.April, 29, 0, 0, 0, 0, time.Local), "天皇誕生日"},
	{time.Date(1989, time.April, 29, 0, 0, 0, 0, time.Local), "みどりの日"},
	{time.Date(1988, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1989, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1990, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1991, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1993, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1994, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1995, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1996, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(1999, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2000, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2001, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2002, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2004, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2005, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2006, time.May, 4, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2032, time.September, 21, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2049, time.September, 21, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2060, time.September, 21, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2077, time.September, 21, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2088, time.September, 21, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2094, time.September, 21, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2009, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2015, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2026, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2037, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2043, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2054, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2071, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2099, time.September, 22, 0, 0, 0, 0, time.Local), "国民の休日"},
	{time.Date(2016, time.August, 11, 0, 0, 0, 0, time.Local), "山の日"},
}

func TestJpTime_HolidayOther(t *testing.T) {
	for _, test := range holidayothertests {
		jpt := NewJpTime(test.time)
		_, newHoliday := jpt.Holiday()
		if newHoliday != test.str {
			t.Errorf("JpTime_HolidayOther %v : %v = %v", test.time, test.str, newHoliday)
		}
	}
}

func BenchmarkFmtIntKanjiMeisuu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(FmtIntKanjiMeisuu(i))
	}
}

func BenchmarkJpTime_Holiday(b *testing.B) {
	t := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.Local)
	for i := 0; i < b.N; i++ {
		jpt := NewJpTime(t.AddDate(0, 0, i))
		_, holidayName := jpt.Holiday()
		fmt.Println(jpt.Format(time.ANSIC) + " = " + holidayName)
	}
}
