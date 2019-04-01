package generator

import "distributed-uid-generator/bits"

type DefaultGenerator struct {
	bits.IdBits
}

func (defaultgenerator *DefaultGenerator) GetNextId() uint64 {
	return 0
}

func (defaultgenerator *DefaultGenerator) ParseId() *IdDetail {
	return nil
}