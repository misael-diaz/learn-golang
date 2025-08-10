package main;
import (
	"os"
	"log"
	"fmt"
	"bufio"
	"encoding/csv"
	"strings"
	"strconv"
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
	ncorrect := 0;
	nproblems := len(records);
	for _, ln := range records {
		problem_type, problem_statement, problem_answer := ln[0], ln[1], ln[2];
		fmt.Printf("prompt: %s\n", problem_statement);
		response, err := console_reader.ReadString('\n');
		res := strings.TrimSpace(response);
		if (nil != err) {
			log.Fatal(err);
		}
		if (0 == strings.Compare("text", problem_type)) {
			if (0 == strings.Compare(res, problem_answer)) {
				ncorrect++;
			}
		} else if (0 == strings.Compare("math", problem_type)) {
			ans, _ := strconv.Atoi(problem_answer);
			num, _ := strconv.Atoi(res);
			if (ans == num) {
				ncorrect++;
			}
		} else {
			log.Fatal("UXProblemTypeError");
		}
	}
	fmt.Printf("total number of problems: %d\n", nproblems);
	fmt.Printf("total number of correct answers: %d\n", ncorrect);
	file.Close();
}
