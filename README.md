# Lock free collections

This library use unsafe and atomic. So it is not crossplatform.


## Queue:
    Use Michael Scott queue algorithm.
### Bench AMD 5545M:
    BenchmarkQueueEnq-4	 2000000	       959 ns/op       
    BenchmarkQueueDeq-4	 2000000	       766 ns/op

## Stack:
### Bench AMD 5545M:
    BenchmarkStackPush-4	 1000000	      1417 ns/op
    BenchmarkStackPop-4 	 1000000	      1229 ns/op

# Tasks

- [x] Queue
- [x] Stack
- [ ] Map
- [ ] Linked List
