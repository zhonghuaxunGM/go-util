package main

import (
	"fmt"
)

func main() {
	// consul utils
	s, _ := GetConsulInfo("test", "ip:port")
	fmt.Println(string(s))
	WriteConsulInfo("test", `{'test':tes}`, "ip:port")
	DelConsulInfo("test", "ip:port")
}
