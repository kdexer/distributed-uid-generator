package generator

type IdDetail struct {
	timestamp uint64
	workid uint64
	seq uint64
}

type NextId interface {
	GetNextId() uint64
}

type ParseId interface {
	ParseId() *IdDetail
}