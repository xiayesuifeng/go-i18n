# go-i18n
> Golang's i18n implementation, provides internationalization support for your application

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/xiayesuifeng/go-i18n)](https://goreportcard.com/report/gitlab.com/xiayesuifeng/go-i18n)
[![GoDoc](https://godoc.org/gitlab.com/xiayesuifeng/go-i18n?status.svg)](https://godoc.org/gitlab.com/firerainos/xiayesuifeng/go-i18n)
[![Sourcegraph](https://sourcegraph.com/gitlab.com/xiayesuifeng/go-i18n/-/badge.svg)](https://sourcegraph.com/gitlab.com/xiayesuifeng/go-i18n)

## Example Code
`zh_CN.go`
```go
package main

import (
    "gitlab.com/xiayesuifeng/go-i18n"
)

func init() {
    i18n.AddTranslation(i18n.Translation{
        "Hello":            "你好",
    }, "zh_CN")
}
```
`en_US.go`
```go
package main

import (
    "gitlab.com/xiayesuifeng/go-i18n"
)

func init() {
    i18n.AddTranslation(i18n.Translation{
        "Hello":            "Hello",
    }, "en_US")
}
```

`main.go`
```go
package main

import (
    "fmt"
    "gitlab.com/xiayesuifeng/go-i18n"
)

func main() { 
    // Set language automatically based on LANG variables
    i18n.LoadTranslator()

    // Get translation by key name, return key if the translation does not exist
    fmt.Println(i18n.Tr("Hello")) // Hello
  	
    // Set language to zh_CN
    i18n.SetLanguage("zh_CN")
    
    // Get translation by key name
    fmt.Println(i18n.Tr("Hello")) // 你好

    ...
}
```

## License
This package is released under [GPLv3](LICENSE).