package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Readline() string {
	fmt.Print("Array of numbers (space separated) to sort: ")
	if scanner := bufio.NewScanner(os.Stdin); scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func ParseLine(line string, nums *[]int) error {
	fields := strings.FieldsFunc(line, Split)
	if len(fields) < 1 {
		return errors.New("Input at least several numbers, example: 56 34 12")
	}
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return fmt.Errorf("Not a number: %s", field)
		}
		*nums = append(*nums, num)
	}
	fmt.Printf("Parsed array of numbers: %v\n", *nums)
	return nil
}

func GetSplitSizes(n int, parts int) []int {
	dividend := n / parts
	remainder := n - parts*dividend
	sizes := make([]int, 0)
	for i := 0; i < parts; i++ {
		if dividend == 0 && remainder == 0 {
			break
		}
		sizes = append(sizes, dividend)
		if remainder > 0 {
			sizes[i]++
			remainder--
		}
	}
	fmt.Printf("Split sizes: %v\n", sizes)
	return sizes
}

func SortNums(nums []int, ch chan []int) {
	sort.Ints(nums)
	fmt.Printf("Sorted split: %v\n", nums)
	ch <- nums
}

func MergeNums(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func main() {
	for {
		line := Readline()
		if line == "" {
			break
		}
		nums := make([]int, 0)
		err := ParseLine(line, &nums)
		if err != nil {
			fmt.Println(err)
			continue
		}

		sizes := GetSplitSizes(len(nums), 4)
		ch := make(chan []int, sizes[0])

		pos := 0
		for _, size := range sizes {
			go SortNums(nums[pos:pos+size], ch)
			pos += size
		}

		var result []int
		for i := 0; i < len(sizes); i++ {
			sorted := <-ch
			result = MergeNums(result, sorted)
		}
		fmt.Printf("Result: %v\n", result)
	}
}
