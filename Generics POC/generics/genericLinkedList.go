package generics

type GenericNode[T any] struct {
	Data T
	Next *GenericNode[T]
}
