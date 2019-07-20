# btree

Binary search tree implementation in Go.

This was built with the intention of teaching myself Go, and is not meant to be used in production systems.

## Benchmark

```text
goos: linux
goarch: amd64
pkg: github.com/sircelsius/btree
BenchmarkBinarySearchTree_Search100-8      	20000000	        81.9 ns/op
BenchmarkBinarySearchTree_Search1000-8     	20000000	       106 ns/op
BenchmarkBinarySearchTree_Search10000-8    	10000000	       165 ns/op
BenchmarkBinarySearchTree_Search100000-8   	 5000000	       259 ns/op
PASS
ok  	github.com/sircelsius/btree	7.490s

```