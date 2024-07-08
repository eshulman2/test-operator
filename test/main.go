package main

import (
	"fmt"
	"reflect"
	"slices"
)


type extraConfigmapsMounts struct {
	Name      string
	MountPath string
	SubPath   string
}

type some struct{
	part1 []extraConfigmapsMounts
	part2 []extraConfigmapsMounts
}

type some2 struct{
	part1 string
}

type helperStruct struct{
	value any
}

func main() {
	var list1 = []extraConfigmapsMounts{extraConfigmapsMounts{Name: "slice1", MountPath: "a", SubPath: "a"},{Name: "bla", MountPath: "a", SubPath: "a"},{Name: "slice1", MountPath: "a", SubPath: "a" }}
	var list2 = []extraConfigmapsMounts{extraConfigmapsMounts{Name: "list2", MountPath: "a", SubPath: "a" }}

	// bla := some{part1: list1, part2: list2}
	// v1 := reflect.ValueOf(bla).FieldByName("part1")

	// fmt.Println(v1.Type())

    sliceType := reflect.TypeOf(list1)

    //use of MakeSlice() method
    intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)

    intSliceReflect = reflect.AppendSlice(intSliceReflect, reflect.ValueOf(list1))
	intSliceReflect = reflect.AppendSlice(intSliceReflect, reflect.ValueOf(list2))
	unique := reflect.MakeSlice(sliceType, 0, 0)

	// var unique []reflect.Value
	type key struct{ value1, value2 any }
	var m []any

	for i := 0; i < intSliceReflect.Len(); i++ {
		k := key{intSliceReflect.Index(i).FieldByName("Name").Interface(), intSliceReflect.Index(i).FieldByName("MountPath").Interface()}
		if ok := slices.IndexFunc(m, func(h any) bool {return h == k}); ok != -1 {
			unique.Index(ok).Set(intSliceReflect.Index(i))
		} else {
			m = append(m, k)
			unique = reflect.Append(unique, intSliceReflect.Index(i))
		}
	}
	fmt.Println(unique)

	var unique2 []extraConfigmapsMounts
	type key2 struct{ value1, value2 string }
	var m2 []any

	for _, v := range list1 {
		k := key2{v.Name, v.MountPath}
		if ok := slices.IndexFunc(m2, func(h any) bool {return h == k}); ok != -1 {
			unique2[ok] = v
		} else {
			m2 = append(m2, k)
			unique2 = append(unique2, v)
		}
	}
	fmt.Println(unique2)

	result:= map[interface{}]interface{}{3: "banana", "something": 7}
	rv2 := reflect.ValueOf(result)
	someval := rv2.MapIndex(reflect.ValueOf(34))
	if someval == (reflect.Value{}){
		fmt.Println("Oh no this is not defined")
	}else{
		fmt.Println(rv2.MapIndex(reflect.ValueOf(3)))
	}
}
