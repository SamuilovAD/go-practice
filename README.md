# Interview Go Projects

A collection of small Go programs and pattern examples for interview preparation. Each sub-project is a standalone `main.go` demonstrating algorithms, concurrency, data structures, or GoF design patterns.

## Project Structure

```
project/
├── GOFpatterns/
│   └── Creational/
│       ├── Builder/
│       │   └── 3_builder.go - Builder pattern example
│       ├── FactoryMethod/
│       │   └── 2_factoryMethod.go - Factory Method pattern example
│       └── Singleton/
│           └── 1_Singleton.go - Singleton pattern example
├── exercises/
│   ├── 10_designHashMapApp/
│   │   └── main.go - Simple hash map implementation
│   ├── 11_loopInPointerApp/
│   │   └── main.go - Detect loop in pointer-based structures
│   ├── 12_reverseNodeApp/
│   │   └── main.go - Reverse nodes in a linked list
│   ├── 13_twoSumApp/
│   │   └── main.go - Two Sum using hash map
│   ├── 2_dataHandlingPipelineApp/
│   │   └── main.go - Data pipeline with goroutines and channels
│   ├── 3_logsFilesParserApp/
│   │   ├── files/
│   │   │   └── log1.txt
│   │   └── main.go - Parse log files and extract info
│   ├── 4_oddEvenPrintApp/
│   │   └── main.go - Odd/Even printing with synchronization
│   ├── 5_ordersQueueApp/
│   │   └── main.go - Order processing with timeouts
│   ├── 6_PointerTestApp/
│   │   └── main.go - Pointer semantics and addresses demo
│   ├── 7_sortbubbleApp/
│   │   └── main.go - Bubble sort implementation
│   ├── 8_workersPoolApp/
│   │   └── main.go - Worker pool for concurrent tasks
│   └── 9_workersPoolFixed/
│       └── main.go - Fixed-size worker pool
├── go.mod
└── leetcode/
    ├── 1_addTwoNumbers_easy/
    │   └── main.go - Add Two Numbers (linked lists)
    ├── 2_twoSum_medium/
    │   └── main.go - Two Sum (optimized)
    ├── 3_lengthOfLongestSubstring_medium/
    │   └── main.go - Longest substring without repeats (sliding window)
    └── 4_median_of_two_sorted_arrays_hard/
        └── main.go - Median of two sorted arrays (efficient)
```

## How to run examples

From the project root, run any example with `go run` pointing to its `main.go`, for example:

```
# Run an exercise
go run ./exercises/5_ordersQueueApp/main.go

# Run a GoF pattern example
go run ./GOFpatterns/Creational/Builder/3_builder.go

# Run a LeetCode solution
go run ./leetcode/3_lengthOfLongestSubstring_medium/main.go
```
