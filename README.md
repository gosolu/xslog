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
    handler := xslog.UseContext(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
    slog.SetDefault(slog.New(handler))

    ctx := context.Background()
    slog.InfoContext(ctx, "Hi")
}
```

### LICENSE
MIT LICENSE

