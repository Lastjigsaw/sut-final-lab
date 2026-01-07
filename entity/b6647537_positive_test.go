package govalidator

import (
	. "github.com/onsi/gomega"
	"testing"

	"github.com/Lastjigsaw/sut-final-lab/entity"
)

func TestFarmHasCow(t *testing.T) {
	g := NewWithT(t)

	f := farm.New([]string{"Cow", "Horse"})
	g.Expect(f.HasCow()).To(BeTrue(), "Farm should have cow")
}

func TestEach(t *testing.T) {
	// TODO Maybe refactor?
	t.Parallel()
	acc := 0
	data := []interface{}{1, 2, 3, 4, 5}
	var fn Iterator = func(value interface{}, index int) {
		acc = acc + value.(int)
	}
	Each(data, fn)
	if acc != 15 {
		t.Errorf("Expected Each(..) to be %v, got %v", 15, acc)
	}
}
