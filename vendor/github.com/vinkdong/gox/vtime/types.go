package vtime

import (
	"strconv"
	"time"
	"strings"
	"github.com/vinkdong/gox/log"
)

type Time struct {
	Format string    // timestamp or time layout
	Unit   string    // if format is timestamp it can be ms ns s ..
	Value  string    // a string formatted time value
	Time   time.Time // the standard time
	TZ     string    // if don't set timezoo
}

func (t *Time) ToSpecFormat(format string) string {
	return t.Time.Format(format)
}

func (t *Time) FromRelativeTime(relative string) error {
	if strings.HasPrefix(relative, "now") {

		var (
			now      = time.Now().UTC()
			duration time.Duration
			err      error
		)
		duration, err = time.ParseDuration(strings.TrimPrefix(relative, "now"))
		if err != nil {
			log.Error(err)
			return err
		}

		t.Time = now.Add(duration)
		t.FromTime(t.Time)
	}
	return nil
}

/**
transfer a time format to another format
 */
func (t *Time) Transfer(to *Time) error {
	stdTime, err := t.Parser()
	if err != nil {
		return err
	}
	to.FromTime(stdTime)
	return nil
}

// Parser VTime Format to time
func (t *Time) Parser() (time.Time, error) {
	var err error
	if t.Format == "timestamp" {
		formTimeValue, err := strconv.ParseInt(t.Value, 0, 64)
		switch t.Unit {
		case "ms":
			t.Time = ParserTimestampMs(formTimeValue)
			break
		case "ns":
			t.Time = ParserTimestampNs(formTimeValue)
		case "s":
			t.Time = ParserTimestampS(formTimeValue)
		default:
			t.Time = time.Now()
		}
		return t.Time, err
	} else {
		if t.TZ != "" {
			loc, err := time.LoadLocation(t.TZ)
			if err == nil {
				t.Time, err = time.ParseInLocation(t.Format, t.Value, loc)
				return t.Time, err
			}
		}
		t.Time, err = time.Parse(t.Format, t.Value)
	}
	return t.Time, err
}

/**
parser time to vtime
 */
func (t *Time) FromTime(stdTime time.Time) {
	if t.Format == "timestamp" {
		switch t.Unit {
		case "ms":
			ttime := stdTime.UnixNano()
			t.Value = strconv.FormatInt(ttime, 0)
			return
		case "ns":
			ttime := stdTime.UnixNano()
			t.Value = strconv.FormatInt(ttime, 0)
			return
		case "s":
			ttime := stdTime.Unix()
			t.Value = strconv.FormatInt(ttime, 0)
		default:
			return
		}
	} else {
		if t.Format == ""{
			t.Format = "2006-01-02 15:04:05"
		}
		if t.TZ != ""{
			loc, err := time.LoadLocation(t.TZ)
			if err == nil{
				t.Value = stdTime.In(loc).Format(t.Format)
				return
			}
		}
		t.Value = stdTime.Format(t.Format)
	}
}
