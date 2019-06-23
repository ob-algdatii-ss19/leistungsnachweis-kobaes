// Running this Benchmark with (within the same directory as main.go):
// $ go test -bench=.

package main

import (
	bufio "bufio"
	csv "encoding/csv"
	fmt "fmt"
	io "io"
	os "os"
	strconv "strconv"
	testing "testing"
)

func BenchmarkDynamic(b *testing.B) {

	// reading a .csv file with items for testing purposes
	csvFile, err := os.Open("configs/configtest.csv")

	// if no config file could be opened an error occurs and the program is ended
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var initItems []item
	for {
		line, err := reader.Read() // read single lines from the .csv file

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
		}

		if line[0] == "#" {
			continue
		}

		v, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("[ERROR] something went wrong while reading \"volume\" as int: %v\n", err)
		}

		w, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Printf("[ERROR] something went wrong while reading \"worth\" as int: %v\n", err)
		}
		initItems = append(initItems, item{
			name:   line[0],
			volume: v,
			worth:  w,
		})
	}

	kd := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	for n := 0; n < b.N; n++ {
		dynamic(initItems, &kd)
	}
}
