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
	MaxWorkerID   = -1 ^ (-1 << WorkerIDBits)
	SenquenceBits = 12

	WorkerIDBitsMask = 0x3ff
	SenquenceMask    = 0xfff
)

type IDWorker struct {
	workerID      int64
	senquenceID   int64
	lastTimeStamp int64
	lock          *sync.Mutex
}

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

func (iw *IDWorker) NextID() (id int64, err error) {

}

func timeGen() (ct int64) {
	ct = time.Now().UnixNano() / 1000 / 1000
	return
}

func timeReGen()
