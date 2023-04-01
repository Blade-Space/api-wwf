# wwf - working with files 📃

Package for working with files

## import to progect

```golang
package main

import (
  wwf "api/wwf/routes"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  api := r.Group("/api/wwf")
  wwf.RegisterRoutes(api)

  r.Run(":3000")
}
```

## `File` methods

- `api/files` -> Get files in directory
- `api/read_file` -> Read file by path
- `api/rename_file` -> Rename file
- `api/delete_file` -> Delete file
- `api/create_file` -> Create file
