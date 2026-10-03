## 2024-05-24 - Slice Backward Iteration Pattern
**Learning:** In Go backend services where new items (like spaces or users) are frequently appended to global slice singletons, locating elements via forward iteration is inefficient.
**Action:** When searching elements in global slices with frequent appends, use backward iteration (`for i := len(slice) - 1; i >= 0; i--`). Also, dereference slice pointers once outside the loop to avoid redundant operations on each iteration.
