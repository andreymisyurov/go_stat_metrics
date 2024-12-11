package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	txt := sc.Text()
	parts := strings.Fields(txt)
	var nums []int
	var sum float64

	for i := 0; i < len(parts); i++ {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s - is not a number\n", parts[i])
			os.Exit(1)
		}
		sum += float64(num)
		nums = append(nums, num)
	}
	fmt.Println(sum / float64(len(nums)))
}









