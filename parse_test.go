package cpuusage

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestCoreParse(t *testing.T) {
	assert := assert.Make(t)

	assert(parserCoreInfo("cpu0 29508 4564 8526 717502 4034 0 54 0 0 0")).Equal(coreinfo{"cpu0", 29508, 4564, 8526, 717502}, nil)
	assert(parserCoreInfo("cpu0 29508 4564 8526")).HasError()
	assert(parserCoreInfo("cup0 29508 4564 8526 717502 4034 0 54 0 0 0")).HasError()
}

func TestCPUParse(t *testing.T) {
	assert := assert.Make(t)
	cpu, err := parseCPU(procstat)
	assert(err).NoError()
	assert(cpu.overall).Equal(coreinfo{"cpu", 152770, 16132, 44016, 1109080})
	assert(len(cpu.cores)).Equal(4)
	assert(cpu.cores[0]).Equal(coreinfo{"cpu0", 41421, 4804, 11600, 1083260})
	assert(cpu.cores[1]).Equal(coreinfo{"cpu1", 42209, 3088, 12371, 8576})
	assert(cpu.cores[2]).Equal(coreinfo{"cpu2", 34999, 3901, 9530, 8582})
	assert(cpu.cores[3]).Equal(coreinfo{"cpu3", 34139, 4336, 10514, 8661})

	assert(parseCPU(procstaterr)).HasError()
}

var procstat = []string{
	"cpu  152770 16132 44016 1109080 4988 0 218 0 0 0",
	"cpu0 41421 4804 11600 1083260 4734 0 64 0 0 0",
	"cpu1 42209 3088 12371 8576 64 0 32 0 0 0",
	"cpu2 34999 3901 9530 8582 141 0 44 0 0 0",
	"cpu3 34139 4336 10514 8661 49 0 76 0 0 0",
	"intr 7096213 19 42514 0 0 0 0 0 0 1 419 0 0 897192 0 0 0 69 0 0 0 0 0 0 80 25220 40 159216 0 316593 16 230562 13 213",
	"ctxt 30364916",
	"btime 1465222999",
	"processes 51345",
	"procs_running 2",
	"procs_blocked 0",
	"softirq 3821819 59 1367751 103 4564 155102 6 28624 1240883 0 1024727",
}

var procstaterr = []string{
	"cpu  152770 16132 44016 1109080 4988 0 218 0 0 0",
	"cpu0 41421 4804 11600 1083260 4734 0 64 0 0 0",
	"cpu1 xxxx 3088 12371 8576 64 0 32 0 0 0",
	"cpu2 34999 3901 9530 8582 141 0 44 0 0 0",
	"cpu3 34139 4336 10514 8661 49 0 76 0 0 0",
	"intr 7096213 19 42514 0 0 0 0 0 0 1 419 0 0 897192 0 0 0 69 0 0 0 0 0 0 80 25220 40 159216 0 316593 16 230562 13 213",
	"ctxt 30364916",
	"btime 1465222999",
	"processes 51345",
	"procs_running 2",
	"procs_blocked 0",
	"softirq 3821819 59 1367751 103 4564 155102 6 28624 1240883 0 1024727",
}
