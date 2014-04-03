package model 
type Captcha struct {
  Pattern, LeftOperand, Operator, RightOperand int
}

func (captcha Captcha) GetLeftOperand() string {
  var NumberInString = [3]string{"One", "Two", "Three"}
  return NumberInString[captcha.LeftOperand-1]
}