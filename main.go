package main

import (
	"fmt"
	"strings"
)

func main() {

	dataString := "Android 10, android 10 sdk, android 11 SDK, android 11,"
	splitString := strings.Split(dataString, ", ")
	// splitString = strings.Split(splitString, ",")

	var dataOutput []interface{}
	for _, item := range splitString {

		if strings.Contains(strings.ToLower(item), "android 10") {
			// fmt.Println("Kena validasi")
			continue
		} else {
			dataOutput = append(dataOutput, item)
			// fmt.Println("Engga kena")
		}

	}
	fmt.Println("dataOutput: ", dataOutput)
}
