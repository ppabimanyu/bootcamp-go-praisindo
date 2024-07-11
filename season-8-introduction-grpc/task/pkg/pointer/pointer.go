/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package pointer

import (
	"fmt"
	"time"
)

// String function takes any type of data and returns a pointer to a string if the data is not nil.
// If the data is nil, it returns nil.
func String(data any) *string {
	if data != nil {
		v := data.(string)
		return &v
	}
	return nil
}

// DefaultString function takes any type of data and returns a string if the data is not nil.
// If the data is nil, it returns an empty string.
func DefaultString(data any) string {
	if data != nil {
		v := data.(string)
		return v
	}
	return ""
}

// Time function takes any type of data and returns a pointer to a time.Time if the data is not nil.
// If the data is nil, it returns nil.
func Time(data any) *time.Time {
	if data != nil {
		v := data.(time.Time)
		return &v
	}
	return nil
}

// Int function takes any type of data and returns a pointer to an int if the data is not nil.
// If the data is nil, it returns nil.
func Int(data any) *int {
	if data != nil {
		v := data.(int)
		return &v
	}
	return nil
}

// Int64 function takes any type of data and returns a pointer to an int64 if the data is not nil.
// If the data is nil, it returns nil.
func Int64(data any) *int64 {
	if data != nil {
		v := data.(int64)
		return &v
	}
	return nil
}

// Float64 function takes any type of data and returns a pointer to a float64 if the data is not nil.
// If the data is nil, it returns nil.
func Float64(data any) *float64 {
	if data != nil {
		v := data.(float64)
		return &v
	}
	return nil
}

// Float64ToString function takes any type of data and returns a pointer to a string representation of a float64 if the data is not nil.
// If the data is nil, it returns nil.
func Float64ToString(data any) *string {
	if data != nil {
		v := fmt.Sprint(data.(float64))
		return &v
	}
	return nil
}

// Bool function takes any type of data and returns a pointer to a bool if the data is not nil.
// If the data is nil, it returns nil.
func Bool(data any) *bool {
	if data != nil {
		v := data.(bool)
		return &v
	}
	return nil
}

// DefaultBool function takes any type of data and returns a bool if the data is not nil.
// If the data is nil, it returns false.
func DefaultBool(data any) bool {
	if data != nil {
		return data.(bool)
	}
	return false
}
