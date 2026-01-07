package govalidator

import (
	. "github.com/onsi/gomega"
	"testing"

	"github.com/Lastjigsaw/sut-final-lab/entity"
)

func TestBooksPositive(t *testing.T) {
	g := NewWithT(t)

    b := Books.New([]
		string{"Cow", "Horse"}
	)
    g.Expect(b.Books()).To(BeTrue(), "It positive test")

}