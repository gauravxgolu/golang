package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
)

func main() {
	//fileName  := ""
	//if len(os.Args) == 2 {
	//	fileName := os.Args[1]
	//}
	if len(os.Args)!=2 {
		fmt.Println("Invalid aregument, single argument of file path")
		return
	}
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	count := 0
    for scanner.Scan() {
        count++
		byteArr := scanner.Bytes()
		for i:=0; i < len(byteArr); i++ {
			//fmt.Println(byteArr[i])
			if byteArr[i] > 127 {
				fmt.Println("Line no.",  count, " contains UTF-8 character.")
				break
			}
		}
        //fmt.Println(scanner.Bytes())

    }
}
