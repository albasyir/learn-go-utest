# U Test on Golang

describe how to run unit test on golang

## Unit Test

- must be declared by func `Test..` prefix, ex. `TestSayHello`

### Run utest

```bash
go test -v ./...
```

### Run one package

```bash
go test -v ./helper/
```

`./helper/` is package name

### mocking and assert

this example using `testify` package, see on project how to implement mocking and assert

### Sub Test

can declare subtest that can be use for **table testing implementation**, see example on this project

## Banchmark

### Run banchmark without unit test

```bash
go test -v -run=UnitTestHasNotExists -bench=.
```

note that you must define a unit test that has not implemented, like `UnitTestHasNotExists`, this has not implemented by this project

### Run spesific banchmark without unit test

```bash
go test -v -run=UnitTestHasNotExists -bench=BenchmarkSayHello
```

### Sub Banchmark

can declare sub banchmark that can be use for **table banchmarking implementation**, see example on this project
