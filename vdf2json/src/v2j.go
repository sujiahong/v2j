package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"bytes"
)

func describe(i interface{}){
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func readVdf(fileName string) {
	f, err := os.Open(fileName)
	if (err != nil){
		return
	}
	defer f.Close()
	var appid string
	m := make(map[string]interface{})
	// curM := m
	br := bufio.NewReader(f)
	for{
		arrByte, _, err := br.ReadLine()
		if (err == io.EOF){
			break
		}
		// fmt.Println(string(arrByte))
		if (appid != ""){
			strLen := len(arrByte)
			fmt.Println(len(arrByte))
			if (strLen == 1){
				if (string(arrByte) == "{"){
					continue
				}else{

				}
			}else{
				bytes
			}
		}else{
			arr := bytes.Split(arrByte, []byte(","))
			fmt.Println(len(arr), cap(arr))
			if (len(arr) > 1){
				for _, item := range arr{
					idx := bytes.Index(item, []byte("AppID"))
					if (idx > -1){
						appid = string(item[8:])
						m[appid] = make(map[string]interface{})
						break
					}
				}
			}
		}
	}
}

func main(){
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
	//readVdf("sample.vdf")
}