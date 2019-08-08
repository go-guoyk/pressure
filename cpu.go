package main

import (
	"errors"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

func measureCPU() (pressure int, err error) {
	var buf []byte

	if buf, err = ioutil.ReadFile("/proc/loadavg"); err != nil {
		return
	}

	las := whitespaces.Split(strings.TrimSpace(string(buf)), -1)
	if len(las) != 5 {
		err = errors.New("invalid loadavg")
		return
	}

	var la float64
	if la, err = strconv.ParseFloat(strings.TrimSpace(las[0]), 64); err != nil {
		return
	}

	pressure = int(la * 100 / float64(runtime.NumCPU()))

	return
}

