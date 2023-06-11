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

func counter(fileName string) (calories_per_elf []int, elfs_cals [][]int) {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	calories_per_elf = make([]int, 1)
	elfs_cals = append(elfs_cals, []int{})

	for scanner.Scan() {
		if scanner.Text() != "" {
			scan_conv, err := strconv.Atoi(scanner.Text())
			check(err)
			calories_per_elf[len(calories_per_elf) - 1] += scan_conv
			elfs_cals[len(elfs_cals) - 1] = append([]int(elfs_cals[len(elfs_cals) - 1]), scan_conv)
		} else {
			calories_per_elf = append(calories_per_elf, 0)
			elfs_cals = append(elfs_cals, []int{})
		}
	}
	return
}

func findNumber(array []int, needle int) (x int, found bool) {
    for i, v := range array {
        fmt.Println(v)
        if v == needle {
            return i, true
        }
    }
    return -1, false
}

func main() {
	sorted_cals, elf_cals := counter("./calories") // i don't need the ./ here but might as well
    unsorted_cals := make([]int, len(sorted_cals))
    copy(unsorted_cals, sorted_cals)
	sort.Slice(sorted_cals, func(i, j int) bool {
		return sorted_cals[i] > sorted_cals[j]
	})
    index, _ := findNumber(unsorted_cals, sorted_cals[0])
	fmt.Println(unsorted_cals, elf_cals)
    fmt.Println(index)
}
