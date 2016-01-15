package php_test

import (
	"fmt"
	"github.com/ha1t/go-php-function"
	"testing"
)

func ExampleBasename() {
	fmt.Println(php.Basename("/usr/bin/hogehoge.php"))
	// Output: hogehoge.php
}

func TestBasename(t *testing.T) {
	if i := php.Basename("/hoge/huga/base.php"); i != "base.php" {
		t.Error("failed")
	} else {
		t.Log("ok")
	}
}
