package jsontest

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonTest() {
	jsonTestInner()
}

type configuration struct {
	Enabled bool
	Path    string
}

func jsonTestInner() {
	file, _ := os.Open("/Users/jinxin/go/src/geometry/jsontest/conf.json")

	defer file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)

	conf := configuration{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("path:" + conf.Path)
}
