package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "problems.csv"

type question struct {
	firstOperand  int
	secondOperand int
	result        int
}

func openCsvFile() *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open '%s'\n", fileName)
		os.Exit(1)
	}
	return file
}

func getRecordsFromFile(file *os.File) [][]string {
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read all records from '%s'\n", fileName)
		os.Exit(1)
	}
	return records
}

func parseRecords(records [][]string) ([]question, error) {
	questions := make([]question, len(records))
	for i, record := range records {
		// Assume that each record comes in the format [xxx+yyy zzz]
		operands := strings.Split(record[0], "+")
		firstOperand, err := strconv.ParseInt(operands[0], 10, 32)
		if err != nil {
			return nil, err
		}
		secondOperand, err := strconv.ParseInt(operands[1], 10, 32)
		if err != nil {
			return nil, err
		}
		result, err := strconv.ParseInt(record[1], 10, 32)
		if err != nil {
			return nil, err
		}
		questions[i] = question{
			firstOperand:  int(firstOperand),
			secondOperand: int(secondOperand),
			result:        int(result),
		}
	}
	return questions, nil
}

func askQuestions(questions []question, s *bufio.Scanner) {
	for i, q := range questions {
		fmt.Printf("Question %d: %d + %d = ", i, q.firstOperand, q.secondOperand)
		for s.Scan() {
			answer, err := strconv.ParseInt(s.Text(), 10, 32)
			if err != nil {
				fmt.Println("Invalid number")
			}
		}
	}
}

func main() {
	file := openCsvFile()
	records := getRecordsFromFile(file)
	questions, err := parseRecords(records)
	if err != nil {
		fmt.Println("Failed to parse csv records")
		os.Exit(1)
	}
	fmt.Println(questions)
}
