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

//created marks and total as global variables, so that they can be accessed by all functions
var marks, total int = 0, 0

func ask(records *[][]string) {
	total = len(*records)

	//asking questions one by one, and taking answers from user
	for _, record := range *records {
		fmt.Println(record[0])
		var ans string
		fmt.Scan(&ans)

		//TrimSpace because in records there can be trailing and leading spaces which combines take the string
		//but while scanning the answer we won't consider spaces... so we trim the spaces from ans in record
		if ans == strings.TrimSpace(record[1]) {
			marks++
		}
	}
	//printing the score (this will execute when user completes the test before test duration)
	fmt.Printf("\n%v out of %v are correct\n", marks, total)

	//now program exits, as we printed score
	os.Exit(0)
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
	//Alerting the user that the exam starts in one second
	fmt.Println("Test starts in one second!!")
	time.Sleep(time.Duration(1) * time.Second)

	//started asking questions through ask function which is another go routine
	go ask(&records)

	//timer starts
	time.Sleep(time.Duration(*timelimit) * time.Second)

	//printing the score how much user got (this will execute when the test duration completes and user is still answering)
	fmt.Printf("\n%v out of %v are correct\n", marks, total)
}
