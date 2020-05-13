package main

import (
  "bytes"
  "encoding/gob"
  "io/ioutil"
  "fmt"
)

func EncodeGob(obj map[interface{}]interface{}) ([]byte, error) {
  for _, v := range obj {
    gob.Register(v)
  }
  buf := bytes.NewBuffer(nil)
  err := gob.NewEncoder(buf).Encode(obj)
  return buf.Bytes(), err
}


func main() {
  var data []byte
  var kv = make(map[interface{}]interface{})
  fmt.Print("Input uname:")
  var uname string
  fmt.Scanf("%s", &uname)
  kv["uname"] = uname
  kv["uid"] = int64(1)
  fmt.Println(kv)

  data, err := EncodeGob(kv);
  if err !=nil {
    fmt.Println(err)
  }
  ioutil.WriteFile("payload", data, 0644)
}
