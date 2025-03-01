# 批处理模式

The batching resiliency pattern for golang.

Creating a batcher takes two parameters:
- the timeout to wait while collecting a batch
- the function to run once a batch has been collected

You can also optionally set a prefilter to fail queries before they enter the
batch.

```go
b := batcher.New(10*time.Millisecond, func(params []interface{}) error {
	// do something with the batch of parameters
	return nil
})

b.Prefilter(func(param interface{}) error {
	// do some sort of sanity check on the parameter, and return an error if it fails
	return nil
})

for i := 0; i < 10; i++ {
	go b.Run(i)
}
```
