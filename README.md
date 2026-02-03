# Notes
Format and save changes against original file; 
```bash
$ gofmt -w main.go
```

Run file;
```bash
$ go run main.go
```

Build file;
```bash
$ go build main.go
```

Declare variables
```go
// full declaration
<varName> <varType> = <varValue>
carType string = "Wagon"

// short declaration
carType := "Wagon"

// declaration with default value
carFueled bool // default to False

// declare variables as a group
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```