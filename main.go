package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"realm-challenge/assignment"
	"realm-challenge/utils"
	"strings"
)

var partnersRecords, inputRecords [][]string

func main() {
	output := assignment.SolveProblem1(inputRecords, partnersRecords)
	fmt.Println("The output of problem is:")
	utils.WriteConsole(output)
}

func init() {
	/* loadding csv */
	inputRecords = InputReader("input.csv")
	for i := 0; i < len(inputRecords); i++ {
		for j := 0; j < len(inputRecords[i]); j++ {
			inputRecords[i][j] = strings.TrimSpace(inputRecords[i][j])
		}
	}

	partnersRecords = InputReader("partners.csv")
	for i := 0; i < len(partnersRecords); i++ {
		for j := 0; j < len(partnersRecords[i]); j++ {
			partnersRecords[i][j] = strings.TrimSpace(partnersRecords[i][j])
		}
	}
}

//InputReader ...
func InputReader(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
