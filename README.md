# Package to create simple translations for a project


Usage

```go
import (
    "fmt"
    "github.com/hanspr/lang"
)

// Global to be used anywhere in the project

var (
    Lang *lang.Lang
)

func main() {
    y := 10
    x := 2
    Lang = lang.NewLang("PATHT/es_MX.lang")
    if x==0 {
        fmt.Print(Lang.Translate("x can not have a value of zero"))
    }
    c:= y / x
    fmt.Printf(Lang.Translate("Division of %d / %d, is equal to: %d"),y,x,c)
}

```

es_MX.lang
```text
x can not have a value of zero|x no puede tener valor de cero
Division of %d / %d, is equal to: %d|División de %d / %d, es igual a: %d
```
