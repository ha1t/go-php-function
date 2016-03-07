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

func TestInArray(t *testing.T) {
	itemList := []string{"hoge", "huga", "moge"}

	if php.InArray("hoge", itemList) {
		t.Log("ok")
	} else {
		t.Error("failed")
	}

	if php.InArray("not_found", itemList) {
		t.Error("failed")
	} else {
		t.Log("ok")
	}

	if php.InArray("", itemList) {
		t.Error("failed")
	} else {
		t.Log("ok")
	}
}
