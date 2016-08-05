[![Build Status](https://travis-ci.org/macrat/go-funyu.svg?branch=master)](https://travis-ci.org/macrat/go-funyu)

go-funyu
========

The [funyu markup language](https://bitbucket.org/MacRat/funyu) parser written by golang.

## Installation
``` bash
$ go get github.com/macrat/go-funyu
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
	metadata, document, err := funyu.Parse(`
title: test

this is test of [[funyu]].`)

	fmt.Println(document)
}
```

This code will do display like below.

``` HTML
<article>
<p>
this is test of <strong>funyu</strong>.<br>
</p>
</article>
```

Detail about **funyu**, please see [reference of funyu parser for python](https://bitbucket.org/MacRat/funyu/src/tip/REFERENCE.fny).

## License
[MIT License](https://opensource.org/licenses/MIT) (c)2016 [MacRat](http://blanktar.jp) &lt;m@crat.jp&gt;
