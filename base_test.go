package main

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
	}
}
