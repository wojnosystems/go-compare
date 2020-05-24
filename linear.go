package go_compare

type Linear uint8

const (
	Equal = iota
	GreaterThan
	LessThan
)

func IsGreaterThanOrEqual(comparable Linear) bool {
	return comparable == GreaterThan || comparable == Equal
}

func IsLessThanOrEqual(comparable Linear) bool {
	return comparable == LessThan || comparable == Equal
}
