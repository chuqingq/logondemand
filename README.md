# logondemand
log on demand for Go.

# Example

client: 

```Go
filename := "logondemand_test"
l, err := logondemand.New(filename)
if err != nil {
    log.Fatalf("New error: %v", err)
}
logger := log.New(l, "", log.Flags()|log.Lshortfile|log.Lmicroseconds)
for {
    logger.Printf("123")
    time.Sleep(1 * time.Second)
}
```

watch:

```Go
filename := "logondemand_test"
err := logondemand.Cat(filename)
if err != nil {
    log.Fatalf("Cat error: %v", err)
}
```