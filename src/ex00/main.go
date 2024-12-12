package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
	"math"
)

func input_num() []float64 {
	sc := bufio.NewScanner(os.Stdin);
	sc.Scan();
	txt := sc.Text();
	var parts []string = strings.Fields(txt);
	var nums []float64;

	for i := 0; i < len(parts); i++ {
		var num, err = strconv.ParseFloat(parts[i], 64);
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s - is not a number\n", parts[i]);
			os.Exit(1);
		}
		nums = append(nums, num);
	}
	
	sort.Float64s(nums);
	return nums;
}

func get_sum(nums []float64) float64 {

	var sum float64;
	for i := 0; i < len(nums); i++ {
		sum += float64(nums[i]);
	}

	return sum;
}

func get_median(nums []float64) float64 {
	if len(nums) % 2 == 1 {
		return nums[len(nums) / 2];
	} else {
		return (nums[len(nums) / 2 - 1] + nums[len(nums) / 2]) / 2;
	}
}

func get_mean(sum float64, count int) float64 {
	return sum / float64(count);
}

func get_mode(nums []float64) float64 {
	var freq map[float64]int;
	freq = make(map[float64]int);
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++;
	}

	var count_max int = -1;
	var val_max float64;
	for num, count := range freq {
		if count > count_max || (count == count_max && val_max > num) {
			count_max = count;
			val_max = num;
		}
	}

	return val_max;
}

func get_sd(nums []float64) float64 {
	var mean float64 = get_mean(get_sum(nums), len(nums));
	var var_diff []float64 = append([]float64(nil), nums...);
	for i := 0; i < len(var_diff); i++ {
		var_diff[i] = math.Pow((mean - var_diff[i]), 2);
	}
	var sum float64 = get_sum(var_diff);
	return math.Sqrt(sum / float64(len(var_diff) - 1));
}

func main() {
	var arr []float64 = input_num();
	fmt.Printf("Mean: %.2f\n", get_mean(get_sum(arr), len(arr)));
	fmt.Printf("Median: %.2f\n", get_median(arr));
	fmt.Printf("Mode: %.2f\n", get_mode(arr));
	fmt.Printf("SD: %.2f\n", get_sd(arr));
}









