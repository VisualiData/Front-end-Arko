package main

import (
	"strconv"
	"encoding/json"
	"strings"
	"html/template"
)
// template functions
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	// Add whatever other types you need
	default:
		return ""
	}
}

func Marshal(value interface{}) template.JS {
	a, _ := json.Marshal(value)
	return template.JS(a);
}

func Join(value []interface{}) string {
	size := len(value)
	types := make([]string, size, size * 2)
	for i, v  := range value{
		//fmt.Println(value[i])
		types[i] = v.(string)
	}

	return strings.Join(types, ",")
}