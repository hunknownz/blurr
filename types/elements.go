package types

type Elements interface {
	Size() int
	IGet(int) (Element, error)
	ISet(int, Element) (bool, error)
	DType() DType
	Clone() Elements
}
