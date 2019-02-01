# pager

[![Build Status](https://travis-ci.org/visig9/pager.svg?branch=master)](https://travis-ci.org/visig9/pager)

A lightweight golang pager for any slice datatype.



## Usage

```go
package main

import (
    "fmt"
    "gitlab.com/visig/pager"
)

func main() {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	p := pager.Pager{
		Items:    data,
		PageSize: 3,
	}

	fmt.Println("Page Count: " + p.PageCount()) // 4
	fmt.Println("Page 2: " + p.RawPage(2))      // [d e f]
}
```

Check `godoc` for more API.

## Download

```bash
go get gitlab.com/visig/pager
```


## License

MIT
