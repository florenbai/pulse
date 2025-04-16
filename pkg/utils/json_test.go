package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack"
)

func TestJsonAA(t *testing.T) {
	Json6()
}

func Jso() {
	var dd = map[string]string{"ab": "ab1", "ac": "ac1", "ad": "ad1"}
	var aa = map[string]json.RawMessage{}
	marshal, err := json.Marshal(dd)
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}
	err = json.Unmarshal(marshal, &aa)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println("marshal:", string(marshal))
	fmt.Println("aa:", aa)
	for k, v := range aa {
		fmt.Println("k:", k)
		fmt.Println("v:", string(v))
	}
}

func Jsov2() {
	var ee = "kkabcccc"
	var ff = json.RawMessage{}
	//marshal, err := json.Marshal(&ee)
	//if err != nil {
	//	fmt.Println("Marshal err:", err)
	//	return
	//}
	err := json.Unmarshal([]byte(ee), &ff)
	//err = json.Unmarshal(marshal, &ff)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	//fmt.Println("marshal:", marshal, string(marshal), ee)
	fmt.Println("ff:", ff)
	fmt.Println("ff:", string(ff))
}

func Jsov3() {
	var ee = "kkabcccc"
	var ff = json.RawMessage{}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(ee); err != nil {
		fmt.Println("gob enc err:", err)
		return
	}
	//bytes.NewBuffer和bytes.Buffer类似，只不过可以传入一个初始的byte数组，返回一个指针
	dec := gob.NewDecoder(bytes.NewBuffer(buf.Bytes()))
	//调用Decode方法,传入结构体对象指针，会自动将buf.Bytes()里面的内容转换成结构体
	if err := dec.Decode(&ff); err != nil {
		fmt.Println("gob dec err:", err)
		return
	}

	//fmt.Println("marshal:", marshal, string(marshal), ee)
	fmt.Println("ff:", ff)
	fmt.Println("ff:", string(ff))
}
func Jsov4() {
	var ee = "kkabcccc"
	var ff = json.RawMessage{}
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	// 序列化
	structJson, err := jsonIterator.Marshal(ee)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	// 打印输出结果
	fmt.Println("输出序列化结果: ", structJson, string(structJson))

	//fmt.Println("marshal:", marshal, string(marshal), ee)
	fmt.Println("ff:", ff)
	fmt.Println("ff:", string(ff))
}

func Jsov5() {
	var ee = "kkabcccc"
	var ff = json.RawMessage{}
	// 序列化
	structJson, err := msgpack.Marshal(ee)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	// 打印输出结果
	fmt.Println("输出序列化结果: ", structJson, string(structJson))
	err = msgpack.Unmarshal(structJson, &ff)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	//fmt.Println("marshal:", marshal, string(marshal), ee)
	fmt.Println("ff:", ff)
	fmt.Println("ff:", string(ff))
}
