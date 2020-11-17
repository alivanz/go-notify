# Notify Once

```Go
func main() {
  once := notify.NewNotifyOnce()
  for i := 0; i < 3; i++ {
    go func(i){
      <-notify.Listen()
      log.Printf("%d first")
      <-notify.Listen()
      log.Printf("%d second")
    }(int i)
  }
  once.Notify()
}
```

```txt
0 first
1 first
2 first
0 second
1 second
2 second
```

# Notify Bytes
```Go
func main() {
  bytesNotify := notify.NewBytes([]byte("first"))
  go func() {
    for {
      data, done := bytesNotify.Listen()
      log.Printf("alice got %s", string(data))
      <-done
    }
  }()
  go func() {
    for {
      data, done := bytesNotify.Listen()
      log.Printf("bob got %s", string(data))
      <-done
    }
  }()
  go func() {
    for {
      data, done := bytesNotify.Listen()
      log.Printf("lucy got %s", string(data))
      <-done
    }
  }()
  for i := 0; ; i++ {
    time.Sleep(1 * time.Second)
    bytesNotify.Notify([]byte(fmt.Sprintf("msg#%d", i)))
  }
}
```

```txt
alive got first
bob got first
lucy got first
alive got msg#0
bob got msg#0
lucy got msg#0
alive got msg#1
bob got msg#1
lucy got msg#1
alive got msg#2
bob got msg#2
lucy got msg#2
....
```

# Network notify
Server side
```go
func main() {
  n := notify.NewInterface()
}
```
