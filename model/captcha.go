package model 
type Captcha struct {
  Pattern, LeftOperand, Operator, RightOperand int
}

func (captcha Captcha) GetLeftOperand() string {
  if captcha.LeftOperand == 1 {
  	return "One"
  } else {
    return "Two"
  }
}