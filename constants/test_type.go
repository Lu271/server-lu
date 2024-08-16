package constants

type TestType int32

const (
	Test1 = iota + 1
	Test2
	Test3
)

func (t TestType) String() string {
	switch t {
	case Test1:
		return "test1"
	case Test2:
		return "test2"
	case Test3:
		return "test3"
	default:
		return "unknown"
	}
}
