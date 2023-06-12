package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func counter(fileName string) (elfs_cals [][]int) {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	elfs_cals = append(elfs_cals, []int{})
	for scanner.Scan() {
		if scanner.Text() != "" {
			scan_conv, err := strconv.Atoi(scanner.Text())
			check(err)
			elfs_cals[len(elfs_cals)-1] = append([]int(elfs_cals[len(elfs_cals)-1]), scan_conv)
		} else {
			elfs_cals = append(elfs_cals, []int{})
		}
	}
	return
}

func add_grouped_cals(array [][]int) (grouped []int) {
    for _, v := range array {
        temp := 0
        for _, va := range v {
            temp += va
        }
        grouped = append(grouped, temp)
    }
    return
}

func findNumber(array []int, needle int) (x int, newArray []int) {
	for i, v := range array {
		if v == needle {
			temp := make([]int, 0)
			temp = append(temp, array[:i]...)
			temp = append(temp, array[i+1:]...)
			return i, temp
		}
	}
	return -1, []int{}
}

func eat(array [][]int, index int) [][]int {
	first_dimension := array[index]
	sort.Slice(first_dimension, func(i, j int) bool {
		return first_dimension[i] > first_dimension[j]
	})
	_, temp := findNumber(array[index], first_dimension[0])
	array[index] = temp
	return array
}

func main() {
	elf_cals := counter("./calories") // i don't need the ./ here but might as well
	unsorted_cals := make([]int, 0)
	for i := range elf_cals {
		for _, va := range elf_cals[i] {
			unsorted_cals = append(unsorted_cals, va)
		}
	}
	for len(unsorted_cals) > 0 {
		sorted_cals := make([]int, len(unsorted_cals))
		copy(sorted_cals, unsorted_cals)
		sort.Slice(sorted_cals, func(i, j int) bool {
			return sorted_cals[i] > sorted_cals[j]
		})
        added_cals := add_grouped_cals(elf_cals)
        sorted_added_cals := make([]int, len(added_cals))
        copy(sorted_added_cals, added_cals)
        // fmt.Println(sorted_added_cals)
        sort.Slice(sorted_added_cals, func(i, j int) bool {
            return sorted_added_cals[i] > sorted_added_cals[j]
        })
        fmt.Println(sorted_added_cals)
		index, _ := findNumber(added_cals, sorted_added_cals[0])
		unsorted_cals = nil
		elf_cals = eat(elf_cals, index)
		for _, v := range elf_cals {
			for _, va := range v {
				unsorted_cals = append(unsorted_cals, va)
			}
		}
	}
}
