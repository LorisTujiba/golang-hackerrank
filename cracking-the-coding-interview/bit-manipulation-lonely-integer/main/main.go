package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var length int
var input string
var bucket []string
var lonely int64
var datas = make(map [string]int)

func main(){


	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader,&length)
	input,_ = reader.ReadString('\n')

	if length != 1{
		bucket = strings.Fields(input)
		fmt.Print(check())

	} else{
		fmt.Print(input)
	}

}

func check()int64 {
	for i:=0;i<length;i++{
		if _,exist := datas[bucket[i]];!exist{
			datas[bucket[i]] = 1
		}else if val,exist := datas[bucket[i]];exist{
			val += 1
			datas[bucket[i]]= val
		}
	}
	for val,n := range datas{
		if n == 1{
			lonely,_ = strconv.ParseInt(val,10,64)
		}
	}
	return lonely
}