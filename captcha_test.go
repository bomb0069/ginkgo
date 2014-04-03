package tdd

import (
	. "./model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Captcha", func() {
  Context("First pattern with plus sign", func() {
    It("should return One when input for Left operand is 1", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 1, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("One"))
    })
    It("should return Two when input for Left operand is 2", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 2, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("Two"))
    })    
    It("should return Three when input for Left operand is 3", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 3, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("Three"))
    })
    It("should return Four when input for Left operand is 4", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 4, Operator: 1, RightOperand: 1}
      Expect(captcha.GetLeftOperand()).To(Equal("Four"))
    })
    
    It("should return 1 when input for Right operand is 1", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 4, Operator: 1, RightOperand: 1}
      Expect(captcha.GetRightOperand()).To(Equal("1"))
    })
    
    It("should return 2 when input for Right operand is 2", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 4, Operator: 1, RightOperand: 2}
      Expect(captcha.GetRightOperand()).To(Equal("2"))
    })
    
    It("should return 3 when input for Plus with Left Operand is 1 and Right operand is 2", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 1, Operator: 1, RightOperand: 2}
      Expect(captcha.GetResult()).To(Equal("3"))
    })
    It("should return 4 when input for Plus with Left Operand is 2 and Right operand is 2", func() {
      var captcha = Captcha{Pattern: 1, LeftOperand: 2, Operator: 1, RightOperand: 2}
      Expect(captcha.GetResult()).To(Equal("4"))
    })
  })
})
