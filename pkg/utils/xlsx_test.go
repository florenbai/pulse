package utils

import (
	"encoding/json"
	"testing"
)

func TestAAA(t *testing.T) {
	//var names = []string{"aa", "bbb"}
	StrPoiRefV1()
}

func StrPoiRefV1() {
	type abc struct {
		CA  string          `json:"ca"`
		A   string          `json:"a"`
		B   string          `json:"b"`
		CB  string          `json:"cb"`
		Jso json.RawMessage `json:"jso"`
	}
	var aaa = json.RawMessage{}
	err := aaa.UnmarshalJSON([]byte(`jso1`))
	if err != nil {
		return
	}
	var bbb json.RawMessage = []byte("1111131423148903124098903214")

	var array = []abc{
		{CA: "ca1", A: "a1", B: "b1", CB: "cb1", Jso: bbb},
		{CA: "ca2", A: "a2", B: "b2", CB: "cb2", Jso: bbb},
		{CA: "ca3", A: "a3", B: "b3", CB: "cb3", Jso: bbb},
		{CA: "ca4", A: "a4", B: "b4", CB: "cb4", Jso: bbb},
		{CA: "ca5", A: "a5", B: "b5", CB: "cb5", Jso: bbb},
	}
	var data interface{}
	data = &array
	//firstCharacter := 65

	//s := reflect.ValueOf(data)
	//t := reflect.TypeOf(data)
	//fmt.Println("xlsx:",s.Kind() == reflect.Ptr )

	//var mapArray = make([]map[string]interface{}, 0)
	//if err := AnyToAny(&array, &mapArray); err != nil {
	//	fmt.Println("err:", err)
	//}
	//data = mapArray
	//mv := reflect.ValueOf(data)
	//mt := reflect.TypeOf(data)
	//fmt.Println("xlsx:", s.Kind() == reflect.Ptr)
	//if s.Kind() == reflect.Ptr {
	//	data = reflect.Value.Elem(s).Interface()
	//	s = reflect.ValueOf(data)
	//	//t = reflect.TypeOf(records)
	//}
	//fmt.Println("xlsx:", s.Kind() == reflect.Map, s.Kind())
	//for i := 0; i < s.Len(); i++ {
	//	elem := s.Index(i).Interface()
	//	elemType := reflect.TypeOf(elem)
	//	elemValue := reflect.ValueOf(elem)
	//	for j := 0; j < elemType.NumField(); j++ {
	//		field := elemType.Field(j)
	//		//fieldVal := elemValue.Field(j)
	//		//fmt.Println("field:", field.Name)
	//		fmt.Println(" elemValue.Field(j).Interface():", elemValue.Field(j).Interface())
	//		name := field.Name
	//		fmt.Println("name:", name)
	//		column := string(rune(firstCharacter + j))
	//		fmt.Println("column:", column)
	//
	//	}
	//}
	//var f = StructWriteXlsx("Sheet1", data)
	var f, _ = MapWriteXlsx("Sheet1", data)
	if err := f.SaveAs("_ums.xlsx"); err != nil {
		println(err.Error())
	}
}
