package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ask(records *[][]string, marks *int, total *int, ch chan int) {
	*total = len(*records)

	//asking questions one by one, and taking answers from user
	for _, record := range *records {
		fmt.Println(record[0])
		var ans string
		fmt.Scan(&ans)

		//TrimSpace because in records there can be trailing and leading spaces which combines take the string
		//but while scanning the answer we won't consider spaces... so we trim the spaces from ans in record
		if ans == strings.TrimSpace(record[1]) {
			*marks++
		}
	}

	//sending data to channel that the user completed test
	ch <- 1
}

func main() {
	//creating flags

	//Csv flag is for taking csv filename from user (default is "problem.csv")
	filename := flag.String("Csv", "problem.csv", "this csv file contains the question answer")

	//TestDuration flag is for taking the timelimit in seconds of the test from user (default is 30s)
	timelimit := flag.Int("TestDuration", 30, "this is used for test duration")

	//parsing the flags
	flag.Parse()

	//opening the user specified filename
	quiz, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	//creating a NewReader for reading
	quizreader := csv.NewReader(quiz)

	//reading the records
	records, err := quizreader.ReadAll()
	if err != nil {
		fmt.Println("error occured")
	}

	// creating variables for marks and total
	var marks, total int = 0, 0
	var ch = make(chan int)

	//Alerting the user that the exam starts in one second
	fmt.Println("Test starts in one second!!")
	time.Sleep(time.Duration(1) * time.Second)

	//started asking questions through ask function which is another go routine
	go ask(&records, &marks, &total, ch)

	//timer starts
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	//this blocks until the user completes the test or testduration completes
	select {
	case <-timer.C:
	case <-ch:
	}

	//prints score
	fmt.Printf("\n%v out of %v are correct\n", marks, total)
}
