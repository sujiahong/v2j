package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func readVdf(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	var appid string
	// m := make(map[string]interface{})
	// var curM = m
	// var mapArr []map[string]interface{}
	br := bufio.NewReader(f)
	for {
		arrByte, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		// fmt.Println(string(arrByte))
		if appid != "" {
			var startIdx = -1
			var kvArr []string
			var c string
			for i, char := range arrByte {
				if string(char) == "\"" {
					if startIdx < 0 {
						startIdx = i + 1
					} else {
						kvArr = append(kvArr, string(arrByte[startIdx:i]))
						startIdx = -1
					}
				} else if string(char) == "{" {
					c = "{"
					break
				} else if string(char) == "}" {
					c = "}"
					break
				}
			}
			switch c {
			case "{":
				// 	append(mapArr, curM)
				// 	curM[appid] = make(map[string]interface{})
				// 	curM = curM[appid].(map[string]interface{})
				break
			case "}":
				// 	curM = mapArr[len(mapArr)-1]
				// 	mapArr = mapArr[:len(mapArr)-1]
				break
			default:
				fmt.Println(kvArr)
			}
			kvArr = kvArr[0:0]
		} else {
			arr := bytes.Split(arrByte, []byte(","))
			fmt.Println(len(arr), cap(arr))
			if len(arr) > 1 {
				for _, item := range arr {
					idx := bytes.Index(item, []byte("AppID"))
					if idx > -1 {
						appid = string(item[8:])
						break
					}
				}
			}
		}
	}
}

func main() {
	fmt.Printf("hello ,, go!!\n")
	var a = "runoob"
	fmt.Println(a)
	i := 55
	st := struct {
		name string
	}{
		name: "difosoooofh!!!",
	}
	describe(a)
	describe(i)
	describe(st)
	m := make(map[string]interface{})
	fmt.Println(m)
	readVdf("sample.vdf")
}
