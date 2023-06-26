# Benchmark and Unit Testing

Benchmark testing
```bash
go test -bench=JSON ./testing/benchmark/main_test.go -benchmem
go test -bench=Sonic ./testing/benchmark/main_test.go -benchmem
```

Unit testing
```bash
go test -v ./testing/unit/main_test.go -run=JSON
```