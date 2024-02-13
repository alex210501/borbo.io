package ds

import "fmt"

type DataStucture struct {
	Name string
	Age  int
}

func (d *DataStucture) String() string {
	return fmt.Sprintf("%v (%v years)", d.Name, d.Age)
}
