# Benchmarks
Test Environment
```
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
```

## String Concatenation
每次串連一個字符的情況，有些方法可以事先分配所需空間
* Time
```
Assign > StringsBuiderWithGrow > Copy > StringsBuider > BytesBuffer > NormalConcat
```
* Memory
```
StringsBuiderWithGrow > Assign = Copy > BytesBuffer > StringsBuider > NormalConcat
```

```
BenchmarkNormalConcat-8                   959862             49092 ns/op          483923 B/op          1 allocs/op
BenchmarkBytesBuffer-8                  207401966                5.632 ns/op           3 B/op          0 allocs/op
BenchmarkStringsBuider-8                621018559                3.504 ns/op           5 B/op          0 allocs/op
BenchmarkStringsBuiderWithGrow-8        801707186                1.501 ns/op           1 B/op          0 allocs/op
BenchmarkCopy-8                         405525543                2.949 ns/op           2 B/op          0 allocs/op
BenchmarkAssign-8                       1000000000               0.6008 ns/op          2 B/op          0 allocs/op
```

每次串連未知長度字符串的情況
* Time
```
StringsBuider >  BytesBuffer > NormalConcat
```
* Memory
```
BytesBuffer > StringsBuider > NormalConcat
```
```
BenchmarkNormalConcatLongString-8         208395            103228 ns/op         1046032 B/op          1 allocs/op
BenchmarkBytesBufferLongString-8        100000000               18.58 ns/op           34 B/op          0 allocs/op
BenchmarkStringsBuiderLongString-8      100000000               10.82 ns/op           61 B/op          0 allocs/op
```

## Convert Bytes and String
標準轉換會有一次重新分配記憶體的操作，而 unsafe.Pointer 是直接替換指標的指向，所以後者的效能會比較好。

```
BenchmarkConvertBytesToString-8                         272225202                4.561 ns/op           0 B/op          0 allocs/op
BenchmarkConvertBytesToStringWithUnsafePointer-8        1000000000               0.2768 ns/op          0 B/op          0 allocs/op
BenchmarkConvertStringToBytes-8                         234166994                5.235 ns/op           0 B/op          0 allocs/op
BenchmarkConvertStringToBytesWithUnsafePointer-8        1000000000               0.2792 ns/op          0 B/op          0 allocs/op
```
不過 BytesToString 和 StringToBytes 有差異是，String 所需的空間大小比 Bytes 大，且 String 是不可修改的，
如果直接透過 unsafe.Pointer 將 Bytes 轉成 String 會有以下問題：
```
s := "Test"
bytes := *(*[]byte)(unsafe.Pointer(&s))
bytes[0] = 'X'

unexpected fault address 0x110d635
fatal error: fault
[signal SIGBUS: bus error code=0x2 addr=0x110d635 pc=0x10e3d00]
```
所以 StringToBytes 的部分建議還是使用官方提供的標準函式來轉換比較安全。