package main

import (
	//"sync"
	"fmt"
	"github.com/swgloomy/go-common"
	"math"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	find := float64(3) / float64(13)

	fmt.Println(int(find * 100))

	fmt.Println(fmt.Sprintf("【智能营销】您好：到达率%.2f%%", common.RoundingPercentageByInt(77, 717, 2)))
	return

	go func() {
		for {
			fmt.Println(3)
			return
		}
		fmt.Println(4)
	}()

	fmt.Println(2)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	//信号等待
	<-c
}

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
