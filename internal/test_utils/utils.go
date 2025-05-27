package test_utils

import "reflect"

func HasMutex(c interface{}) bool {
	val := reflect.ValueOf(c).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == "mu" {
			return true
		}
	}
	return false
}
