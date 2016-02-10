maptostruct
===
[![GoDoc](https://godoc.org/github.com/Iwark/maptostruct?status.svg)](https://godoc.org/github.com/Iwark/maptostruct)

Package ``maptostruct`` convert a map to a struct.

## Installation

```
$ go get github.com/Iwark/maptostruct
```

## Example

```go
package main

import (
  "fmt"
  "github.com/Iwark/maptostruct"
)

type MyStruct struct {
  Name string
  Age  int64
}

func main() {
    myData := make(map[string]interface{})
    myData["Name"] = "Tony"
    myData["Age"] = int64(23)

    result := &MyStruct{}
    err := maptostruct.Do(myData, result)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}
```

## License

maptostruct is released under the MIT License.