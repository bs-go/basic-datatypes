package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ASCII codes for printable characters (33 - 127)
// 33 corresponds to "!"
// Total # of printable characters in the ASCII table is 94
var MIN = 0
var MAX = 94

func random2(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random2(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func genPass() {
	var LENGTH int64 = 8
	arguments := os.Args
	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			LENGTH = t
		}
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("Using default values...")
	}
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	fmt.Println(getString(LENGTH))
}
