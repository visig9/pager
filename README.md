# pager

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

	fmt.Println(p.Page(2)) // with meta data
}
```

## Download

```bash
go get gitlab.com/visig/pager
```


## License

MIT
