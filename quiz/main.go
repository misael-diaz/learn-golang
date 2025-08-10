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
	fmt.Println(records);
	file.Close();
}
