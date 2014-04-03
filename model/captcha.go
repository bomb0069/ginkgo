package model 
import "strconv"
type Captcha struct {
  Pattern, LeftOperand, Operator, RightOperand int
}

func (captcha Captcha) GetLeftOperand() string {
  var NumberInString map[int]string = map[int]string{ 1: "One", 2: "Two", 3: "Three", 4: "Four",5: "Five",6: "Six",7: "Seven",8: "Eight", 9: "Nine"}
  return NumberInString[captcha.LeftOperand]
}
  
func (captcha *Captcha) GetRightOperand() string {
  return strconv.Itoa(captcha.RightOperand) 
}
  
func (captcha *Captcha) GetResult() string {
  return strconv.Itoa(captcha.LeftOperand + captcha.RightOperand)
}