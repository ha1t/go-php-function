# go-php-function

## basename

```go
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

```go
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

```go
package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Split("hello_world", "_")[0])
}
```

## in_array

```php
<?php
if (in_array('hoge', ['hoge', 'huga'])) {
    echo "ok";
}
```

```go
package main

import (
	"fmt"
	"github.com/ha1t/go-php-function"
)

func main() {
	itemList := []string{"hoge", "huga"}
	if php.InArray("hoge", itemList) {
		fmt.Println("ok")
	}
}
```
