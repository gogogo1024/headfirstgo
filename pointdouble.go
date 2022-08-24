// 指针
package headfirstgo

type Student struct {
	Name string
	age int
	Address
	province string
}
type Address struct {
	province string
}
func CreatePoint() *float64 {
	var myFloat = 18.8
	return &myFloat
}
func DoubleNumber(number *int) {
	*number *= 2
}
