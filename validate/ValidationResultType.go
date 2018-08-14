package validate

type ValidationResultType int

const (
	Information ValidationResultType = iota
	Warning
	Error
)
