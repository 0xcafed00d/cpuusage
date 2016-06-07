package cpuusage

import (
	"bufio"
	"os"
)

func readProcStat() (cpuinfo, error) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return cpuinfo{}, err
	}
	defer f.Close()

	lines := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return cpuinfo{}, err
		}
		lines = append(lines, scanner.Text())
	}

	return parseCPU(lines)
}
