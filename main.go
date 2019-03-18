package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	defer func() {
		fatal := recover()
		if fatal != nil {
			fmt.Println("程序崩溃，错误信息为:", fatal)
			fmt.Println("按回车退出")
			fmt.Scanln()
		}
	}()
	filepath := "./data.json"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读文件错误")
		fmt.Println("按回车退出")
		fmt.Scanln()
		return
	}
	var ob interface{}
	err = json.Unmarshal(data, &ob)
	if err != nil {
		fmt.Println("json反格式化失败")
		fmt.Println("按回车退出")
		fmt.Scanln()
		return
	}
	_, result, er := SetType(ob)
	if er != nil {
		fmt.Println("格式转义错误")
		fmt.Println("按回车退出")
		fmt.Scanln()
		return
	}
	esData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json格式化失败")
		fmt.Println("按回车退出")
		fmt.Scanln()
		return
	}
	ioutil.WriteFile("./target.json", esData, 0644)
	fmt.Println("按回车退出")
	fmt.Scanln()
}

//SetType ...
func SetType(value interface{}) (isEnd bool, result []string, err error) {
	log.Println(value)
	switch value.(type) {
	case string:
		{
			return true, make([]string, 0), nil
		}
	case map[string]interface{}:
		{
			obmap, ok := value.(map[string]interface{})
			if ok {
				arr := make([]string, 0, len(obmap))
				for k, v := range obmap {
					is, res, e := SetType(v)
					if e != nil {
						log.Fatalln(e)
						return false, make([]string, 0), e
					}
					if is {
						arr = append(arr, k)
					} else {
						for i := 0; i < len(res); i++ {
							arr = append(arr, k+"."+res[i])
						}
					}
				}
				return false, arr, nil
			}
			return false, make([]string, 0), nil
		}
	case []interface{}:
		{
			return true, make([]string, 0), nil
		}
	default:
		{
			return true, make([]string, 0), nil
		}
	}
}

//EsValue ... EsValue
type EsValue map[string]interface{}
