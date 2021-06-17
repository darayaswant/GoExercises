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

func ask(records *[][]string, t int) {
	total = len(*records)
	//timer starts
	timer := time.NewTimer(time.Duration(t) * time.Second)

	//if time ends, it prints score
	//otherwise it keeps asking the questions

	for _, record := range *records {
		ch := make(chan string)
		//creating go routine because while scanning, go routine stops
		go func() {
			fmt.Println(record[0])
			var temp string
			fmt.Scan(&temp)
			ch <- temp
		}()

		//this select statement chooses whether to end test or continue it..
		select {
		case <-timer.C:
			fmt.Printf("\n %v out of %v are correct\n", marks, total)
			return
		case ans := <-ch:

			//TrimSpace because in records there can be trailing and leading spaces which combines take the string
			//but while scanning the answer we won't consider spaces... so we trim the spaces from ans in record
			if ans == strings.TrimSpace(record[1]) {
				marks++
			}
		}
	}

	// this will only execute when the test completes before timelimit
	fmt.Printf("\n %v out of %v are correct\n", marks, total)
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

	//started asking questions through ask function
	ask(&records, *timelimit)
}
