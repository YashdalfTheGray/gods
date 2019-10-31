package gods

// Iterator defines an interface for creating iterators
// over data structures.
type Iterator interface {
	Next() bool
	Get() interface{}
	Error() error
}
