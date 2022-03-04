package helpers

import (
	"fmt"
	"reflect"
	"time"
)

func Empty(val interface{}) bool {
	if val == nil {
		return true
	}

	value := reflect.ValueOf(val)

	switch value.Kind() {
	case reflect.String, reflect.Array:
		return value.Len() == 0
	case reflect.Map, reflect.Slice:
		return value.Len() == 0 || value.IsNil()
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()

	}
	return reflect.DeepEqual(val, reflect.Zero(value.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
