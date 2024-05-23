* try to avoid passing pointers to a function
* a map in go is implemented as a pointer to a struct; so if you pass a map to a function you can actually modify the map in the function
* rather than pass a map around, use a struct
* if you pass a slice to a function then modifications to the slice will be captured but any appends will not; because you are actually dealing with a copy of the slice in the function, defined as a struct with 3 elements -- int for length, int for capacity, and a pointer to a block of memory; these values are copied into a new struct; and then if you add values then the length field is updated in the copy but not the original; so upon exiting the function you are back to a slices with an unchanged original length value -- it ignores the new extra units of memory added
* here is an example of using a slice as a buffer
  * with "file.Read" we are passing a 100-byte slice to the "Read" function; that function reads in 100b or less, MODIFYING THE ORIGINAL SLICE, but not going beyond that exact capacity
  * "Read" also returns a count of the # of bytes read
  * then the "process" function (not specified here) is called against all of the bytes read (up to the "count" value)
  * if "Read" also returned an error then we stop if it was an EOF file and return nil (and this is how we WANT to exit the for loop) or return an actual error
  * ~~~~
  * the reason why this is a good way to read the file is that we are not creating a brand new variable every time we read from the file, which is better from a memory allocation perspective
```
file, err := os.Open(fileName)
if err != nil {
    return err }
defer file.Close()
data := make([]byte, 100)
for {
    count, err := file.Read(data)
    process(data[:count])
    if err != nil {
        if errors.Is(err, io.EOF) {
            return nil
        }
        return err
    }
}
```
* garbage collection
  * heap vs stack?
    * a STACK is a consecutive block of memory; memory allocations are added to this, one on top of another, as more variables are allocated; but you can only store something on the stack if the size of what you are storing is known, e.g. if you are allocating primitive values, arrays, or structs; since pointer sizes are also known, pointers are placed on the stack; since you know the size of what you are storing, a "stack pointer" is simply moved when allocating to the stack; when a function is invoked a new "stack frame" is memory is created; and when the function exits the memory on the stack heap is automatically deallocated
      * again in order to allocate to the stack:
        * the data must be a local variable whose size is known at compile time
        * a pointer cannot be returned from the function
    * if the compilier determines that something that needs to be allocated to memory doesn't meet the stack requirements then it is placed on the HEAP
      * garbage collection in go works with the HEAP, not the STACK
      * something on the HEAP is NOT eligible for garbage collection so long as it has at least one pointer referencing it
  * you can run `go build -gcflags="-m"` to see which values "escape to the heap"
    * you'll notice that pointers do
  * garbage collection will work better in go if you use pointers less, as the less you use pointers the more memory use will be sequential
  * ~~~~~
  * tuning garbage collection by environment variables:
    * GOGC environment variable
      * size heap can get to before a garbage collection occurring
      * set to 100 by default
      * making it larger will make garbage collection less frequent, speeding up program
      * setting it to `off` will disable garbage collection entirely, making program even faster but on long-running programs will certainly cause all memory to get used up
      * ~~~~~~
      * if you fool around with raising this value -- see exercise 3, you will see diminishing returns
    * GOMEMLIMIT
      * limit on total amount of memory the progam can use
      * disabled by default -- well, set to math.MaxInt64, which is very large
      * it is a soft limit that can be exceeded in particular circumstances
  * ~~~~~~~~
  * you can set the GODBUG environment variable (`GODEBUG=gctrace=1`) to see when garbage collections occur and how changing GOGC affects that


