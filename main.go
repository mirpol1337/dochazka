package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {
	xlFile, err := xlsx.OpenFile("Dochazka.xlsx")
	if err != nil {
		panic(err)
	}

	var durations []time.Duration

	for i, row := range xlFile.Sheets[0].Rows {
		index := 0
		if row.Cells[2].Value == "příchod" {
			startTime := row.Cells[0].Value
			start, err := time.Parse(layout, startTime)

			if err != nil {
				return
			}

			if i+1 < len(xlFile.Sheets[0].Rows) && xlFile.Sheets[0].Rows[i+1].Cells[2].Value == "odchod" {
				nextRow := xlFile.Sheets[0].Rows[i+1]
				endTime := nextRow.Cells[0].Value
				end, err := time.Parse(layout, endTime)

				if err != nil {
					return
				}

				duration := end.Sub(start)
				durations = append(durations, duration)
				fmt.Println(duration.String())
				index++
			}
		}
	}

	sum := time.Duration(0)
	for _, d := range durations {
		sum += d
	}

	avg := sum / time.Duration(len(durations))
	fmt.Println("Average time:", avg.String())

	expectedWorkHours := time.Duration(float64(len(durations)) * 8.5 * float64(time.Hour))
	totalWorkHours := sum.Round(time.Minute)
	overtime := totalWorkHours - expectedWorkHours

	if overtime >= 0 {
		fmt.Println("Overtime:", overtime.String())
	} else {
		fmt.Println("Undertime:", (-overtime).String())
	}
}

const (
	layout = "2.1.2006 15:04:05"
)
