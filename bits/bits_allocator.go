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
type IdBits struct {
	signBits uint8
	timestampBits uint8
	workerIdBits uint8
	sequenceBits uint8

	maxDeltaSeconds uint64
	maxWorkerId uint64
	maxSequence uint64

	timestampShift uint8
	workerIdShift uint8
}

// allocate uid,the highest bit always zero mean allocate uid always positive
func (idBits *IdBits) Allocate(timestamp uint64,workid uint64,seq uint64) uint64 {
	return timestamp << idBits.timestampShift | workid << idBits.workerIdShift | seq
}

// get max Delta Seconds
func (idBits *IdBits) GetMaxDeltaSeconds() uint64 {
	return idBits.maxDeltaSeconds
}

// get max work id
func (idBits *IdBits) GetMaxWorkerId() uint64 {
	return idBits.maxWorkerId
}

// get max sequence
func (idBits *IdBits) GetMaxSequence() uint64 {
	return idBits.maxSequence
}

// get the spec timestampBits,workerIdBits,sequenceBits bit_allocator reference
func New(tb uint8,wib uint8,sb uint8) *IdBits {
	return &IdBits{
		signBits:1,
		timestampBits:tb,
		workerIdShift:wib,
		sequenceBits:sb,
		maxDeltaSeconds:-1 ^ (-1 << tb),
		maxWorkerId:-1 ^ (-1 << wib),
		maxSequence:-1 ^ (-1 << sb),
		timestampShift:wib+ sb,
		workerIdBits:sb,
	}
}

// get default bit_allocator reference
func NewDefault() *IdBits {
	return &IdBits{
		signBits:1,
		timestampBits:28,
		workerIdBits:22,
		sequenceBits:13,
		maxDeltaSeconds:-1 ^ (-1 << 28),
		maxWorkerId:-1 ^ (-1 << 22),
		maxSequence:-1 ^ (-1 << 13),
		timestampShift:22+13,
		workerIdShift:13,
	}
}