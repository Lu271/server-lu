package util

import "server-lu/constants"

func GetTestString(id int32) string {
	return constants.TestType(id).String()
}
