// Running this Benchmark with (within the same directory as main.go):
// $ go test -bench=.

package main

import (
	"testing"
)

func BenchmarkDynamic(b *testing.B) {
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

	kd := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	for n := 0; n < b.N; n++ {
		dynamic(initItems, &kd)
	}
}

// func BenchmarkDynamicParallel(b *testing.B) {
// 	initItems := []item{
// 		item{name: "Apple", volume: 3, worth: 30},
// 		item{name: "Apple", volume: 3, worth: 30},
// 		item{name: "Orange", volume: 4, worth: 30},
// 		item{name: "Orange", volume: 4, worth: 30},
// 		item{name: "Pencil", volume: 1, worth: 10},
// 		item{name: "Pencil", volume: 1, worth: 10},
// 		item{name: "Pencil", volume: 1, worth: 10},
// 		item{name: "Mirror", volume: 5, worth: 40},
// 		item{name: "Mirror", volume: 5, worth: 40},
// 	}

// 	itemList := make([]item, 0)

// 	kd := knapsackParallel{items: &itemList, totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

// 	for n := 0; n < b.N; n++ {
// 		dynamicParallel(initItems, &kd)
// 	}
// }
