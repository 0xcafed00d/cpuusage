package main

import (
	"fmt"
	"os"
	"time"

	"github.com/simulatedsimian/cpuusage"
)

func main() {
	u := cpuusage.Usage{}

	for {
		err := u.GetUsage()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Overall: %d Cores: %v\n", u.Overall, u.Cores)
		time.Sleep(1 * time.Second)
	}
}
