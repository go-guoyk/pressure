package main

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	measureMEM = measure{
		name: "MEM",
		measure: func() (val int, err error) {
			var buf []byte
			if buf, err = ioutil.ReadFile("/proc/meminfo"); err != nil {
				return
			}

			var total, avail uint64
			if total, avail, err = decodeMeminfo(buf); err != nil {
				return
			}
			val = int(float64(total-avail) * 100 / float64(total))
			return
		},
	}
)

func decodeMeminfo(buf []byte) (total uint64, avail uint64, err error) {
	num := 2
	s := bufio.NewScanner(bytes.NewReader(buf))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		l := s.Text()
		parts := whitespaces.Split(strings.TrimSpace(l), -1)
		if len(parts) != 3 || parts[2] != "kB" {
			continue
		}
		if parts[0] == "MemTotal:" {
			if total, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return
			}
			num--
		}
		if parts[0] == "MemAvailable:" {
			if avail, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return
			}
			num--
		}
	}
	if num != 0 {
		err = errors.New("missing MemTotal or MemAvailable")
		return
	}
	if avail > total || total == 0 {
		err = errors.New("invalid MemTotal or MemAvailable")
		return
	}
	return
}
