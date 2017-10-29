package demo
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
	TArray
	TSlice
	TType
	TMap
	TConcurrency
)
func NewDemo(t DemoType) IDemo {
	switch t {
	case TPointer: return NewPointerDemo()
	case TArray: return NewArrayDemo()
	case TSlice: return NewSliceDemo()
	case TMap: return NewMapDemo()
	case TType: return NewTypeDemo()
	case TConcurrency: return NewConcurrencyDemo()
	default: return &Demo{"void","abstract demo."}
	}
}


type Demo struct {
	name, desc string
}

func (d Demo) String() string {
	return fmt.Sprintf("%-12s #%s",d.name, d.desc)
}

func (d Demo) Do() {
	panic("not implemented")
}


