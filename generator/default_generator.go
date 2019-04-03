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
		sequence:0,
		workid:wid,
		timestamp:getCurrentTime(),
		beginDateTime:bdt,
	}
}

func (defaultgenerator *DefaultGenerator) GetNextId() uint64 {
	currentTime := getCurrentTime()
	duration := currentTime - defaultgenerator.timestamp
	if (duration < 0) {
		panic("occur clock move back stop generator id,wait time catch up")
	}

	if (duration == 0) {

	} else {

		defaultgenerator.timestamp = currentTime
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

func checkError(err error,msg string) {
	if nil != err {
		panic(msg + err.Error())
	}
}