package books_test

import (
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func lala(s string) {
	fmt.Println("lala")
}

var _ = ginkgo.Describe("test", func() {
	gomega.Expect(true).To(gomega.BeTrue())
})
