package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPass(t *testing.T){
	g := NewGomegaWithT(t)
	
	pass := Customer{
		Name : "B6327552", //ผิด
		Email : "b6327552@g.sut.ac.th",
		CustomerID: "H6327552",
	}

	ok,err := govalidator.ValidateStruct(pass)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
}