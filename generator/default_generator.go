package generator

import (
	"time"
)

type DefaultGenerator struct {
	sequence uint64
	workid uint64
	timestamp int64
	beginDateTime string
}

func New(wid uint64,bdt string) (dg *DefaultGenerator) {
	return &DefaultGenerator{
		workid:wid,
		beginDateTime:bdt,
	}
}

func (defaultgenerator *DefaultGenerator) GetNextId() uint64 {
	duration := getCurrentTime() - defaultgenerator.timestamp
	if (duration < 0) {
		panic("occur clock move back stop generator id,wait time catch up")
	}

	return 0
}

func (defaultgenerator *DefaultGenerator) ParseId() *IdDetail {
	return nil
}

// get current time sec since January 1, 1970 UTC.
func getCurrentTime() int64 {
	return time.Now().Unix()
}