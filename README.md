
```golang
package main

import (
	"fmt"
	"github.com/ha1t/go-php-function"
)

func main() {
	filepath := "/usr/bin/hogehoge.php"
	fmt.Printf("before %v\n", filepath)
	fmt.Printf("after %v\n", php.Basename(filepath))
}
```

```
before /usr/bin/hogehoge.php
after hogehoge.php
```

## glob

```golang
package main

import (
	fmt "fmt"
	"path/filepath"
)

func main() {

	items, err := filepath.Glob("/home/*")
	if err != nil {
		panic(err)
	}
	for _, item := range items {
		fmt.Printf(item + "\n")
	}
}
```
## explode (split)

```php
<?php
echo explode("_", "hello_world")[0];
```

```golang
package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Split("hello_world", "_")[0])
}
```
