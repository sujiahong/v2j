package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readVdf(fileName string, m *map[string]interface{})  {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	var curM = *m
	var mapArr []map[string]interface{}
	br := bufio.NewReader(f)
	for {
		arrByte, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
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
			}else if(arrLen == 1){
				mapArr = append(mapArr, curM)
				curM[kvArr[0]] = make(map[string]interface{})
				curM = curM[kvArr[0]].(map[string]interface{})
			}
		}
		kvArr = kvArr[0:0]
		startIdx = -1
		bracesChar = ""
	}
}

func main() {
	m := make(map[string]interface{})
	readVdf("sample.vdf", &m)
	m1 := m["381210"].(map[string]interface{})["depots"].(map[string]interface{})["branches"].(map[string]interface{})["public"].(map[string]interface{})["buildid"]
	fmt.Println(m1)
}
