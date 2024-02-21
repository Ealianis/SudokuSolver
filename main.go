package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("nothing to see here, yet...")
}

// Not used yet, under construction.
func FindDepth(a any) int {
	i := 0
	for {
		if isHigherDimension, ok := a.([]any); ok {
			i++
			a = isHigherDimension[0]
		} else {
			return i
		}
	}
}

// Not used yet, under construction.
func DelveDimensionalArray(a any, depth int) {
	value := reflect.ValueOf(a)

	// If value is not a slice or array, return
	if kind := value.Kind(); kind != reflect.Slice && kind != reflect.Array {
		return
	}

	// If value is a 1D array or slice, print it
	if value.Len() == 0 || value.Index(0).Kind() != reflect.Slice && value.Index(0).Kind() != reflect.Array {
		fmt.Println(value.Interface())
		return
	}

	// Otherwise, recurse on each element
	for i := 0; i < value.Len(); i++ {
		DelveDimensionalArray(value.Index(i).Interface(), 0)
	}
}
