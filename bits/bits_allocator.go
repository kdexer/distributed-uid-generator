package bits

// signBits is highest bit for uid and used for positive or negative (should always 0 means positive)
// timestampBits used for time
// workerIdBits used for work id
// sequenceBits used for sequence
// maxDeltaSeconds means the max value for time
// maxWorkerId means the max value for work id
// maxSequence means the max value for sequence
// timestampShift means should move bits used for generator uid
// workerIdShift means should move bits used for generator uid
// sequenceMask means use sequenceMask and sequence operation
type IdBits struct {
	signBits        uint8
	timestampBits   uint8
	workerIdBits    uint8
	sequenceBits    uint8
	sequenceMask    int64
	maxDeltaSeconds int64
	maxWorkerId     int64
	maxSequence     int64
	timestampShift  uint8
	workerIdShift   uint8
}

// allocate uid,the highest bit always zero mean allocate uid always positive
func (idBits *IdBits) Allocate(timestamp int64, workid int64, seq int64) int64 {
	return timestamp<<idBits.timestampShift | workid<<idBits.workerIdShift | seq
}

// get sequenceMask value
func (idBits *IdBits) GetSequenceMask() int64 {
	return idBits.sequenceMask
}

// get max Delta Seconds
func (idBits *IdBits) GetMaxDeltaSeconds() int64 {
	return idBits.maxDeltaSeconds
}

// get max work id
func (idBits *IdBits) GetMaxWorkerId() int64 {
	return idBits.maxWorkerId
}

// get max sequence
func (idBits *IdBits) GetMaxSequence() int64 {
	return idBits.maxSequence
}

// get the spec timestampBits,workerIdBits,sequenceBits bit_allocator reference
func NewBitAllocator(tbit uint8, wbit uint8, sbit uint8) *IdBits {
	return &IdBits{
		signBits:        1,
		timestampBits:   tbit,
		workerIdShift:   wbit,
		sequenceBits:    sbit,
		maxDeltaSeconds: -1 ^ (-1 << tbit),
		maxWorkerId:     -1 ^ (-1 << wbit),
		maxSequence:     -1 ^ (-1 << sbit),
		timestampShift:  wbit + sbit,
		workerIdBits:    sbit,
		sequenceMask:    -1 ^ (-1 << sbit),
	}
}

// get default bit_allocator reference
func NewDefaultBitAllocator() *IdBits {
	return &IdBits{
		signBits:        1,
		timestampBits:   28,
		workerIdBits:    22,
		sequenceBits:    13,
		maxDeltaSeconds: -1 ^ (-1 << 28),
		maxWorkerId:     -1 ^ (-1 << 22),
		maxSequence:     -1 ^ (-1 << 13),
		timestampShift:  22 + 13,
		workerIdShift:   13,
		sequenceMask:    -1 ^ (-1 << 13),
	}
}
