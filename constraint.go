package wrench

type Number interface {
	int | int64 | uint | int8 | int32
}
