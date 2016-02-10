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

type Person struct {
  Name string
  Age  int64
  Gender string `mts:"gender"`
}

func main() {
    person := map[string]interface{}{
        "Name":   "Iwark",
        "Age":    24,
        "gender": "man",
        "hobby":  "playing the piano",
    }

    result := &Person{}
    err := maptostruct.Do(person, result)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}
```

## License

maptostruct is released under the MIT License.