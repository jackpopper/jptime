package jptime

import (
	"testing"
	"time"
)

type JpFormatTest struct {
	name   string
	format string
	time   time.Time
	result string
}

var jpformatTests = []JpFormatTest{
	{"ISO8601", ISO8601, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "2006-01-02T15:04:05+09:00"},
	{"JISX0301", JISX0301, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "H18.01.02"},
	{"JISX0301_OutOfRange", JISX0301, time.Date(1872, time.January, 2, 15, 4, 5, 0, time.Local), ""},
	{"JISX0301JP", JISX0301JP, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "平18.01.02"},
	{"KanjiDate", KanjiDate, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "二〇〇六年一月二日"},
	{"KanjiTime", KanjiTime, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "十五時四分五秒"},
	{"WarekiDate", WarekiDate, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "平成18年1月2日"},
	{"WarekiKanjiDate", WarekiKanjiDate, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "平成十八年一月二日"},
	{"WarekiKanjiDate2", WarekiKanjiDate, time.Date(1872, time.January, 2, 15, 4, 5, 0, time.Local), "一八七二年一月二日"},
	{"JpWeekdayBrackets", JpWeekdayBrackets, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "（月）"},
	{"JpWeekdayString", JpWeekdayString, time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local), "月曜日"},
}

func TestJpTime_JpFormat(t *testing.T) {
	for _, test := range jpformatTests {
		jpt := NewJpTime(test.time)
		result := jpt.JpFormat(test.format)
		if result != test.result {
			t.Errorf("%s expected %q got %q", test.name, test.result, result)
		}
	}
}
