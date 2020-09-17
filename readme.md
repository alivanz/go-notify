# Notify Once

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

```
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
