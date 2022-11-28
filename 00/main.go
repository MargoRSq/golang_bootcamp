package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := getInput()
	sort.Float64s(arr)
	fgs := initFlags()
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
		fmt.Printf("Median: %.2f\n", CalcSD(arr))
	}
}

func initFlags() (fgs flags) {
	fgs.meanFlag = flag.Bool("mean", false, "-mean")
	fgs.medFlag = flag.Bool("median", false, "-median")
	fgs.modeFlag = flag.Bool("mode", false, "-mode")
	fgs.sdFlag = flag.Bool("sd", false, "-sd")
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
	var amount int
	fmt.Scan(&amount)

	numbers := make([]float64, amount)
	var tmp float64
	for i := 0; i < amount; i++ {
		_, err := fmt.Scan(&tmp)
		if err != nil {
			fmt.Println("Invalid input")
		} else {
			if tmp > -100000 && tmp < 100000 {
				numbers[i] = tmp
			} else {
				fmt.Println("Invalid number")
			}
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
			median = arr[(arrLen/2)-1] + arr[arrLen/2]
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
