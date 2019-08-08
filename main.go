package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	thresholdMEM int
	thresholdCPU int
	verbose      bool
)

var (
	whitespaces = regexp.MustCompile(`\s+`)
)

type Measure struct {
	Name      string
	Threshold int
	Measure   func() (int, error)
}

func exit(err *error) {
	if *err != nil {
		fmt.Println((*err).Error())
		os.Exit(1)
	}
}

func main() {
	var err error
	defer exit(&err)

	flag.IntVar(&thresholdCPU, "cpu", 80, "CPU threshold, in percentage")
	flag.IntVar(&thresholdMEM, "mem", 80, "MEM threshold, in percentage")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	var cpuLoad, memLoad int
	if cpuLoad, err = measureCPU(); err != nil {
		return
	}
	if memLoad, err = measureMEM(); err != nil {
		return
	}

	var msg string
	var bad bool

	if cpuLoad >= thresholdCPU || memLoad >= thresholdMEM {
		bad = true
	}

	if cpuLoad >= thresholdCPU || verbose {
		msg += fmt.Sprintf("CPU Pressure: %d%%; ", cpuLoad)
	}
	if memLoad >= thresholdMEM || verbose {
		msg += fmt.Sprintf("MEM Pressure: %d%%; ", memLoad)
	}

	if len(msg) > 0 {
		fmt.Println(msg)
	}

	if bad {
		os.Exit(1)
	}
}

