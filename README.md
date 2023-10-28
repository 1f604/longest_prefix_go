# longest_prefix_go

Benchmark longest prefix matching algorithms in Go

Includes brute force and Trie implementations.

To run:

```
$ go test -bench . -benchmem -count 10
```

Results (repeating the search each time):

```
100 prefixes:
Benchmark_brute_force-4         1000000000               0.0004339 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004730 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004233 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002044 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004709 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004425 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004498 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004333 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004364 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0004430 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009243 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0008939 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001186 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001059 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009180 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001082 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.008919 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0008859 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009051 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009660 ns/op               0 B/op          0 allocs/op

200 prefixes:
Benchmark_brute_force-4         1000000000               0.0008050 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0008464 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0009719 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0009954 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0009179 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0007887 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0009607 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0007892 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.0009670 ns/op               0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001257 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0008922 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001132 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009905 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009497 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001020 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001003 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001134 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001200 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009214 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001063 ns/op        0 B/op          0 allocs/op

300 prefixes:
Benchmark_brute_force-4         1000000000               0.001215 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001286 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001225 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001205 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001225 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001216 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001185 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001844 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.001192 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.005851 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.003052 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009906 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009440 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001003 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001094 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001058 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001004 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0009754 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.0008557 ns/op               0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001059 ns/op        0 B/op          0 allocs/op

500 prefixes:
Benchmark_brute_force-4         1000000000               0.002014 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002054 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002112 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002452 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002174 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002104 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002221 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002075 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.002240 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003273 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001339 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001142 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001457 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.005727 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001278 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001388 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001411 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001459 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001899 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001086 ns/op        0 B/op          0 allocs/op

1,000 prefixes:
Benchmark_brute_force-4         1000000000               0.003930 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.005062 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.004669 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003774 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003873 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003811 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003627 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003810 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.003839 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.004597 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.002341 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001529 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001199 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001232 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001462 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001557 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.005368 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001423 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001095 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001523 ns/op        0 B/op          0 allocs/op

2,000 prefixes:
Benchmark_brute_force-4         1000000000               0.008191 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007451 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007611 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007673 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.008356 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007481 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007797 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.009720 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007594 ns/op        0 B/op          0 allocs/op
Benchmark_brute_force-4         1000000000               0.007353 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001425 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001683 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001553 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001489 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.002044 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001367 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001960 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001659 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.002187 ns/op        0 B/op          0 allocs/op
Benchmark_Trie-4                1000000000               0.001259 ns/op        0 B/op          0 allocs/op
```

Results suggest that brute force is faster when number of prefixes is less than 200, but Trie is faster when number of prefixes exceeds 300.

You may get different results on your machine.