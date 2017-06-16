package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type record struct {
	Debit   string
	Remarks string
	Date    string
}

func newRecord(row []string, x, y, z int) record {
	//println(row[y])
	//debit, _ := strconv.ParseFloat(row[y], 64)
	return record{
		Debit:   row[y],
		Remarks: row[x],
		Date:    row[z],
	}
}

func main() {

	f, err := os.Open("myacct.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Date        |      Debit   |   Remarks")

	for i, row := range rows {
		if i < 19 {
			continue
		} else if !(strings.Contains(row[6], "UCHE") && strings.Contains(row[6], "SALARY")) {
			continue
		}
		record := newRecord(row, 6, 3, 0)
		fmt.Printf("%s  |   %s  |  %s\n", record.Date, record.Debit, record.Remarks)

		//openValues = append(openValues, fmt.Sprintf("%.2f", record.Open))

	}

}
