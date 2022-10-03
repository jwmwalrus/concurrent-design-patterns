# concurrent-design-patterns
Concurrent design patterns.


## References

* Go Concurrency Patterns, examples: 
    * https://github.com/lotusirous/go-concurrency-patterns

* Rob Pyke’s Talks on Concurrency Patterns:
    * https://go.dev/talks/2012/concurrency.slide#1
    * https://go.dev/talks/2013/advconc.slide#1 
    * https://www.youtube.com/watch?v=oV9rvDllKEg

* Concurrency in Go, Chapter 4:
    * https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html 

* Rethinking Classical Concurrency Patterns:
    * https://www.youtube.com/watch?v=5zXAHh5tJqQ 

* Wikipedia:
    * https://en.wikipedia.org/wiki/Communicating_sequential_processes
    * https://en.wikipedia.org/wiki/Futures_and_promises 
    * https://en.wikipedia.org/wiki/Generator_(computer_programming) 

* Fan-Out, Fan-In Scenarios In Durable Functions:
    * https://learn.microsoft.com/en-us/azure/azure-functions/durable/durable-functions-cloud-backup?tabs=csharp


## Futher examples

### Futures

| Example | Description |
| ------- | ----------- |
| [Google 1.0](https://github.com/lotusirous/go-concurrency-patterns/tree/main/9-google1.0) | Non-concurrent search for web, images and video |
| [Google 2.0](https://github.com/lotusirous/go-concurrency-patterns/tree/main/10-google2.0) | Concurrent search, using the futures pattern to search for web, images and video, respectively |
| [Google 2.1](https://github.com/lotusirous/go-concurrency-patterns/tree/main/11-google2.1) | Variation of Google 2.0 with timeouts |
| [Google 3.0](https://github.com/lotusirous/go-concurrency-patterns/tree/main/12-google3.0) | Variation of Google 2.1 with replicas for web, images and video |


### Generator

| Example | Description |
| ------- | ----------- |
| [Generator](https://github.com/lotusirous/go-concurrency-patterns/tree/main/3-generator) | A generator that returns the provided message, multiple times |
| [Subscription](https://github.com/lotusirous/go-concurrency-patterns/tree/main/14-adv-subscription) | An advanced subscription to a resource |

### Fan-Out, Fan-In

| Example | Description |
| ------- | ----------- |
| [Merge files](https://www.youtube.com/watch?v=x6vBvgKGvxU) | Reads two files and merges them into a single result, without using the sync package (Fan-In) |
| [Semaphore](https://dev.to/b0r/go-concurrency-patterns-fan-out-semaphore-1ojf) | Simulates the work of 10 employees, when only two seats/workstations are available (Fan-Out) |

