package books_test

import (
	"fmt"
	. "github.com/TaylorOno/ginkgo-example/books"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Book", func() {
	var (
		longBook  *Book
		shortBook *Book
	)

	BeforeEach(func() {
		longBook = &Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		shortBook = &Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		Context("With fewer than 10 pages", func() {
			Specify("should be a article", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		Context("Failed Test", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("Slow Test", func() {
			It("should be a short story", func() {
				time.Sleep(5 * time.Second)
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		When("When Test", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		When("Test", func() {
			It("should be a (short) story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHOR1T STORY"))
			})
		})

		Context("A test has a special ( character", func() {
			It("focus should still work", func() {
				test := "lala"
				fmt.Println(test)
				Expect(true).To(BeTrue())
			})
		})

		Context("This test will panic", func() {
			It("test reporter should show exception", func() {
				panic("panic!!")
			})
		})
	})

	DescribeTable("Book Category Table",
		func(b Book, expected string) {
			Expect(b.CategoryByLength()).To(Equal(expected))
		},
		Entry("Novel", Book{Pages: 2783}, "NOVEL"),
		Entry("Short Story", Book{Pages: 24}, "SHORT STORY"),
	)

	DescribeTable("addition",
		func(a, b, c int) {
			Expect(a + b).To(Equal(c))
		},
		EntryDescription("%d + %d = %d"),
		Entry(nil, 1, 2, 3),
		Entry(EntryDescription("%[3]d = %[1]d + %[2]d"), 10, 100, 110),
		Entry(func(a, b, c int) string { return fmt.Sprintf("%d = %d", a+b, c) }, 4, 3, 7),
	)
})
