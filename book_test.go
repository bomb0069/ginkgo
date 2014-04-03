package tdd 

import (
  . "./model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

  var longBook Book = Book{Title: "Harry Potter", Author: "JK Rolling", Pages: 1000}
  var shortBook Book = Book{Title: "Turn the ship around", Author: "Someone", Pages: 100}

  Describe("Categorizing book length", func() {
    Context("With more than 300 pages", func() {
      It("should be a novel", func() {
        Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
      })
    })
    Context("With fewer than 300 pages", func() {
      It("should ne an article", func() {
        Expect(shortBook.CategoryByLength()).To(Equal("ARTICLE"))
      })
    })
  })
})
