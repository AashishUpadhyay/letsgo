package menu

import (
	"testing"
)

func TestAdd(t *testing.T) {

	// arrange
	itemName := "Coffee"
	test_data := menu{
		{name: "Caramel Machiato", prices: map[string]float64{"small": 1.65, "medium": 1.95, "large": 2.15}},
		{name: "Penne Pollo", prices: map[string]float64{"half": 5.65, "full": 10.95, "double": 20.15}},
		{name: "Coffee", prices: map[string]float64{}},
	}
	want := "item already exists!"

	// act
	got := test_data.add(itemName)

	// assert
	if got.Error() != want {
		t.Errorf("got %q, wanted %q", got.Error(), want)
	}
}
