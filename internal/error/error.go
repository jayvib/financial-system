package error

type Error interface {
	error
	String() string
	OrigErr() error
}
