package demo

import (
	"fmt"
)

type MapDemo struct {
	Demo
}

func NewMapDemo() *MapDemo {
	return &MapDemo{ Demo{"map", "typical Map concept."} }
}

func (d MapDemo) Do() {
	//NOTE: basic
	type Vertex struct{
		Lat, Long float64
	}
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	//NOTE: map iterals
	m = map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google": Vertex{37.42202, -122.08408},
	}
	for k,v := range m {
		fmt.Println(k,v)
	}

	//NOTE: get, delete and exist
	val, ok := m["Google"]
	fmt.Println(val, ok)
	delete(m, "Google")
	val, ok = m["Google"]
	fmt.Println(val, ok)
}
