package main;
import (
	"os"
	"log"
	"fmt"
	"bufio"
	"encoding/csv"
);

func main() {
	file, err := os.Open("problems.csv");
	if (nil != err) {
		log.Fatal(err);
	}
	console_reader := bufio.NewReader(os.Stdin);
	r := csv.NewReader(file);
	records, err := r.ReadAll();
	if (nil != err) {
		log.Fatal(err);
	}
	nproblems := len(records);
	for _, ln := range records {
		problem_type, problem_statement, problem_answer := ln[0], ln[1], ln[2];
		fmt.Printf("prompt: %s\n", problem_statement);
		response, err := console_reader.ReadString('\n');
		if (nil != err) {
			log.Fatal(err);
		}
		fmt.Println(problem_type);
		fmt.Println(problem_answer);
		fmt.Println(response);
	}
	fmt.Println(nproblems);
	file.Close();
}
