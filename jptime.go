// Package jptime provides functionality for japanese calendar & time
// 日本の祝日・暦・季節・時刻表示のためのパッケージ
package jptime

import "time"

// A JpTime represents an instant in japanese time.
type JpTime struct {
	time.Time
}

// NewJpTime returns JpTime.
func NewJpTime(t time.Time) JpTime {
	return JpTime{t.In(time.FixedZone("Asia/Tokyo", 9*60*60))}
}

var chineseNumerals = [...]rune{
	'〇',
	'一',
	'二',
	'三',
	'四',
	'五',
	'六',
	'七',
	'八',
	'九',
}

var digitChineseNumerals = [...]rune{
	'一',
	'十',
	'百',
	'千',
	'万',
}

// FmtInt returns 数字（記数法）[0<].
func FmtInt(i int) string {
	buf := make([]rune, 0, 10)
	for i > 0 {
		buf = append(buf, rune(i%10)+'0')
		i /= 10
	}
	return string(reverseRune(buf))
}

// FmtIntKanji returns 漢数字（記数法）[0<]
func FmtIntKanji(i int) string {
	buf := make([]rune, 0, 10)
	for i > 0 {
		buf = append(buf, chineseNumerals[i%10])
		i /= 10
	}
	return string(reverseRune(buf))
}

// FmtIntKanjiMeisuu returns 漢数字（命数法）[0, 99999]
func FmtIntKanjiMeisuu(i int) string {
	if i < 0 || i >= 100000 {
		return ""
	}
	if i == 0 {
		return "零"
	}

	buf := make([]rune, 0, 10)
	for digit := 0; i > 0; digit++ {
		if digit > 0 && i%10 > 0 {
			buf = append(buf, digitChineseNumerals[digit])
		}
		if (digit == 0 && i%10 > 0) || i%10 > 1 {
			buf = append(buf, chineseNumerals[i%10])
		}
		i /= 10
	}
	return string(reverseRune(buf))
}

func reverseRune(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

// A JpYear specifies a year.
type JpYear int

// JpYear returns the year specifies by t.
func (t JpTime) JpYear() JpYear { return JpYear(int(t.Year())) }
func (y JpYear) String() string {
	year := int(y)
	if year <= 0 {
		return "紀元前" + FmtIntKanji(1-year) + "年"
	}
	return FmtIntKanji(year) + "年"
}

// A JpMonth specifies a month of the year.
type JpMonth int

// JpMonth returns the month of the year specifies by t.
func (t JpTime) JpMonth() JpMonth { return JpMonth(int(t.Month())) }
func (m JpMonth) String() string  { return FmtIntKanjiMeisuu(int(m)) + "月" }

var jpdays = [...]string{
	"日",
	"月",
	"火",
	"水",
	"木",
	"金",
	"土",
}

// A JpWeekday specifies a day of the week.
type JpWeekday int

// JpWeekday returns the day of the week specifies by t.
func (t JpTime) JpWeekday() JpWeekday { return JpWeekday(int(t.Weekday())) }
func (w JpWeekday) String() string    { return jpdays[int(w)] }

// A JpDay specifies a day.
type JpDay int

// JpDay returns the day specifies by t.
func (t JpTime) JpDay() JpDay  { return JpDay(t.Day()) }
func (d JpDay) String() string { return FmtIntKanjiMeisuu(int(d)) + "日" }

// A JpHour specifies a hour of the day.
type JpHour int

// JpHour returns the hour within the day specifies by t, in the range [0, 23].
func (t JpTime) JpHour() JpHour { return JpHour(t.Hour()) }
func (h JpHour) String() string { return FmtIntKanjiMeisuu(int(h)) + "時" }

// A JpMinute specifies a minute of the hour.
type JpMinute int

// JpMinute returns the minute offset within the hour specifies by t, in the range [0, 59].
func (t JpTime) JpMinute() JpMinute { return JpMinute(t.Minute()) }
func (m JpMinute) String() string   { return FmtIntKanjiMeisuu(int(m)) + "分" }

// A JpSecond specifies a second of the minute.
type JpSecond int

// JpSecond returns the second offset within the minute specifies by t, in the range [0, 59].
func (t JpTime) JpSecond() JpSecond { return JpSecond(t.Second()) }
func (s JpSecond) String() string   { return FmtIntKanjiMeisuu(int(s)) + "秒" }

var kyuurekimonths = [...]string{
	"睦月",
	"如月",
	"弥生",
	"卯月",
	"皐月",
	"水無月",
	"文月",
	"葉月",
	"長月",
	"神無月",
	"霜月",
	"師走",
}

// KyuurekiMonth returns 月（旧暦）.
func (t JpTime) KyuurekiMonth() string { return kyuurekimonths[t.Month()-1] }

// A Wareki specifies a month of the year in wareki.
type Wareki int

// These are predefined wareki.
const (
	Kigenzen Wareki = iota
	Seireki
	Meiji
	Taisho
	Showa
	Heisei
)

type jpTimeWareki struct {
	name    string
	initial string
	start   time.Time
}

var wareki = [...]jpTimeWareki{
	{name: "紀元前", initial: "B.C."},
	{name: "西暦", initial: "A.D.", start: time.Date(1, time.January, 1, 0, 0, 0, 0, time.Local)},
	{name: "明治", initial: "M", start: time.Date(1873, time.January, 1, 0, 0, 0, 0, time.Local)},
	{name: "大正", initial: "T", start: time.Date(1912, time.July, 30, 0, 0, 0, 0, time.Local)},
	{name: "昭和", initial: "S", start: time.Date(1926, time.December, 25, 0, 0, 0, 0, time.Local)},
	{name: "平成", initial: "H", start: time.Date(1989, time.January, 8, 0, 0, 0, 0, time.Local)},
}

// Wareki returns 和暦.
func (t JpTime) Wareki() (string, string, int) {
	var name, initial string
	var year int
	for i, w := range wareki {
		if (Wareki(i) == Kigenzen || t.After(w.start) || t.Equal(w.start)) &&
			(Wareki(i) == Heisei || t.Before(wareki[i+1].start)) {
			name = w.name
			initial = w.initial
			switch Wareki(i) {
			case Kigenzen:
				year = 1 - t.Year()
			case Meiji:
				// 明治6年からグレゴリオ暦採用のため
				year = t.Year() - w.start.Year() + 5 + 1
			default:
				year = t.Year() - w.start.Year() + 1
			}
			break
		}
	}
	return name, initial, year
}

var eto = [...]string{
	"子",
	"丑",
	"寅",
	"卯",
	"辰",
	"巳",
	"午",
	"未",
	"申",
	"酉",
	"戌",
	"亥",
}

// Eto returns 干支.
func (t JpTime) Eto() string { return eto[(t.Year()+8)%12] }

// Sekku returns 節句.
func (t JpTime) Sekku() (bool, string, string) {
	switch t.Month() {
	case time.January:
		if t.Day() == 7 {
			return true, "人日", "七草の節句"
		}
	case time.March:
		if t.Day() == 3 {
			return true, "上巳", "桃の節句"
		}
	case time.May:
		if t.Day() == 5 {
			return true, "端午", "菖蒲の節句"
		}
	case time.July:
		if t.Day() == 7 {
			return true, "七夕", "笹の節句"
		}
	case time.September:
		if t.Day() == 9 {
			return true, "重陽", "菊の節句"
		}
	}
	return false, "", ""
}

// http://www.h3.dion.ne.jp/~sakatsu/sekki24_topic.htm
type jpTimeSekki24 struct {
	name  string
	month time.Month
	d     float64
	a     float64
}

var sekki24 = [...]jpTimeSekki24{
	{"小寒", time.January, 6.3811, 0.242778},
	{"大寒", time.January, 21.1046, 0.242765},
	{"立春", time.February, 4.8693, 0.242713},
	{"雨水", time.February, 19.7062, 0.242627},
	{"啓蟄", time.March, 6.3968, 0.242512},
	{"春分", time.March, 21.4471, 0.242377},
	{"清明", time.April, 5.6280, 0.242231},
	{"穀雨", time.April, 20.9375, 0.242083},
	{"立夏", time.May, 6.3771, 0.241945},
	{"小満", time.May, 21.9300, 0.241825},
	{"芒種", time.June, 6.5733, 0.241731},
	{"夏至", time.June, 22.2747, 0.241669},
	{"小暑", time.July, 8.0091, 0.241642},
	{"大暑", time.July, 23.7317, 0.241654},
	{"立秋", time.August, 8.4102, 0.241703},
	{"処暑", time.August, 24.0125, 0.241786},
	{"白露", time.September, 8.5186, 0.241898},
	{"秋分", time.September, 23.8896, 0.242032},
	{"寒露", time.October, 9.1414, 0.242179},
	{"霜降", time.October, 24.2487, 0.242328},
	{"立冬", time.November, 8.2396, 0.242469},
	{"小雪", time.November, 23.1189, 0.242592},
	{"大雪", time.December, 7.9152, 0.242689},
	{"冬至", time.December, 22.6587, 0.242752},
}

// Sekki24 returns 二十四節気.
func (t JpTime) Sekki24() (bool, string) {
	year := t.Year()
	if year >= 2100 {
		return false, ""
	}

	if t.Month() == time.January || t.Month() == time.February {
		year--
	}

	for _, s := range sekki24 {
		if t.Month() == s.month && t.Day() == int(s.d+s.a*float64(year-1900)-float64(int((year-1900)/4))) {
			return true, s.name
		}
	}
	return false, ""
}

// Holiday returns 祝日名
func (t JpTime) Holiday() (bool, string) {
	isHoliday, holidayName := t.holiday()
	if !isHoliday {
		isHoliday, holidayName = t.furikaeHoliday()
		if !isHoliday {
			isHoliday, holidayName = t.kokuminHoliday()
		}
	}
	return isHoliday, holidayName
}

// 祝日名.
func (t JpTime) holiday() (bool, string) {
	if t.Year() < 1948 {
		return false, ""
	}

	switch t.Month() {
	case time.January:
		if t.Day() == 1 {
			return true, "元日"
		} else if (t.Year() >= 2000 && t.happyMonday(2)) || (t.Year() < 2000 && t.Day() == 15) {
			return true, "成人の日"
		}
	case time.February:
		if t.Day() == 11 {
			return true, "建国記念の日"
		}
	case time.March:
		if _, sekki := t.Sekki24(); sekki == "春分" {
			return true, "春分の日"
		}
	case time.April:
		if t.Day() == 29 {
			if t.Year() >= 2007 {
				return true, "昭和の日"
			} else if t.Year() >= 1989 {
				return true, "みどりの日"
			} else {
				return true, "天皇誕生日"
			}
		}
	case time.May:
		if t.Day() == 3 {
			return true, "憲法記念日"
		} else if t.Year() >= 2007 && t.Day() == 4 {
			return true, "みどりの日"
		} else if t.Day() == 5 {
			return true, "こどもの日"
		}
	case time.June:
	case time.July:
		if (t.Year() >= 2003 && t.happyMonday(3)) || (t.Year() < 2003 && t.Year() >= 1995 && t.Day() == 20) {
			return true, "海の日"
		}
	case time.August:
		if t.Year() >= 2016 && t.Day() == 11 {
			return true, "山の日"
		}
	case time.September:
		if (t.Year() >= 2003 && t.happyMonday(3)) || (t.Year() < 2003 && t.Year() >= 1966 && t.Day() == 15) {
			return true, "敬老の日"
		} else if _, sekki := t.Sekki24(); sekki == "秋分" {
			return true, "秋分の日"
		}
	case time.October:
		if (t.Year() >= 2003 && t.happyMonday(2)) || (t.Year() < 2003 && t.Year() >= 1966 && t.Day() == 10) {
			return true, "体育の日"
		}
	case time.November:
		if t.Day() == 3 {
			return true, "文化の日"
		} else if t.Day() == 23 {
			return true, "勤労感謝の日"
		}
	case time.December:
		if t.Year() >= 1989 && t.Day() == 23 {
			return true, "天皇誕生日"
		}
	}
	return false, ""
}

// ハッピーマンデー.
func (t JpTime) happyMonday(wnum int) bool {
	if t.Weekday() == time.Monday && t.weekNumber() == wnum {
		return true
	}
	return false
}

// 何週目.
func (t JpTime) weekNumber() int {
	return ((t.Day() - 1) / 7) + 1
}

// 振替休日.
func (t JpTime) furikaeHoliday() (bool, string) {
	if t.Year() < 1973 || t.Weekday() == time.Sunday {
		return false, ""
	}

	for i := 1; i <= 5; i++ {
		prevday := NewJpTime(t.AddDate(0, 0, -i))
		if isHoliday, _ := prevday.holiday(); isHoliday {
			if prevday.Weekday() == time.Sunday {
				return true, "振替休日"
			} else if t.Year() < 2007 {
				// 2007年までは前日のみチェック
				break
			}
		} else {
			break
		}
	}
	return false, ""
}

// 国民の休日.
func (t JpTime) kokuminHoliday() (bool, string) {
	if t.Year() < 1988 || t.Weekday() == time.Sunday {
		return false, ""
	}

	if yIsHoliday, _ := NewJpTime(t.AddDate(0, 0, -1)).holiday(); yIsHoliday {
		if tIsholiday, _ := NewJpTime(t.AddDate(0, 0, 1)).holiday(); tIsholiday {
			return true, "国民の休日"
		}
	}
	return false, ""
}
