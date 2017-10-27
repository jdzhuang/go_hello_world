package main

type IDemo interface {
	Do()
}

type VoidDemo struct {
}

func (d VoidDemo) String() string {
	return "void"
}

func (d VoidDemo) Do() {
	panic("not implemented")
}
