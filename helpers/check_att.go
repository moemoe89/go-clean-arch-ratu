//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package helpers

import (
	"reflect"
	"strings"
)

// CheckInDBAtt represent the helpers for check attribute inside of struct
func CheckInDBAtt(v interface{}, filterField string) []string {

	filterField = strings.Replace(filterField, " ", "", -1)
	field := strings.Split(filterField, ",")

	out := []string{}
	mapField := make(map[string]bool)
	for _, f := range field {
		mapField[f] = true
	}
	val := reflect.ValueOf(v)
	for i := 0; i < val.Type().NumField(); i++ {
		if tag, ok := val.Type().Field(i).Tag.Lookup("db"); ok {
			if _, ok := mapField[tag]; ok {
				if tag != "-" {
					out = append(out, tag)
				}
			}
		}
	}

	return out
}

// CheckMatchDBAtt represent the helpers for check attribute match of struct
func CheckMatchDBAtt(v interface{}, field string) string {

	var out string
	val := reflect.ValueOf(v)
	for i := 0; i < val.Type().NumField(); i++ {
		if tag, ok := val.Type().Field(i).Tag.Lookup("db"); ok {
			if tag == field {
				if tag != "-" {
					out = tag
				}
			}
		}
	}

	return out
}
