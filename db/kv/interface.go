package db

// user may end the iteration prematurely by returning false
type ApplyFn func([]byte, []byte) bool

type KvDb interface {
	GetKvValue([]byte) ([]byte, error)
	GetAllKeyValue(ApplyFn) error
	SetKvKeyValue([]byte, []byte) error
	UpdateKvExistingValue([]byte, []byte) error
	Close()
}
