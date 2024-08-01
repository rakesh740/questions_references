To limit the number of goroutines running at any given time to 100, you can use a semaphore pattern. This can be implemented using a buffered channel in Go. The buffered channel will act as a semaphore that limits the number of goroutines that can run concurrently.

Here is how you can modify your code to achieve this:

```go
package main

import (
 "time"
)

func main() {
 const maxGoroutines = 100
 sem := make(chan struct{}, maxGoroutines)

 for i := 0; i < 1000; i++ {
  sem <- struct{}{}  // acquire a slot
  go func(i int) {
   defer func() { <-sem }()  // release the slot after the goroutine is done
   time.Sleep(1 * time.Second)
  }(i)
 }

 // Wait for all goroutines to finish
 for i := 0; i < maxGoroutines; i++ {
  sem <- struct{}{}
 }
}
```

### Explanation:

1. **Semaphore Creation**: `sem := make(chan struct{}, maxGoroutines)` creates a buffered channel that can hold up to `maxGoroutines` elements. This channel acts as a semaphore to limit the number of concurrent goroutines.

2. **Acquiring a Slot**: `sem <- struct{}{}` sends an empty struct to the channel. If the channel is full (i.e., it already holds `maxGoroutines` elements), this operation will block until a slot becomes available.

3. **Releasing a Slot**: `defer func() { <-sem }()` receives from the channel, releasing the slot once the goroutine is done.

4. **Waiting for Completion**: The loop `for i := 0; i < maxGoroutines; i++ { sem <- struct{}{} }` ensures that the main function waits for all the goroutines to finish. This is done by filling the semaphore channel again, effectively ensuring that all running goroutines have completed.

This way, you ensure that at most 100 goroutines are running concurrently at any given time.