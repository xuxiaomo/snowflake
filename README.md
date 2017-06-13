# Unique ID Generator

## func NewWorker
    func NewWorker(workerid int64) (iw *IDWorker, err error)
NewWorker renturn an unique id generator

## func (iw *IDWorker) NextID
    func (iw *IDWorker) NextID() (uid int64, err error)
NextID return an unique id

## example
    iw := NewWorker(1)
    id := iw.NextID()

