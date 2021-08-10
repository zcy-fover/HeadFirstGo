package calendar

import (
	"errors"
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetDate(year int, month int, day int) error {
	err := d.SetDay(day)
	if err != nil {
		return err
	}
	err = d.SetMonth(month)
	if err != nil {
		return err
	}
	err = d.SetYear(year)
	if err != nil {
		return err
	}
	return nil
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("输入年无效")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("输入月份无效")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("输入天数无效")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Display() {
	fmt.Println(*d)
}
