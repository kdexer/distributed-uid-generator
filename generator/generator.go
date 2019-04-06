package generator

type IdDetail struct {
	timestamp uint64
	workid    uint64
	seq       uint64
}

type NextId interface {
	GetNextId() int64
}

type ParseId interface {
	ParseId() *IdDetail
}
