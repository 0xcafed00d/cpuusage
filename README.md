# cpuusage [![GoDoc](https://godoc.org/github.com/simulatedsimian/cpuusage?status.svg)](https://godoc.org/github.com/simulatedsimian/cpuusage)

Go library that monitors real time CPU usage (total and per-core) (linux only)

## Installation:
```bash
$ go get github.com/simulatedsimian/cpuusage
```

## Example:
```go
import "github.com/simulatedsimian/cpuusage"
```
```go
   // display the cpu usage once per second
u := cpuusage.Usage{}

for {
  err := u.Measure()
  if err != nil {
    // handle error....
  }
  fmt.Printf("Overall %%: %d Per Core %%: %v\n", u.Overall, u.Cores)
  time.Sleep(1 * time.Second)
}
```
