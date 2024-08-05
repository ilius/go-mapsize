# MapSize
How much memory is occupied by map.

```txt
go get github.com/520MianXiangDuiXiang520/MapSize
```

```go
package main

import (
	"fmt"
	"github.com/520MianXiangDuiXiang520/MapSize"
)

func main() {
	m := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		m[i] = struct{}{}
	}
	fmt.Println(mapsize.Size(m)) // 1344

	type A struct{}
	m2 := map[int]*A{}
	for i := 0; i < 100; i++ {
		m2[i] = &A{}
	}
	fmt.Println(mapsize.Size(m)) // 1344
}
```