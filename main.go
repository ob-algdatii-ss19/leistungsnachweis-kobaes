package main

import (
		"fmt"
		"sort"
		"errors"
)

type item struct {
	name   string
	volume, worth  int
}

type knapsack struct {
	items       []item
	totalWorth, currentItemsVolume, maxVolume  int
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

func main() {
	initItems := []item{
		item{name: "Apple", volume: 3, worth: 20},
		item{name: "Orange", volume: 4, worth: 30},
		item{name: "Pencil", volume: 1, worth: 10},
		item{name: "Mirror", volume: 5, worth: 40},
	}

	k := knapsack{items: make([]item, 0), totalWorth: 0, currentItemsVolume: 0, maxVolume: 10}

	greedy(initItems, &k)
	for _, it := range k.items {
		fmt.Println(it.name)
	}

}
