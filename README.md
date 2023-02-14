## version

```
> go version
go version go1.20 linux/amd64
```

## result

```
> go test -v ./...
=== RUN   TestJsonIter_behavioral_deference
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey1. org = {"Foo":"foo","Baz":"bar"}, jsoniter = {"Foo":"foo","Baz":"baz"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey3. org = {"Foo":"foo"}, jsoniter = {"Foo":"foo","Baz":"qux"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey4. org = {"Foo":"foo","Bar":"bar","Baz":"baz"}, jsoniter = {"Foo":"foo","Baz":"baz","Bar":"barbar"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey1. org = {"Foo":"foo","Baz":"bar"}, jsoniter = {"Foo":"foo","Baz":"baz"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey3. org = {"Foo":"foo"}, jsoniter = {"Foo":"foo","Baz":"qux"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey4. org = {"Foo":"foo","Bar":"bar","Baz":"baz"}, jsoniter = {"Foo":"foo","Baz":"baz","Bar":"barbar"}
--- FAIL: TestJsonIter_behavioral_deference (0.00s)
FAIL
FAIL    github.com/ngicks/test-jsoniter.git     0.002s
FAIL
```
