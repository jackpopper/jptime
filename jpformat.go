package jptime

// These are predefined layouts for use in JpTime.JpFormat
const (
	ISO8601           = "2006-01-02T15:04:05-07:00"
	JISX0301          = "H18.01.02"
	JISX0301JP        = "平18.01.02"
	KanjiDate         = "二〇〇六年一月二日"
	KanjiTime         = "一五時四分五秒"
	WarekiDate        = "平成18年1月2日"
	WarekiKanjiDate   = "平成一八年一月二日"
	JpWeekdayBrackets = "（月）"
	JpWeekdayString   = "月曜日"
)

// JpFormat returns a textual representation for japanese format.
func (t JpTime) JpFormat(layout string) string {
	switch layout {
	case JISX0301:
		_, initial, year := t.Wareki()
		if len(initial) == 1 {
			return initial + FmtInt(year) + t.Format(".01.02")
		}
	case JISX0301JP:
		name, initial, year := t.Wareki()
		if len(initial) == 1 {
			buf := []rune(name)
			return string(buf[0]) + FmtInt(year) + t.Format(".01.02")
		}
	case KanjiDate:
		return t.JpYear().String() + t.JpMonth().String() + t.JpDay().String()
	case KanjiTime:
		return t.JpHour().String() + t.JpMinute().String() + t.JpSecond().String()
	case WarekiDate:
		name, _, year := t.Wareki()
		return name + FmtInt(year) + "年" + t.Format("1月2日")
	case WarekiKanjiDate:
		name, initial, year := t.Wareki()
		if len(initial) == 1 {
			return name + FmtIntKanjiMeisuu(year) + "年" + t.JpMonth().String() + t.JpDay().String()
		}
		return t.JpYear().String() + t.JpMonth().String() + t.JpDay().String()
	case JpWeekdayBrackets:
		return "（" + t.JpWeekday().String() + "）"
	case JpWeekdayString:
		return t.JpWeekday().String() + "曜日"
	default:
		return t.Format(layout)
	}
	return ""
}
