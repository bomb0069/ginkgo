package model 
type Captcha struct {
  Pattern, LeftOperand, Operator, RightOperand int
}

func (captcha Captcha) GetLeftOperand() string {
  return "One"
}
