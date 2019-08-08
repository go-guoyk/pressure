package main

import (
	"errors"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

var measureCPU = measure{
	name: "CPU",
	measure: func() (val int, err error) {
		var buf []byte

		if buf, err = ioutil.ReadFile("/proc/loadavg"); err != nil {
			return
		}

		var la float64
		if la, err = decodeLoadavg(buf); err != nil {
			return
		}

		val = int(la * 100 / float64(runtime.NumCPU()))

		return
	},
}

func decodeLoadavg(buf []byte) (float64, error) {
	las := whitespaces.Split(strings.TrimSpace(string(buf)), -1)
	if len(las) != 5 {
		return 0, errors.New("invalid loadavg data")
	}
	if la, err := strconv.ParseFloat(strings.TrimSpace(las[0]), 64); err != nil {
		return 0, errors.New("invalid loadavg number")
	} else {
		return la, nil
	}
}
