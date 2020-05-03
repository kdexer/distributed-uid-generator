package generator

import (
	"github.com/kdexer/distributed-uid-generator/bits"
	"sync"
	"time"
)

var mu sync.Mutex

/*
	default generator
 */
type DefaultGenerator struct {
	sequence  int64
	workid    int64
	timestamp int64
	epochStr  string
	*bits.IdBits
}


func New(wid int64, epoch string, tbit uint8, wbit uint8, sbit uint8) (dg *DefaultGenerator) {
	allocator := bits.NewBitAllocator(tbit, wbit, sbit)
	if wid > allocator.GetMaxWorkerId() || wid < 0 {
		panic("worker Id can't be greater than " + string(allocator.GetMaxWorkerId()) + " or less than 0")
	}
	return &DefaultGenerator{
		sequence:  0,
		workid:    wid,
		timestamp: getCurrentTimestamp(),
		epochStr:  epoch,
		IdBits:    allocator,
	}
}

func (dg *DefaultGenerator) GetNextId() int64 {
	//todo use channel
	mu.Lock()
	defer mu.Unlock()
	currentTimestamp := getCurrentTimestamp()
	lastTimestamp := dg.timestamp
	if currentTimestamp < lastTimestamp {
		panic("occur clock move back stop generator id,wait time catch up")
	}
	if currentTimestamp == lastTimestamp {
		dg.sequence = (dg.sequence + 1) & dg.IdBits.GetSequenceMask()
		if 0 == dg.sequence {
			// when get uid over maxSequence value at same time,should wait next time
			currentTimestamp = getNextTime(lastTimestamp)
		}
	} else {
		dg.sequence = 0
	}
	dg.timestamp = currentTimestamp
	return dg.IdBits.Allocate(getDuration(dg.epochStr, currentTimestamp), dg.workid, dg.sequence)
}

// get current time sec since January 1, 1970 UTC.
func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func checkError(err error, msg string) {
	if nil != err {
		panic(msg + err.Error())
	}
}

func getDuration(epoch string, currentTimestamp int64) int64 {
	epochTimestamp, e := time.Parse("2006-01-02", epoch)
	checkError(e, "epoch time string parse error")
	duration := currentTimestamp - epochTimestamp.Unix()
	if duration < 0 {
		panic("current time less then epoch time")
	}
	return duration
}

func getNextTime(lastTimestamp int64) int64 {
	currentTimeStamp := getCurrentTimestamp()
	for currentTimeStamp <= lastTimestamp {
		currentTimeStamp = getCurrentTimestamp()
	}
	return currentTimeStamp
}
