package Abstract_factory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factory := NewSimpleLunchFactory()
	food := factory.CreateLunch()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
