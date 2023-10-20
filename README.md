go-slugify
==============

[![Go Report Card](https://goreportcard.com/badge/github.com/kong/go-slugify)](https://goreportcard.com/report/github.com/kong/go-slugify)

Make Pretty Slugs.

Installation
------------

```
go get -u github.com/kong/go-slugify
```

Install CLI tool:

```
go get -u github.com/kong/go-slugify/slugify
$ slugify "北京kožušček,abc"
bei-jing-kozuscek-abc
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/kong/go-slugify


Usage
------

```go
package main

import (
	"fmt"
	"github.com/kong/go-slugify"
)

func main() {
	slugifier := sligify.NewSlugifier()
	s := "北京kožušček,abc"
	fmt.Println(slugifier.Slugify(s))
	// Output: bei-jing-kozuscek-abc
}
```
