package main

import "testing"

// Benchmark should be prefixed in benchmarking func
func BenchmarkAnalyzeText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := AnalyzeText("moby.txt")
		if err != nil {
			b.Fatalf("Error analyzing text: %v", err)
		}
	}
}

func BenchmarkOptimizedAnalyzeText(b *testing.B) {
	filename := "moby.txt"
	for i := 0; i < b.N; i++ {
		_, err := OptimizedAnalyzeText(filename)
		if err != nil {
			b.Fatalf("Error analyzing text: %v", err)
		}
	}
}

// the below command is to run benchmark for a specific number of time
//go test -run none -bench . -benchtime 3s -benchmem -v
//BenchmarkAnalyzeText-8: This is the name of the benchmark test. The -8 part means that the benchmark was run on 8 threads to parallelize and speed up the execution.
//235: This is the number of iterations that were performed while benchmarking the function. The benchmark function is executed b.N times, where b.N increases in each iteration until the benchmark function lasts long enough to be measured accurately.
//15517327 ns/op: This means that each operation in the benchmark took an average of approximately 15.517327 milliseconds to execute. This is a measurement of time consumed per operation (ns/op stands for "nanoseconds per operation").
//27750673 B/op: This means that each operation in the benchmark caused about 27,750,673 bytes (or approximately 27MB) of allocations in memory. (B/op stands for "bytes per operation"). This can help in identifying memory-intensive operations.
//710 allocs/op: This means that every execution of the benchmarked function resulted in 710 allocations from the heap (allocs/op stands for "heap allocations per operation"). This gives an insight into how much work the garbage collector will need to do as a result of running this function.

//below command generates memory profile
//go test -run none -bench . -benchtime 3s -benchmem -v -memprofile p.out
// go tool pprof p.out
// list AnalyzeText
// top 5
// weblist AnalyzeText // to see ui version of the report

/*

  Total:      1.59GB     8.57GB (flat, cum)   100%
     12            .          .           	AvgWordLength float64
     13            .          .           	LongestWord   string
     14            .          .           }
     15            .          .
     16            .          .           func AnalyzeText(filename string) (TextAnalysis, error) {
     17            .   397.96MB           	content, err := os.ReadFile(filename)
     18            .          .           	if err != nil {
     19            .          .           		return TextAnalysis{}, err
     20            .          .           	}
     21            .          .
     22     403.30MB   403.30MB           	text := string(content)
     23            .     6.59GB           	words := strings.Fields(text)
     24            .          .
     25            .          .           	frequency := make(map[string]int)
     26            .          .           	totalLength := 0
     27            .          .           	longest := ""
     28            .          .
     29            .          .           	for _, word := range words {
     30       1.19GB     1.19GB           		frequency[word]++
     31            .          .           		if len(word) > len(longest) {
     32            .          .           			longest = word
     33            .          .           		}
     34            .          .           		totalLength += len(word)
*/

/*
explanation of the chart
This line is the header of your pprof result. The "flat" column corresponds to the memory allocated by a function itself and the "cumulative" column corresponds to the memory allocated by a function including its child calls.
** The cum value, short for 'cumulative', tells you the resources used by the function and all the functions it calls.** , so basically cumulative means how much total memory allocation happened
The values 1.59GB and 8.57GB are the total memory measured diagonally from the top function to the bottom taking into consideration the flat and cumulative values respectively.



*/
