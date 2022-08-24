/**
 * @author [gogogo1024]
 * @email [jxycbjhc@163.com]
 * @create date 2022-08-23 16:07:59
 * @modify date 2022-08-23 16:07:59
 * @desc [description] 接口
 */
package headfirstgo

import (
	"fmt"
)

type TapePlayer struct {
	Batteries string
}

type TapeRecord struct {
	songs []string
}

type Player interface {
	Play(string)
	Stop()
}

func PlayList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
func (tp TapePlayer) Play(song string) {
	fmt.Println("playing", song)
}
func (tp TapePlayer) Stop() {
	fmt.Println("stopped")
}

func (tr TapeRecord) Play(song string) {
	fmt.Println("playing", song)
}
func (tr TapeRecord) Stop() {
	fmt.Println("stopped")
}

// 既有接口，又有类型上的方法
type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

// error
// 1. 自定义error
type ComedyError string

func (ce ComedyError) Error() string {
	return string(ce)
}

// 2.除了error字符串外，更多的信息
type OverHeatError float64

func (vhe OverHeatError) Error() string {
	return fmt.Sprintf("overheating by %0.2f degrees!", vhe)
}
func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverHeatError(excess)
	}
	return nil
}

// Stringer其实就是其他语言的重写string

type Meters float64
type Centimeters float64
type Millimeters float64

func (me Meters) String() string {
	return fmt.Sprintf("%0.2f Meters", me)
}
func (ct Centimeters) String() string {
	return fmt.Sprintf("%0.2f Centimeters", ct)
}
func (mi Millimeters) String() string {
	return fmt.Sprintf("%0.2f Millimeters", mi)
}
