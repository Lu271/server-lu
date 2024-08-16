package util

import (
	"fmt"
	"server-lu/constants"
	"testing"
)

func TestGetTestString(t *testing.T) {
	str := GetTestString(constants.Test2)
	fmt.Println(str)
}
