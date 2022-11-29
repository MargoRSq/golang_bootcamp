package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	fgs := initFlags()
	arr := getInput()
	sort.Float64s(arr)
	output(arr, fgs)
}

type flags struct {
	meanFlag *bool
	modeFlag *bool
	medFlag  *bool
	sdFlag   *bool
}

func output(arr []float64, fgs flags) {
	if *fgs.meanFlag {
		fmt.Printf("Mean: %.2f\n", CalcMean(arr))
	}
	if *fgs.medFlag {
		fmt.Printf("Median: %.2f\n", CalcMedian(arr))
	}
	if *fgs.modeFlag {
		fmt.Println("Mode: ", FindMode(arr))
	}
	if *fgs.sdFlag {
		fmt.Printf("SD: %.2f\n", CalcSD(arr))
	}
}

func initFlags() (fgs flags) {
	fgs.meanFlag = flag.Bool("mean", false, "use -mean to calculate mean")
	fgs.medFlag = flag.Bool("median", false, "use -median to calculate median")
	fgs.modeFlag = flag.Bool("mode", false, "use -mode to find mode")
	fgs.sdFlag = flag.Bool("sd", false, "use -sd to calculate standard deviation")
	flag.Parse()
	if flag.NFlag() == 0 {
		*fgs.meanFlag = true
		*fgs.medFlag = true
		*fgs.modeFlag = true
		*fgs.sdFlag = true
	}
	return
}

func getInput() []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	amount, err := strconv.Atoi(scanner.Text())
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount format, err:", err)
		os.Exit(1)
	}

	numbers := make([]float64, amount)
	for i := 0; i < amount; {
		scanner.Scan()
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input")
		} else {
			if f > -100000 && f < 100000 {
				numbers[i] = f
			} else {
				fmt.Println("Invalid number")
			}
			i++
		}
	}
	return numbers
}

func CalcMean(arr []float64) (mean float64) {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	mean = sum / float64(len(arr))
	return
}

func CalcMedian(arr []float64) (median float64) {
	arrLen := len(arr)
	if arrLen != 0 {
		if arrLen%2 == 0 {
			median = (arr[(arrLen/2)-1] + arr[arrLen/2]) / 2
		} else {
			median = arr[arrLen/2]
		}
	} else {
		median = 0
	}
	return
}

func FindMode(arr []float64) (mode int) {
	countMap := make(map[float64]int)
	for _, value := range arr {
		countMap[value] += 1
	}

	max := 0
	for _, key := range arr {
		count := countMap[key]
		if count > max {
			max = count
			mode = int(key)
		}
	}
	return
}

func CalcSD(arr []float64) (sd float64) {
	mean := CalcMean(arr)
	for _, value := range arr {
		sd += math.Pow(value-mean, 2)
	}
	sd = math.Sqrt(sd / float64(len(arr)))
	return
}
