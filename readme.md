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
  for i := 0; i < 3; i++ {
    go func(i){
      for {
        data, done := bytesNotify.Listen()
        log.Printf("%d %s", i, string(data))
        <-done
      }
    }(int i)
  }
  for i := 0; ; i++ {
    time.Sleep(1 * time.Second)
    bytesNotify.Notify([]byte(fmt.Sprintf("msg#%d", i)))
  }
}
```

```txt
0 first
1 first
2 first
0 msg#0
1 msg#0
2 msg#0
0 msg#1
1 msg#1
2 msg#1
0 msg#2
1 msg#2
2 msg#2
....
```
