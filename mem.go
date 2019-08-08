package main

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func measureMEM() (pressure int, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile("/proc/meminfo"); err != nil {
		return
	}

	s := bufio.NewScanner(bytes.NewReader(buf))
	s.Split(bufio.ScanLines)

	var total, avail uint64
	num := 2

	for s.Scan() {
		l := s.Text()
		parts := whitespaces.Split(strings.TrimSpace(l), -1)
		if len(parts) != 3 {
			continue
		}
		if parts[0] == "MemTotal:" {
			if err = decodeMeminfoLine(parts, &total); err != nil {
				return
			}
			num--
		}
		if parts[0] == "MemAvailable:" {
			if err = decodeMeminfoLine(parts, &avail); err != nil {
				return
			}
			num--
		}
	}

	if num != 0 {
		err = errors.New("missing MemTotal or MemAvailable in meminfo")
		return
	}
	if avail > total || total == 0 {
		err = errors.New("invalid value in meminfo")
		return
	}
	pressure = int(float64(total-avail) * 100 / float64(total))
	return
}

func decodeMeminfoLine(parts []string, out *uint64) (err error) {
	if parts[2] != "kB" {
		err = errors.New("meminfo line is not ended in 'kB'")
		return
	}
	if *out, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
		return
	}
	return
}

