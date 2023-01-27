# Comparison of Quicksort algorithm in NodeJS and GoLang

| Language | Execution Time (ms) |
| -------- | ------------------- |
| Go       | 12.39 ± 1.99        |
| NodeJS   | 24.27 ± 1.55        |

The above table shows the execution time of quicksort algorithm implemented in Go and Node.js languages. It's clear from the results that Go outperforms Node.js in terms of execution time by ~52%.

It is worth noting that the performance of a programming language depends on various factors such as the implementation of the algorithm, the hardware, and the operating system being used. The above results should be taken as a rough estimate and not as a definitive answer.

## How to test

### GoLang
To run first you need [Go](https://go.dev/dl/).

the correct way is build the app:
```sh
go build main.go
```

then run 

```sh
./main
```

but you also can run with:
```sh
go run main.go
```

### Node

You also need [NodeJS](https://nodejs.org/en/download/).

JS is a interpreted language, so you only need to run:

```sh
node index.js
```