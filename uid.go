//Package snowflake is the only ID Generator
package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	CEpoch        = 1497960000000
	WorkerIDBits  = 10
	SenquenceBits = 12

	WorkerIDShift  = SenquenceBits
	TimeStampShift = SenquenceBits + WorkerIDBits

	MaxWorkerID   = -1 ^ (-1 << WorkerIDBits)
	SenquenceMask = -1 ^ (-1 << SenquenceBits)
)

type IDWorker struct {
	workerID      int64
	senquenceID   int64
	lastTimeStamp int64
	lock          *sync.Mutex
}

//NewWorker renturn a id generator
func NewWorker(workerid int64) (iw *IDWorker, err error) {
	iw = new(IDWorker)
	if workerid > MaxWorkerID || workerid < 0 {
		return nil, errors.New("workerid is over range")
	}
	iw.workerID = workerid
	iw.senquenceID = 0
	iw.lastTimeStamp = -1
	iw.lock = new(sync.Mutex)
	return iw, nil
}

//NextID return a unique id
func (iw *IDWorker) NextID() (id int64, err error) {
	iw.lock.Lock()
	defer iw.lock.Unlock()

	ct := timeGen()

	if iw.lastTimeStamp > ct {
		return 0, errors.New("Clock moved backwards, Refuse gen id")
	}

	if iw.lastTimeStamp == ct {
		iw.senquenceID = (iw.senquenceID + 1) & SenquenceMask
		if iw.senquenceID == 0 {
			ct = timeReGen(iw.lastTimeStamp)
		}
	} else {
		iw.senquenceID = 0
	}

	iw.lastTimeStamp = ct

	id = (iw.lastTimeStamp-CEpoch)<<TimeStampShift | iw.workerID<<WorkerIDShift | iw.senquenceID
	return id, nil
}

//timeGen return current time
func timeGen() (ct int64) {
	ct = time.Now().UnixNano() / 1000 / 1000
	return
}

//timeReGen wait until the next millisecond
func timeReGen(last int64) int64 {
	ct := timeGen()
	for {
		if ct <= last {
			ct = timeGen()
		} else {
			break
		}
	}
	return ct
}
