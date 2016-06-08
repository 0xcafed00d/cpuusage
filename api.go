// Package cpuusage library that monitors real time CPU usage (total and per-core) (linux only)
//
// Installation:
//   go get github.com/simulatedsimian/cpuusage
//
// Example:
//   // display the cpu usage once per second
//   u := cpuusage.Usage{}
//
//   for {
//       err := u.Measure()
//       if err != nil {
//           // handle error....
//       }
//       fmt.Printf("Overall %%: %d Per Core %%: %v\n", u.Overall, u.Cores)
//       time.Sleep(1 * time.Second)
//  }
package cpuusage

// Usage struct holds the overall cpu usage and the per-core usage resulting from the last call to Measure().
// Each usage is an int in the range 0..100 representing the percentage cpu utilisation
type Usage struct {
	// Overall is the combined cpu usage of all cores in the system
	Overall int
	// Cores is a slice of ints holding the usage of each core (physical or virtual) in the system
	Cores    []int
	previous *cpuinfo
}

// Measure reads the /proc/stat file and extracts the cpu usage information, writing it into the Usage struct.
// The first call to Measure gives the usage since the system was booted, with subsequent calls giving the
// usage during the period of time from the previous call to the current one.
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
