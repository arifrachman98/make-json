package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var m map[string]string
	m = make(map[string]string)

	fmt.Println("Please enter your name : ")
	br := bufio.NewReader(os.Stdin)
	bname, _, _ := br.ReadLine()
	name := string(bname)
	m["name"] = name

	fmt.Println("Please enter your address : ")
	br_n := bufio.NewReader(os.Stdin)
	b_address, _, _ := br_n.ReadLine()
	address := string(b_address)
	m["address"] = address

	b, err := json.Marshal(m)

	if err != nil {
		fmt.Println("encoding failed !")
	} else {
		fmt.Print("Encode Data : ")
		fmt.Println(b)
		fmt.Print("Decode Data : ")
		fmt.Println(string(b))
	}
}
