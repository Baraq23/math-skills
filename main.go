package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"mathskills/stats" // Package_name/Folder_name
)

func main() {
	// checking if input is correctly entered
	arguments := len(os.Args)
	if arguments != 2 {
		fmt.Println("Not enough arguements")
		return
	} else if os.Args[1] != "data.txt" {
		fmt.Println("Input should be: go run program-name.go data.txt")
		return
	}
	// open and read data file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []float64{}
	lineNum := 0
	nilData := []int{}
	for scanner.Scan() {
		lineNum++
		str := scanner.Text()
		if str == "" {
			nilData = append(nilData, lineNum)
			continue
		}
		// Removing trailing or leading spaces
		str = strings.TrimSpace(str)
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("%v from line %v is an invalid representation of a number\n", str, lineNum)
			os.Exit(1)
		}
		// Number bigger than the maximum float64 cappacity will not be computed
		if num > math.MaxInt64 {
			fmt.Printf("Number at line %v in data file is too big to compute!\n", lineNum)
			os.Exit(1)
		}
		nums = append(nums, float64(num))
	}
	// this check if the data file has values
	if len(nums) == 0 {
		fmt.Println("Empty data file found.")
		os.Exit(1)
	}

	// Printing results
	fmt.Println("Average: ", int(math.Round(mathskills.Average(nums))))
	fmt.Println("Median: ", int(math.Round(mathskills.Median(nums))))
	fmt.Println("Variance: ", int(math.Round(mathskills.Variance(nums))))
	fmt.Println("Standard Deviation: ", int(math.Round(mathskills.SD(nums))))

	if len(nilData) != 0 {
		fmt.Println()
		for i := 0; i < 69; i++ {
			fmt.Print("=")
		}
		fmt.Println()
		fmt.Println("Please note, data was not found in the following lines in the data file:")
		fmt.Println(nilData)
		fmt.Println("Therefore the results shown are only for the data as is.")
	}
}
