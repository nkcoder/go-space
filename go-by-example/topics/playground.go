package topics

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func TimeNow() (time.Time, time.Time) {
	now := time.Now()
	nowUTC := time.Now().UTC()
	return now, nowUTC
}

func StartOfDayInUTC(t time.Time) time.Time {
	loc, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		panic(err)
	}

	now := time.Now().In(loc).UTC()
	nowUtc2 := time.Now().UTC()

	fmt.Println("now in utc", now, "nowUtc2", nowUtc2)

	return now
}

func TestUUID() {
	x := ""
	v, err := uuid.Parse(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
