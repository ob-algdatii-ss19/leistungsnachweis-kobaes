// Running this Benchmark with (within the same directory as main.go):
// $ go test -bench=.

package main

import (
	bufio "bufio"
	csv "encoding/csv"
	fmt "fmt"
	assert "github.com/stretchr/testify/assert"
	io "io"
	// math "math"
	os "os"
	strconv "strconv"
	testing "testing"
)

func TestAddItemSucceeds(t *testing.T) {
	k := knapsack{
		items:              make([]item, 0),
		totalWorth:         0,
		currentItemsVolume: 0,
		maxVolume:          10,
	}

	it := item{
		name:   "Apple",
		volume: 2,
		worth:  20,
	}

	err := k.addItem(it)

	if err != nil {
		t.Errorf("Expected success, but got an error: %v", err)
	}
}

func TestAddItemFails(t *testing.T) {
	k := knapsack{
		items:              make([]item, 0),
		totalWorth:         0,
		currentItemsVolume: 0,
		maxVolume:          1,
	}

	it := item{
		name:   "Apple",
		volume: 2,
		worth:  20,
	}

	err := k.addItem(it)

	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

func TestGreedySucceeds(t *testing.T) {
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

	k := knapsack{
		items:              make([]item, 0),
		totalWorth:         0,
		currentItemsVolume: 0,
		maxVolume:          10,
	}

	greedy(initItems, &k)
	assert.NotEmpty(t, k.currentItemsVolume)
	assert.NotEmpty(t, k.items)
	assert.NotEmpty(t, k.totalWorth)
}

func TestGreedyFails(t *testing.T) {
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

	k := knapsack{
		items:              make([]item, 0),
		totalWorth:         0,
		currentItemsVolume: 0,
		maxVolume:          0,
	}

	greedy(initItems, &k)
	assert.Empty(t, k.currentItemsVolume)
	assert.Empty(t, k.items)
	assert.Empty(t, k.totalWorth)
}

func TestDynamicSucceeds(t *testing.T) {
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

	k := knapsack{
		items:              make([]item, 0),
		totalWorth:         0,
		currentItemsVolume: 0,
		maxVolume:          10,
	}

	dynamic(initItems, &k)
	assert.NotEmpty(t, k.currentItemsVolume)
	assert.NotEmpty(t, k.items)
	assert.NotEmpty(t, k.totalWorth)
}

func TestDynamicFails(t *testing.T) {
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

	k := knapsack{
		items:              make([]item, 0),
		totalWorth:         0,
		currentItemsVolume: 0,
		maxVolume:          0,
	}

	dynamic(initItems, &k)
	assert.Empty(t, k.currentItemsVolume)
	assert.Empty(t, k.items)
	assert.Empty(t, k.totalWorth)
}

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
