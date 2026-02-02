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
```