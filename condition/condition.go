package condition

import "github.com/hunknownz/blurr/types"

const (
	ComparatorGT    = ">"
	ComparatorLT    = "<"
	ComparatorGTE   = ">="
	ComparatorEq    = "="
	ComparatorLTE   = "<="
	ComparatorIn    = "in"
	ComparatorIsNan = "is_nan"
)

type seCompItem struct {
	comparator string
	Value      types.Element
}

type seCondVal struct {
	cond     *SeCondition
	compItem *seCompItem
	IsNot    bool
}

type SeCondition struct {
	params []seCondVal
}

type dfCompItem struct {
	comparator string
	column     string
	Value      types.Element
}

type dfCondVal struct {
	cond     *DFCondition
	compItem *dfCompItem
	IsNot    bool
}

type DFCondition struct {
	params []dfCondVal
}
