package tdd

import (
	. "./model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Captcha", func() {
  Context("First pattern with plus sign", func() {
    It("should return One when input for first operand is 1", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 1, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("One"))
    })
    It("should return Two when input for first operand is 2", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 2, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("Two"))
    })    
    It("should return Three when input for first operand is 3", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 3, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("Three"))
    })
  })
})
