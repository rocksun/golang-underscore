package underscore

import (
	"testing"

	fjson "github.com/json-iterator/go"
)

type cloneModel struct {
	ID   string
	Name string
	Age  int
}

func TestClone_Array(t *testing.T) {
	arr := []int{1, 2, 3}
	duplicate := make([]int, 0)
	Clone(arr, &duplicate)
	ok := All(duplicate, func(n, i int) bool {
		return arr[i] == n
	})
	if !ok {
		t.Error(arr, duplicate)
	}
}

func TestClone_Struct(t *testing.T) {
	m := cloneModel{
		ID:   "id",
		Name: "name",
		Age:  11,
	}
	duplicate := new(cloneModel)
	Clone(m, &duplicate)
	sOld, _ := fjson.MarshalToString(m)
	sNew, _ := fjson.MarshalToString(duplicate)
	if sOld != sNew {
		t.Error(sOld, sNew)
	}
}

func TestChain_Clone_Array(t *testing.T) {
	arr := []int{1, 2, 3}
	duplicate := make([]int, 0)
	Chain(arr).Clone().Value(&duplicate)
	ok := All(duplicate, func(n, i int) bool {
		return arr[i] == n
	})
	if !ok {
		t.Error(arr, duplicate)
	}
}

func TestChain_Clone_Map(t *testing.T) {
	dic := map[string]int{
		"a": 1,
	}
	duplicate := make(map[string]int)
	Chain(dic).Clone().Value(&duplicate)
	if duplicate["a"] != 1 {
		t.Error(duplicate)
	}
}

func TestChain_Clone_Struct(t *testing.T) {
	m := cloneModel{
		ID:   "id",
		Name: "name",
		Age:  11,
	}
	duplicate := new(cloneModel)
	Chain(m).Clone().Value(&duplicate)

	sOld, _ := fjson.MarshalToString(m)
	sNew, _ := fjson.MarshalToString(duplicate)
	if sOld != sNew {
		t.Error(sOld, sNew)
	}
}
