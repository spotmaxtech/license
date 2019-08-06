# License
A tiny license tool package

# Introduction
Create license is very simple. User provide information, we encrypt the information
with MAGIC string. Then use the MAGIC string we can encrypt user information again
to check valid or not. 
Simple, right? Have fun!

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
