package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
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

func (k *knapsack) addItem(i item) error {
	if k.currentItemsVolume+i.volume <= k.maxVolume {
		k.currentItemsVolume += i.volume
		k.totalWorth += i.worth
		k.items = append(k.items, i)
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

// CheckItem for concurrent computation
func checkItemParallel(k *knapsack, i int, j int, is []item, matrix [][]int) {
	if i <= 0 || j <= 0 {
		return
	}

	wg.Add(1) // for starting a new go routine

	pick := matrix[i][j]
	if pick != matrix[i-1][j] {
		k.addItem(is[i-1])
		go checkItemParallel(k, i-1, j-is[i-1].volume, is, matrix)
	} else {
		go checkItemParallel(k, i-1, j, is, matrix)
	}

	wg.Done()
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

	// Printing the matrix just to know what happens
	// fmt.Println("Matrix is:")
	// printMatrix(matrix, capacity)

	checkItem(k, numItems, capacity, is, matrix)
	k.totalWorth = matrix[numItems][capacity]
	//k.totalWeight = k.currentItemsVolume

	return k
}

func dynamicParallel(is []item, k *knapsack) *knapsack {
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

	kdPar := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

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

	// Dynamic computation starts here:
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

	// Dynamic computation with parallel checkItems function:
	dynamic(initItems, &kdPar)

	fmt.Println()
	fmt.Println("Dynamic Algorithm in parallel:")
	resultd = ""
	resultdWorth = kdPar.totalWorth

	for _, it := range kdPar.items {
		resultd += it.name + " "
	}

	fmt.Println(resultd)
	fmt.Println("Total Worth parallel: " + strconv.Itoa(resultdWorth))
	fmt.Println("Total Volume parallel: " + strconv.Itoa(kdPar.currentItemsVolume))
}
