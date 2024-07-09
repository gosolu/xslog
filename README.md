# xslog
Extend golang/slog with some utility functions.


### Usage
This library requires **go 1.21.0**.

```golang
package main

import (
    "log/slog"
    "os"

    "github.com/gosolu/xslog"
)

func main() {
    handler := xslog.NewHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
    slog.SetDefault(slog.New(handler))
}
```

### LICENSE
MIT LICENSE

