package cpuusage

type Usage struct {
	Overall  int
	Cores    []int
	previous *cpuinfo
}

func delta(c1, c2 coreinfo) coreinfo {
	return coreinfo{
		name:   c1.name,
		user:   c2.user - c1.user,
		nice:   c2.nice - c1.nice,
		idle:   c2.idle - c1.idle,
		system: c2.system - c1.system,
	}
}

func calcUsage(c coreinfo) int {
	used := c.nice + c.system + c.user
	return (used * 100) / (used + c.idle)
}

func (u *Usage) GetUsage() error {
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
