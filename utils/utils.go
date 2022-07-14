package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CheckText(value1 string, value2 string) bool {
	return strings.TrimSpace(value1) == strings.TrimSpace(value2)
}

func CheckSizeSlab(slab string, value string) bool {
	slabArr := strings.Split(slab, "-")
	return value >= slabArr[0] && value <= slabArr[1]
}

func ToInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func WriteConsole(inputRecords [][]string) {
	for _, valueLog := range inputRecords {
		fmt.Println(valueLog)
	}
}
