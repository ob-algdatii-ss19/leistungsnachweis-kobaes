package main

import (
	bufio "bufio"
	csv "encoding/csv"
	"errors"
	"fmt"
	io "io"
	"math"
	os "os"
	"sort"
	strconv "strconv"
	"sync"
)

var wg sync.WaitGroup

type item struct {
	name          string
	volume, worth int
}

type knapsack struct {
	items                                     []item
	totalWorth, currentItemsVolume, maxVolume int
}

type knapsackParallel struct {
	items                                     *[]item
	totalWorth, currentItemsVolume, maxVolume int
}

func (k *knapsack) addItem(i item) error {
	if k.currentItemsVolume+i.volume <= k.maxVolume {
		k.currentItemsVolume += i.volume
		k.totalWorth += i.worth
		k.items = append(k.items, i)
		return nil
	}
	return errors.New("item too big!")
}

func (k *knapsackParallel) addItemParallel(i item) error {
	if k.currentItemsVolume+i.volume <= k.maxVolume {
		k.currentItemsVolume += i.volume
		k.totalWorth += i.worth
		*k.items = append(*k.items, i)
		return nil
	}
	return errors.New("item too big!")
}

func greedy(is []item, k *knapsack) {
	sort.Slice(is, func(i, j int) bool {
		return (is[i].worth / is[i].volume) > (is[j].worth / is[j].volume)
	})
	for i := range is {
		k.addItem(is[i])
	}
}

func checkItem(k *knapsack, i int, j int, is []item, matrix [][]int) {
	if i <= 0 || j <= 0 {
		return
	}

	pick := matrix[i][j]
	if pick != matrix[i-1][j] {
		k.addItem(is[i-1])
		checkItem(k, i-1, j-is[i-1].volume, is, matrix)
	} else {
		checkItem(k, i-1, j, is, matrix)
	}
}

func dynamic(is []item, k *knapsack) *knapsack {
	numItems := len(is) // number of items in knapsack
	capacity := k.maxVolume

	// create the empty matrix
	matrix := make([][]int, numItems+1) // items
	for i := range matrix {
		matrix[i] = make([]int, capacity+1) // volumes
	}

	for i := 1; i <= numItems; i++ {
		for j := 1; j <= capacity; j++ {
			if is[i-1].volume <= j {
				valueOne := float64(matrix[i-1][j])
				valueTwo := float64(is[i-1].worth + matrix[i-1][j-is[i-1].volume])
				matrix[i][j] = int(math.Max(valueOne, valueTwo))
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	checkItem(k, numItems, capacity, is, matrix)
	k.totalWorth = matrix[numItems][capacity]
	//k.totalWeight = k.currentItemsVolume

	return k
}

///////////////////////////// Parallel computation ---------------------------------

// CheckItem for concurrent computation
func checkItemParallel(k *knapsackParallel, i int, j int, is []item, matrix [][]int) {
	if i <= 0 || j <= 0 {
		wg.Done()
		return
	}

	pick := matrix[i][j]
	fmt.Printf("\n pick is: %v", pick)
	fmt.Printf("\n matrix[i-1][j] is: %v", matrix[i-1][j])
	if pick != matrix[i-1][j] {
		k.addItemParallel(is[i-1])
		wg.Add(1) // for starting a new go routine
		go checkItemParallel(k, i-1, j-is[i-1].volume, is, matrix)
	} else {
		wg.Add(1) // for starting a new go routine
		go checkItemParallel(k, i-1, j, is, matrix)
	}
}

func dynamicParallel(is []item, k *knapsackParallel) *knapsackParallel {
	numItems := len(is) // number of items in knapsack
	capacity := k.maxVolume

	// create the empty matrix
	matrix := make([][]int, numItems+1) // items
	for i := range matrix {
		matrix[i] = make([]int, capacity+1) // volumes
	}

	for i := 1; i <= numItems; i++ {
		for j := 1; j <= capacity; j++ {
			if is[i-1].volume <= j {
				valueOne := float64(matrix[i-1][j])
				valueTwo := float64(is[i-1].worth + matrix[i-1][j-is[i-1].volume])
				matrix[i][j] = int(math.Max(valueOne, valueTwo))
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	checkItemParallel(k, numItems, capacity, is, matrix)
	k.totalWorth = matrix[numItems][capacity]

	// Waiting for all go routines to end
	wg.Wait()

	return k
}

// Function to easily print a matrix
func printMatrix(mat [][]int, length int) {
	for i, outer := range mat {
		for j := range outer {
			if j == length {
				fmt.Printf("%v\n", mat[i][j])
			} else {
				fmt.Printf("%v\t", mat[i][j])
			}
		}
	}
}

func main() {
	csvFile, err := os.Open("config.csv") // Opens a file for read only

	// If no config file could be opened an error occurs and the program is ended
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var items []item

	for {
		line, err := reader.Read() // read single lines from the .csv file
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
		}

		v, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("[ERROR] something went wrong while reading \"volume\" as int: %v\n", err)
		}

		w, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Printf("[ERROR] something went wrong while reading \"worth\" as int: %v\n", err)
		}
		items = append(items, item{
			name:   line[0],
			volume: v,
			worth:  w,
		})
	}

	fmt.Printf("[INFO] result of reading the .csv file: %v\n", items)

	initItems := []item{
		item{name: "Apple", volume: 3, worth: 30},
		item{name: "Apple", volume: 3, worth: 30},
		item{name: "Orange", volume: 4, worth: 30},
		item{name: "Orange", volume: 4, worth: 30},
		item{name: "Pencil", volume: 1, worth: 10},
		item{name: "Pencil", volume: 1, worth: 10},
		item{name: "Pencil", volume: 1, worth: 10},
		item{name: "Mirror", volume: 5, worth: 40},
		item{name: "Mirror", volume: 5, worth: 40},
	}

	kg := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	kd := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	// initializing a knapsack parallel struct which contains pointer to
	// itemList := make([]item, 0)
	// kdp := knapsackParallel{items: &itemList, totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	// GREEDY Algorithm ------------------------------------------------------------------------------
	greedy(initItems, &kg)
	fmt.Println("Greedy Algorithm:")
	resultg := ""
	resultgWorth := kg.totalWorth
	for _, it := range kg.items {
		resultg += it.name + " "
	}
	fmt.Println(resultg)
	fmt.Println("Total Worth: " + strconv.Itoa(resultgWorth))
	fmt.Println("Total Volume: " + strconv.Itoa(kg.currentItemsVolume))

	// DYNAMIC Algorithm -----------------------------------------------------------------------------
	dynamic(initItems, &kd)

	fmt.Println()
	fmt.Println("Dynamic Algorithm:")
	resultd := ""
	resultdWorth := kd.totalWorth

	for _, it := range kd.items {
		resultd += it.name + " "
	}

	fmt.Println(resultd)
	fmt.Println("Total Worth: " + strconv.Itoa(resultdWorth))
	fmt.Println("Total Volume: " + strconv.Itoa(kd.currentItemsVolume))

	// DYNAMIC PARALLEL Algorithm --------------------------------------------------------------------
	// dynamicParallel(initItems, &kdp)

	// fmt.Println()
	// fmt.Println("Dynamic Algorithm in parallel:")
	// resultdp := ""
	// resultdpWorth := kdp.totalWorth

	// for _, it := range *kdp.items {
	// 	fmt.Println(it)
	// 	resultdp += it.name + " "
	// }

	// fmt.Println(resultdp)
	// fmt.Println("Total Worth parallel: " + strconv.Itoa(resultdpWorth))
	// fmt.Println("Total Volume parallel: " + strconv.Itoa(kdp.currentItemsVolume))
}
