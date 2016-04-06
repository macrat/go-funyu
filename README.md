go-funyu
========

The [funyu markup language](https://bitbucket.org/MacRat/funyu) parser written by golang.

## Installation
``` bash
$ go get https://github.com/macrat/go-funyu.git
```

## Usage
Like this.

``` go
package main

import (
	"fmt"
	"github.com/macrat/go-funyu"
)

func main() {
	fmt.Println(funyu.Parse(`
title: test

this is test of [[funyu]].`).HTML())
}
```

This code will do display like below.

``` HTML
<p>
this is test of <strong>funyu</strong>.<br>
</p>
```

Detail about **funyu**, please see [reference of funyu parser for python](https://bitbucket.org/MacRat/funyu/src/tip/REFERENCE.fny).

## License
[MIT License](https://opensource.org/licenses/MIT) (c)2016 [MacRat](http://blanktar.jp) <m@crat.jp>
