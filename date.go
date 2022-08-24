/**
 * @author [gogogo1024]
 * @email [jxycbjhc@163.com]
 * @create date 2022-08-22 12:10:28
 * @modify date 2022-08-22 12:10:28
 * @desc [description] 属性访问器 getter,setter
 */

// 日历和时间设置
package headfirstgo

import (
	"errors"
	"unicode/utf8"
)

type Date struct {
	year  int
	month int
	day   int
}
type Event struct {
	title string
	Date
}
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) >30  {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
func (e *Event) Title() string {
	return e.title
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil

}
func (d *Date) Month() int {

	return d.month
}
func (d *Date) SetDay(day int) error {
	if day > 31 || day < 1 {
		return errors.New("invalid month")
	}
	d.day = day
	return nil
}
func (d *Date) Day() int {
	return d.day
}
