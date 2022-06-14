package books_test

import (
	"gocode/books"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var (
		// 通过闭包在BeforeEach和It之间共享数据
		longBook  books.Book
		shortBook books.Book
	)
	// 此函数用于初始化Spec的状态，在It块之前运行。如果存在嵌套Describe，则最
	// 外面的BeforeEach最先运行
	BeforeEach(func() {
		longBook = books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = books.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			// 通过It来创建一个Spec
			It("should be a novel", func() {
				// Gomega的Expect用于断言
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

})
