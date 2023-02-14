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

Recursive types causes stack overflow.

```
> go test -v ./...
=== RUN   TestJsonIter_behavioral_deference
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey1. org = {"Foo":"foo","Baz":"bar"}, jsoniter = {"Foo":"foo","Baz":"baz"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey3. org = {"Foo":"foo"}, jsoniter = {"Foo":"foo","Baz":"qux"}
    recursive_test.go:82: not same, type = testjsoniter_test.OverlappingKey4. org = {"Foo":"foo","Bar":"bar","Baz":"baz"}, jsoniter = {"Foo":"foo","Baz":"baz","Bar":"barbar"}
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0xc024700370 stack=[0xc024700000, 0xc044700000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x5ac7a4?, 0x6caa80?})
        /usr/local/go/src/runtime/panic.go:1047 +0x5d fp=0xc000495e18 sp=0xc000495de8 pc=0x435d9d
...error continues...
```
