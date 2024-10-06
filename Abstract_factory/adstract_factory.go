package Abstract_factory

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {
}

func (r Rise) Cook() {
	fmt.Println("It is Rise")
}

type Tomato struct {
}

func (t Tomato) Cook() {
	fmt.Println("It is Tomato")
}

type LunchFactory interface {
	CreateLunch() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct{}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

func (f *SimpleLunchFactory) CreateLunch() Lunch {
	return &Rise{}
}

func (f *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
