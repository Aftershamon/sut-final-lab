package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestTypePass(t *testing.T) {
	g := NewGomegaWithT(t)

	name := Customer{
		Name:       "B6327552",
		Email:      "b6327552@g.sut.ac.th",
		CustomerID: "H1234567", //ผิด
	}

	ok, err := govalidator.ValidateStruct(name)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ผิดรูปแบบ"))

}
func TestTypeNumMax(t *testing.T) {
	g := NewGomegaWithT(t)

	name := Customer{
		Name:       "B6327552",
		Email:      "b6327552@g.sut.ac.th",
		CustomerID: "H12345678", //ผิด
	}

	ok, err := govalidator.ValidateStruct(name)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ผิดรูปแบบ"))

}
func TestTypeNumMin(t *testing.T) {
	g := NewGomegaWithT(t)

	name := Customer{
		Name:       "B6327552",
		Email:      "b6327552@g.sut.ac.th",
		CustomerID: "H123456", //ผิด
	}

	ok, err := govalidator.ValidateStruct(name)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ผิดรูปแบบ"))

}

func TestTypeChack(t *testing.T) {
	g := NewGomegaWithT(t)

	name := Customer{
		Name:       "B6327552",
		Email:      "b6327552@g.sut.ac.th",
		CustomerID: "H1234567", //ผิด
	}

	ok, err := govalidator.ValidateStruct(name)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ผิดรูปแบบ"))

}
