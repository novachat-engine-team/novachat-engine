package util

import (
	"time"
)

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

func (d *Duration) MarshalText() (text []byte, err error) {
	duration := time.Duration(*d).String()
	text = []byte(duration)
	return
}

/////////////////////////////////////////////////////////////
func NowFormatYMDHMS() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatYMDHMSNow(now time.Time) string {
	return now.Format("2006-01-02 15:04:05")
}

