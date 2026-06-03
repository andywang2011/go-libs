package timex

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type JSONTime time.Time

const (
	DateTimeLayout = "2006-01-02 15:04:05"
	DateLayout     = "2006-01-02"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%q", time.Time(t).Format(DateTimeLayout))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+DateTimeLayout+`"`, string(data), time.Local)
	*t = JSONTime(now)
	return err
}

func (t JSONTime) String() string {
	return time.Time(t).Format(DateTimeLayout)
}

func (t JSONTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *JSONTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = JSONTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}
