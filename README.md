# gin-config
Using json for configrations on gin application.

### Usage

```sh
    go get github.com/virskor/gin-config
```

### Config file

Save config.json file as debug/release .json file in defined root path.
``` go
    configs/app/*.json
```
Filename is same with your gin.Mode().

debug.json
```json
    {
        "application": {
            "name": "example"
        },
        "db": {
            "driver": "mysql",
            "password": "",
            ....
        }
    }
```

### Example

```go
import (
    ginc "github.com/virskor/gin-config"
    "fmt"
)

var conf *ginc.Config

func main(){
    applicationName:= ginc.Get(&ginc.ConfigOptions(Key: "application.name"))

    fmt.Println("result", applicationName)
}
```

ginc.ConfigOptions

```go
    type ConfigOptions struct {
        Key  string //important, grammar would like gson
        File string
    }
```
