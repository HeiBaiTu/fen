### Learning and testing usage


```go
package main

import (
  "net/http"

  "github.com/HeiBaiTu/fen"
)

func main() {
  r := fen.Default()
  r.GET("/", func(c *fen.Context) {
    c.JSON(http.StatusOK, fen.H{"message": "hello world"})
  })
  r.Run() 
}
```

```bash
2024/01/03 16:01:26 Fen Listening and serving HTTP on :8080
```
