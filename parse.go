package cpuusage

import (
	"fmt"
	"strings"
)

type coreinfo struct {
	name                     string
	user, nice, system, idle int
}

type cpuinfo struct {
	overall coreinfo
	cores   []coreinfo
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
