
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

