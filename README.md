# cache

Домашнее задание №1 для курса https://www.udemy.com/course/golang-ninja

go get -u github.com/mdotv/cache

## Пример

```go
package main

import (
	"github.com/mdotv/cache"
	"fmt"
)

func main() {
	c := cache.New()

	c.Set("userId", 42)
	userId := c.Get("userId")

	fmt.Println(userId)

	c.Delete("userId")
	userId = c.Get("userId")

	fmt.Println(userId)
}
```