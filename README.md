# Interview Go Projects

A collection of small Go programs and pattern examples for interview preparation. Each sub-project is a standalone `main.go` demonstrating algorithms, concurrency, data structures, or GoF design patterns.

## Project Structure

```text
project/
├── GOFpatterns/        # GoF pattern examples
├── GrokkingAlgorithms/ # Algorithm practice
├── InterviewPractice/  # Interview-focused exercises
├── Leetcode/           # LeetCode tasks grouped by category
├── LeetcodeTheory/     # Supporting theory and data structure examples
├── exercises/          # Small standalone Go exercises
├── simplePractice/     # Small Go practice programs
├── go.mod
└── README.md
```

## LeetCode Structure

All current task folders inside `Leetcode/` are grouped into PascalCase categories.

### Categorized tasks

- `Leetcode/Backtracking/`
  - `17_letter_combinations_of_a_phone_number/`
- `Leetcode/BinarySearch/`
  - `4_median_of_two_sorted_arrays_hard/`
  - `704_Binary_Search/`
- `Leetcode/BitManipulation/`
  - `231_Power_of_Two/`
- `Leetcode/DynamicProgramming/`
  - `70_Climbing_Stairs/`
- `Leetcode/Graphs/`
  - `200_Number_of_Islands/`
  - `490_The_Maze/`
  - `733_Flood_Fill/`
- `Leetcode/HashTables/`
  - `2_twoSum_medium/`
  - `242_Valid_Anagram/`
- `Leetcode/Intervals/`
  - `56_Merge_Intervals/`
- `Leetcode/LinkedLists/`
  - `1_addTwoNumbers_easy/`
  - `876_Middle_of_the_Linked_List/`
- `Leetcode/Math/`
  - `6_reverse_integer/`
  - `8_is_palindrome_number/`
- `Leetcode/Matrices/`
  - `1572_Matrix_Diagonal_Sum/`
  - `867_Transpose_Matrix/`
- `Leetcode/PrefixSum/`
  - `303_Range_Sum_Query_Immutable/`
- `Leetcode/SlidingWindow/`
  - `3_lengthOfLongestSubstring/`
- `Leetcode/Sorting/`
  - `905_Sort_Array_By_Parity/`
- `Leetcode/StacksAndQueues/`
  - `225_Implement_Stack_using_Queues/`
  - `232_Implement_Queue_using_Stacks/`
- `Leetcode/Strings/`
  - `5_zig_zag_conversion/`
  - `7_custom_atoi/`
  - `11_integer_to_roman/`
- `Leetcode/Trees/`
  - `94_Binary_Tree_Inorder_Traversal/`
- `Leetcode/TwoPointers/`
  - `10_container_with_the_most_wate1/`
  - `15_3Sum_medium/`
  - `18_4Sum/`
  - `27_Remove_Element/`
  - `125_Valid_Palindrome/`
  - `344_Reverse_String/`
  - `88_merge_sorted_array/`

## How to Run Examples

From the project root, run any example with `go run` pointing to its `main.go`, for example:

```bash
# Run an exercise
go run ./exercises/5_ordersQueueApp/main.go

# Run a GoF pattern example
go run ./GOFpatterns/Creational/Builder/3_builder.go

# Run a categorized LeetCode solution
go run ./Leetcode/BinarySearch/704_Binary_Search/main.go
```
