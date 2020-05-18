# reading-time

`reading-time` is the GoLang tool that helps you estimate how long an article will take to read. 
It works perfectly with plain text, but also with `markdown` or `html`.

Note that it's focused on performance and simplicity, so the number of words it will extract from other formats than plain text can vary a little. But this is an estimation right?

## Installation

```bash
go get github.com/begmaroman/reading-time
```

## Usage

```go
package main

import (
    "fmt"
 
    readingtime "github.com/begmaroman/reading-time"
)

func main()  {
	myArticle := "some long long long text with ##markdown stuff and <b>html</b>"
	estimation := readingtime.Estimate(myArticle)
	fmt.Println(estimation.Text)        // "1 min read"
	fmt.Println(estimation.Duration)    // 1 min
	fmt.Println(estimation.Words)       // 10
}

```

