package main

import (
	"reflect"
	"testing"
)

func TestInsertionInteger(t *testing.T) {
	inputSlice := []int{2, 1, 3, 1}
	expectedSlice := []int{1, 1, 2, 3}
	result := InsertionSort(inputSlice)
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Fatalf("Valor esperado : %v , Valor encontrado %v ", expectedSlice, inputSlice)
	}
}
