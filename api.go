// Package cpuusage library that monitors real time CPU usage (total and per-core) (linux only)
package cpuusage

// Usage struct holds the overall cpu usage and the per-core usage resulting from the last call to Measure().
// Each usage is an int in the range 0..100 representing the percentage cpu utilisation
type Usage struct {
	// Overall is the combined cpu usage of all cores in the system
	Overall int

	Cores    []int
	previous *cpuinfo
}

// Measure reads the /proc/stat file and extracts the cpu usage information, writing it into the Usage struct.
func (u *Usage) Measure() error {
	cpu, err := readProcStat()
	if err != nil {
		return err
	}

	prev := cpu.clone()
	if u.previous != nil {
		cpu.overall = delta(cpu.overall, u.previous.overall)
		for i := range cpu.cores {
			cpu.cores[i] = delta(cpu.cores[i], u.previous.cores[i])
		}
	}

	u.Overall = calcUsage(cpu.overall)
	u.Cores = nil
	for i := range cpu.cores {
		u.Cores = append(u.Cores, calcUsage(cpu.cores[i]))
	}

	u.previous = &prev
	return nil
}
