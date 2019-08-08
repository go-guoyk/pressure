package main

import (
	"flag"
	"fmt"
	"os"
)

type measure struct {
	name      string
	threshold int
	measure   func() (int, error)
}

func main() {
	flag.IntVar(&measureCPU.threshold, "cpu", 80, "CPU pressure threshold, in percentage")
	flag.IntVar(&measureMEM.threshold, "mem", 80, "MEM pressure threshold, in percentage")
	flag.Parse()

	measures := []measure{measureCPU, measureMEM}

	var fail bool

	for _, m := range measures {
		var val int
		var err error
		if val, err = m.measure(); err != nil {
			fail = true
			fmt.Printf("%s: %s; ", m.name, err.Error())
			continue
		}
		if val >= m.threshold {
			fail = true
		}
		fmt.Printf("%s: %d%%; ", m.name, val)
	}

	if fail {
		fmt.Print("FAIL\n")
		os.Exit(1)
	} else {
		fmt.Print("PASS\n")
	}
}
