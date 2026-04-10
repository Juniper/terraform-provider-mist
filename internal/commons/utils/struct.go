package mist_utils

import "reflect"

// StringOrNumbe

// IsStructEmpty returns true if all pointer/slice/map fields in the
// struct are nil (or empty for slices/maps).  Works with any struct
func IsSdkDataEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Ptr, reflect.Interface:
			if !f.IsNil() {
				return false
			}
		case reflect.Slice, reflect.Map:
			if f.Len() > 0 {
				return false
			}
		case reflect.String:
			if f.String() != "" {
				return false
			}
		default:
			if !f.IsZero() {
				return false
			}
		}
	}
	return true
}
