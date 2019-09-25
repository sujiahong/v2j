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

func readVdf(fileName string, m *map[string]interface{})  {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	var appid string
	var curM = *m
	var mapArr []map[string]interface{}
	br := bufio.NewReader(f)
	for {
		arrByte, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if appid != "" {
			var startIdx = -1
			var kvArr []string
			var bracesChar string
			for i, char := range arrByte {
				if string(char) == "\"" {
					if startIdx < 0 {
						startIdx = i + 1
					} else {
						kvArr = append(kvArr, string(arrByte[startIdx:i]))
						startIdx = -1
					}
				} else if string(char) == "{" {
					bracesChar = "{"
					break
				} else if string(char) == "}" {
					bracesChar = "}"
					break
				}
			}
			switch bracesChar {
			case "{":
				break
			case "}":
				curM = mapArr[len(mapArr)-1]
				mapArr = mapArr[:len(mapArr)-1]
				break
			default:
				arrLen := len(kvArr)
				if (arrLen > 1){
					curM[kvArr[0]] = kvArr[1]
				}else{
					mapArr = append(mapArr, curM)
					curM[kvArr[0]] = make(map[string]interface{})
					curM = curM[kvArr[0]].(map[string]interface{})
				}
			}
			kvArr = kvArr[0:0]
			startIdx = -1
			bracesChar = ""
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
	m := make(map[string]interface{})
	readVdf("sample.vdf", &m)
	m1 := m["381210"].(map[string]interface{})["common"].(map[string]interface{})["name"]
	fmt.Println(m1)
}
