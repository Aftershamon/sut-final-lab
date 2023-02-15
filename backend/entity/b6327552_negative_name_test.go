package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	name := Customer{
		Name: "", //ผิด
	}

	ok, err := govalidator.ValidateStruct(name)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name can not Blank"))

}
