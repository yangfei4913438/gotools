package gotools

import (
	"errors"
	"github.com/yangfei4913438/gotools/base"
	"reflect"
	"testing"
)

func Test_CustomError(t *testing.T) {
	x := base.ErrorCustom("test error type message.")

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

	y := base.Splitstr(x, 1, 3)

	if y != "sta" {
		t.Error("split string, check error! want to get 'sta', but result is:", y)
	} else {
		t.Log("split string, check ok!")
	}
}
