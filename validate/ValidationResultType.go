package validate

type ValidationResultType int

const (
	Information ValidationResultType = iota
	Warning     ValidationResultType = iota
	Error       ValidationResultType = iota
)
