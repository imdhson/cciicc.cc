## YYYY-MM-DD - Initial observation
**Learning:** The application uses global slice singletons for spaces (`var spaces *Spaces`) and users (`var users *Users`).
**Action:** Need to find an O(N^2) loop or something inefficient to optimize. Memory mentions "When removing elements from global slice singletons in the types package (e.g., spaces or users), use backward iteration to prevent index shifting and avoid O(N^2) restart inefficiencies."
## 2023-10-24 - Slice Iteration Optimization
**Learning:** Found unnecessary struct copying and redundant pointer dereferencing in `UnusedSpaceRemoveService.go`. Iterating with `for _, v := range` copies structs, which wastes memory/CPU, and dereferencing pointers repeatedly in loops is inefficient.
**Action:** Use index-based loops (`for i := range`) to avoid struct copies and dereference slice pointers once outside the loop.
## 2023-10-24 - Efficient Slice Iteration for Filter in Go
**Learning:** In Go, filtering slices in place without copying elements into loop variables is significantly faster. Dereferencing pointer to slice over and over again in the loop wastes cycles.
**Action:** Replace `for _, v := range *users` with `usersSlice := *users` then `for i := range usersSlice` and accessing `usersSlice[i]`.
