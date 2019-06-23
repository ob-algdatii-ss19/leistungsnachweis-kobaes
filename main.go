package main

//nolint:goimports
import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	flag "github.com/ogier/pflag"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

//nolint:gochecknoglobals
var (
	flagConfigFile bool
	flagGreedy     bool
	flagDynamic    bool
	flagHelp       bool
	flagAll        bool
	// wg             sync.WaitGroup
)

//nolint:gochecknoinits
func init() {
	flag.BoolVar(&flagConfigFile, "configfile", false, "flag which specifies a path to a config file")
	flag.BoolVar(&flagGreedy, "greedy", false, "flag which specifies to use greedy algorithm")
	flag.BoolVar(&flagDynamic, "dynamic", false, "flag which specifies to use dynamic algorithm")
	flag.BoolVar(&flagAll, "all", false, "flag which specifies to use both greedy and dynamic algorithm")
	flag.BoolVar(&flagHelp, "help", false, "flag which specifies that help should be shown")
}

type item struct {
	name          string
	volume, worth int
}

type knapsack struct {
	items                                     []item
	totalWorth, currentItemsVolume, maxVolume int
}

// type knapsackParallel struct {
// 	items                                     *[]item
// 	totalWorth, currentItemsVolume, maxVolume int
// }

func (k *knapsack) addItem(i item) error {
	if k.currentItemsVolume+i.volume <= k.maxVolume {
		k.currentItemsVolume += i.volume
		k.totalWorth += i.worth
		k.items = append(k.items, i)
		return nil
	}
	return errors.New("item too big")
}

// func (k *knapsackParallel) addItemParallel(i item) error {
// 	if k.currentItemsVolume+i.volume <= k.maxVolume {
// 		k.currentItemsVolume += i.volume
// 		k.totalWorth += i.worth
// 		*k.items = append(*k.items, i)
// 		return nil
// 	}
// 	return errors.New("item too big!")
// }

func greedy(is []item, k *knapsack) {
	sort.Slice(is, func(i, j int) bool {
		return (is[i].worth / is[i].volume) > (is[j].worth / is[j].volume)
	})
	for i := range is {
		err := k.addItem(is[i])
		if err != nil {
			// fmt.Printf("[ERROR] %v", err)
			fmt.Printf("")

		}
	}
}

func checkItem(k *knapsack, i int, j int, is []item, matrix [][]int) {
	if i <= 0 || j <= 0 {
		return
	}

	pick := matrix[i][j]
	if pick != matrix[i-1][j] {
		err := k.addItem(is[i-1])
		if err != nil {
			// fmt.Printf("[ERROR] %v", err)
			fmt.Printf("")

		}
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
// func checkItemParallel(k *knapsackParallel, i int, j int, is []item, matrix [][]int) {
// 	if i <= 0 || j <= 0 {
// 		wg.Done()
// 		return
// 	}

// 	pick := matrix[i][j]
// 	fmt.Printf("\n pick is: %v", pick)
// 	fmt.Printf("\n matrix[i-1][j] is: %v", matrix[i-1][j])
// 	if pick != matrix[i-1][j] {
// 		k.addItemParallel(is[i-1])
// 		wg.Add(1) // for starting a new go routine
// 		go checkItemParallel(k, i-1, j-is[i-1].volume, is, matrix)
// 	} else {
// 		wg.Add(1) // for starting a new go routine
// 		go checkItemParallel(k, i-1, j, is, matrix)
// 	}
// }

// func dynamicParallel(is []item, k *knapsackParallel) *knapsackParallel {
// 	numItems := len(is) // number of items in knapsack
// 	capacity := k.maxVolume

// 	// create the empty matrix
// 	matrix := make([][]int, numItems+1) // items
// 	for i := range matrix {
// 		matrix[i] = make([]int, capacity+1) // volumes
// 	}

// 	for i := 1; i <= numItems; i++ {
// 		for j := 1; j <= capacity; j++ {
// 			if is[i-1].volume <= j {
// 				valueOne := float64(matrix[i-1][j])
// 				valueTwo := float64(is[i-1].worth + matrix[i-1][j-is[i-1].volume])
// 				matrix[i][j] = int(math.Max(valueOne, valueTwo))
// 			} else {
// 				matrix[i][j] = matrix[i-1][j]
// 			}
// 		}
// 	}

// 	checkItemParallel(k, numItems, capacity, is, matrix)
// 	k.totalWorth = matrix[numItems][capacity]

// 	// Waiting for all go routines to end
// 	wg.Wait()

// 	return k
// }

// Function to easily print a matrix
// func printMatrix(mat [][]int, length int) {
// 	for i, outer := range mat {
// 		for j := range outer {
// 			if j == length {
// 				fmt.Printf("%v\n", mat[i][j])
// 			} else {
// 				fmt.Printf("%v\t", mat[i][j])
// 			}
// 		}
// 	}
// }

func main() {
	var csvFile *os.File
	var err error

	// Flags ---------------------------------------------------------------------------
	// Parsing the flags from above
	flag.Parse()

	// Stores all arguments of commandline after the flags in arguments
	arguments := flag.Args()

	if flagHelp {
		fmt.Println("HELP")
		fmt.Println("The following flags are available:")
		fmt.Println("--help:\t\t\t\t\t\tshows this help message")
		fmt.Println("--greedy:\t\t\t\t\tcomputing knapsack using the greedy algorithm")
		fmt.Println("--dynamic:\t\t\t\t\tcomputing knapsack using the dynamic algorithm")
		fmt.Println("--configfile \"path-to-file/filename.csv\":\tusing own .csv file for specifying own items")
		os.Exit(0)
	}

	if flagConfigFile {
		csvFile, err = os.Open(arguments[0]) // Opens a file for read only
		// If no config file could be opened an error occurs and the program is ended
		if err != nil {
			fmt.Printf("[ERROR] %v", err)
			os.Exit(1)
		}
	} else {
		csvFile, err = os.Open("config.csv") // Opens a file for read only
		// If no config file could be opened an error occurs and the program is ended
		if err != nil {
			fmt.Printf("[ERROR] %v", err)
			os.Exit(1)
		}
	}

	// reading and parsing config file for items
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


	fmt.Println("Using the following items for computing knapsack:")
	for _, i := range initItems {
		fmt.Printf("Item: %v, Volume: %v, Worth: %v\n", i.name, i.volume, i.worth)
	}
	fmt.Println()

	// initItems := []item{
	// 	item{name: "Apple", volume: 3, worth: 30},
	// 	item{name: "Apple", volume: 3, worth: 30},
	// 	item{name: "Orange", volume: 4, worth: 30},
	// 	item{name: "Orange", volume: 4, worth: 30},
	// 	item{name: "Pencil", volume: 1, worth: 10},
	// 	item{name: "Pencil", volume: 1, worth: 10},
	// 	item{name: "Pencil", volume: 1, worth: 10},
	// 	item{name: "Mirror", volume: 5, worth: 40},
	// 	item{name: "Mirror", volume: 5, worth: 40},
	// }

	kg := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	kd := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	// initializing a knapsack parallel struct which contains pointer to
	// itemList := make([]item, 0)
	// kdp := knapsackParallel{items: &itemList, totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	// GREEDY Algorithm ------------------------------------------------------------------------------
	if flagGreedy || flagAll {
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
	}

	// DYNAMIC Algorithm -----------------------------------------------------------------------------
	if flagDynamic || flagAll {
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
	}

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
