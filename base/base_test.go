package base

import (
	"errors"
	"reflect"
	"testing"
)

func Test_CustomError(t *testing.T) {
	x := ErrorCustom("test error type message.")

	var check error
	check = errors.New("check error type message. ")

	if reflect.TypeOf(x) != reflect.TypeOf(check) {
		t.Error("custom error type, check error!")
	} else {
		t.Log("custom error type, check ok!")
	}
}

func Test_Splitstr(t *testing.T) {
	x := "start_end"

	y := Splitstr(x, 1, 3)

	if y != "sta" {
		t.Error("split string, check error! want to get 'sta', but result is:", y)
	} else {
		t.Log("split string, check ok!")
	}
}

func Test_IntToStr(t *testing.T) {
	x := int(1)
	y := IntToStr(x)
	if y != "1" {
		t.Error("change int type to string, check error! right result is '1', but result is:", y)
	} else {
		t.Log("change int type to string, check ok!")
	}
}

func Test_Int64ToStr(t *testing.T) {
	x := int64(123456789)
	y := Int64ToStr(x)
	if y != "123456789" {
		t.Error("change int64 type to string, check error! right result is '123456789', but result is:", y)
	} else {
		t.Log("change int64 type to string, check ok!")
	}
}

func Test_StrToInt(t *testing.T) {
	x := "123"
	y, err := StrToInt(x)
	if err != nil {
		t.Error("change string type to int, found error:", err)
	}

	if y != int(123) {
		t.Error("change string type to int, check error! right result is 123, but result is:", y)
	} else {
		t.Log("change string type to int, check ok!")
	}
}

func Test_StrToInt64(t *testing.T) {
	x := "1234567"
	y, err := StrToInt64(x)
	if err != nil {
		t.Error("change string type to int64, found error:", err)
	}

	if y != int64(1234567) {
		t.Error("change string type to int64, check error! right result is 1234567, but result is:", y)
	} else {
		t.Log("change string type to int64, check ok!")
	}
}

func Test_Float64ToStr(t *testing.T) {
	x := float64(3.1415926)
	y := Float64ToStr(x)
	if y != "3.14" {
		t.Error("change float64 type to string, check error! right result is '3.14', but result is:", y)
	} else {
		t.Log("change float64 type to string, check ok!")
	}
}

func Test_StrToFloat64(t *testing.T) {
	x := "3.1415926"
	y, err := StrToFloat64(x)
	if err != nil {
		t.Error("change string type to float64, found error:", err)
	}

	if y != float64(3.1415926) {
		t.Error("change string type to float64, check error! right result is 3.1415926, but result is:", y)
	} else {
		t.Log("change string type to float64, check ok!")
	}
}
