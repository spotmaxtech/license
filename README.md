A tiny license tool package

# Usage
```bash
go get github.com/spotmaxtech/license
```

```go
package main

import "github.com/spotmaxtech/license"

func main() {
	l := license.NewLicenseManger("123456789")
	key, _ := l.CreateLicense([]byte("jinjin"))
}

```
