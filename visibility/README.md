# Структура
```bash
bin
    // executable files
pkg
    // temporary files
src
    packageNameA
        packageFile1.go
        packageFile2.go
    packageNameB
        packageFile1.go
        packageFile2.go
    main.go
```

# Внутри пакета конвеншины
- паременные с маленькой буквы неимпортируемые наружу
- паременные с большой буквы импортируемые наружу

```go
package A

var (
    foo int // not importable
    Bar int // importable
)
```

