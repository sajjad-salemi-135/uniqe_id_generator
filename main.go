package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	sequencebits = 2
	timebits     = 32

	maxsequencebits = (1 << sequencebits) - 1
	maxtimebits     = (1 << timebits) - 1

	sequenceshift = timebits
)

type snowflake struct {
	sequence int64
	time     int64
	mux      sync.Mutex
}

func Newsnowflake() *snowflake {
	return &snowflake{
		sequence: 0,
		time:     -1,
	}
}

func (s *snowflake) Generator() int64 {
	s.mux.Lock()
	defer s.mux.Unlock()

	epoch := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now().UnixMilli() - epoch.UnixMilli()

	if now > maxtimebits {
		fmt.Println("Timestamp exceeds 32-bit range!")
		return -1
	}

	if now == s.time {
		s.sequence = (s.sequence + 1) & maxsequencebits
		if s.sequence == 0 {
			for now <= s.time {
				now = time.Now().UnixMilli() - epoch.UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.time = now

	id := (s.sequence << sequenceshift) | s.time

	return id
}

func main() {
	sr := Newsnowflake()
	for i := 0 ; i < 10 ; i++{
		id := sr.Generator()
		ids := fmt.Sprintf("%034b\n", id)
		fmt.Println(ids)
	}
	
}
