package cpuusage

import (
	"fmt"
	"strings"
)

type coreinfo struct {
	name                     string
	user, nice, system, idle int
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
	total := used + c.idle
	return (used * 100) / total
}

type cpuinfo struct {
	overall coreinfo
	cores   []coreinfo
}

func (c cpuinfo) clone() (cc cpuinfo) {
	cc.overall = c.overall
	cc.cores = append(cc.cores, c.cores...)
	return
}

func parserCoreInfo(s string) (coreinfo, error) {
	core := coreinfo{}

	_, err := fmt.Sscanf(s, "%s %d %d %d %d", &core.name, &core.user, &core.nice, &core.system, &core.idle)

	if err != nil {
		return coreinfo{}, fmt.Errorf("%s: [%s]", err, s)
	}
	if !strings.HasPrefix(core.name, "cpu") {
		return coreinfo{}, fmt.Errorf("stats are not for cpu %s", s)
	}
	return core, nil
}

func parseCPU(lines []string) (cpuinfo, error) {
	cpu := cpuinfo{}

	for _, s := range lines {
		if strings.HasPrefix(s, "cpu") {
			core, err := parserCoreInfo(s)
			if err != nil {
				return cpuinfo{}, err
			}
			if core.name == "cpu" {
				cpu.overall = core
			} else {
				cpu.cores = append(cpu.cores, core)
			}
		}
	}
	return cpu, nil
}
