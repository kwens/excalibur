/**
 * @Author: kwens
 * @Date: 2022-10-10 09:13:43
 * @Description:
 */
package time

import (
	"strconv"
	"time"
)

const (
	DAY_FORMAT_YYYY_MM_DD  = "2006-01-02"
	DAYTIME_FORMAT         = "2006-01-02 15:04:05"
	DAYTIME_COMPACT_FORMAT = "20060102150405"
)

// GetBeforeDay 获取n月前的日期
func GetBeforeMonthDay(n int) string {
	timeNow := time.Now()
	beforeDay := timeNow.AddDate(0, -1*n, 0)
	return beforeDay.Format(DAY_FORMAT_YYYY_MM_DD)
}

func DayStr2Time(s string) (time.Time, error) {
	return time.ParseInLocation(DAY_FORMAT_YYYY_MM_DD, s, time.Local)
}

func MustTimestamp2TimeStr(ts string) string {
	t, _ := strconv.Atoi(ts)
	tm := time.Unix(int64(t), 0)
	return tm.Format(DAYTIME_FORMAT)
}

func DayTimeStr2Timestamp(ts string) (string, error) {
	t, err := time.ParseInLocation(DAY_FORMAT_YYYY_MM_DD, ts, time.Local)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(t.Unix())), nil
}

func Must2DayTimeStr(t time.Time) string {
	return t.Format(DAYTIME_FORMAT)
}

func Must2DayStr(t time.Time) string {
	return t.Format(DAY_FORMAT_YYYY_MM_DD)
}

func Must2CompactDayTimeStr(t time.Time) string {
	return t.Format(DAYTIME_COMPACT_FORMAT)
}

func MustStr2DayTime(s string) (t time.Time) {
	t, _ = time.ParseInLocation(DAYTIME_FORMAT, s, time.Local)
	return t
}

func MustStr2TimestampInt64(s string) int64 {
	t, _ := time.ParseInLocation(DAYTIME_FORMAT, s, time.Local)
	return t.Unix()
}

func MustTimestampInt642TimeStr(ts int64) string {
	t := time.Unix(ts, 0).Format(DAYTIME_FORMAT)
	return t
}
