package main;
import (
	"os"
	"log"
	"fmt"
	"encoding/csv"
);

func main() {
	file, err := os.Open("problems.csv");
	if (nil != err) {
		log.Fatal(err);
	}
	r := csv.NewReader(file);
	records, err := r.ReadAll();
	if (nil != err) {
		log.Fatal(err);
	}
	nproblems := len(records);
	for _, ln := range records {
		fmt.Println(ln);
	}
	fmt.Println(nproblems);
	file.Close();
}
