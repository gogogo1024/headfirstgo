// 接收器替代其他语言中的方法重载
package headfirstgo

import "fmt"

type Liters float64

type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}


func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}
