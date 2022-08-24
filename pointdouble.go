/**
 * @author [gogogo1024]
 * @email [jxycbjhc@163.com]
 * @create date 2022-08-20 17:05:33
 * @modify date 2022-08-20 17:05:33
 * @desc [description] 指针
 */
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
