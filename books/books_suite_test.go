package books_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = BeforeSuite(func() {
	log.Println("this is a before suite line")
	log.Println("logs in this block should be ignored")
	//Expect("error").ToNot(HaveOccurred())
})
