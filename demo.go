package main
import(
	"fmt"
)

type IDemo interface {
	Do()
}

type DemoType int
const (
	TVoid DemoType = iota
	TPointer
	TSlice
)
func NewDemo(t DemoType) IDemo {
	switch t {
	case TPointer: return NewPointerDemo()
	case TSlice: return NewSliceDemo()
	default: return &Demo{"void","abstract demo."}
	}
}


type Demo struct {
	name, desc string
}

func (d Demo) String() string {
	return fmt.Sprintf("%s #%s",d.name, d.desc)
}

func (d Demo) Do() {
	panic("not implemented")
}


