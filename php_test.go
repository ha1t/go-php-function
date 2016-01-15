package php_test

import (
	"github.com/ha1t/go-php-function"
	"testing"
)

func TestBasename(t *testing.T) {
	if i := php.Basename("/hoge/huga/base.php"); i != "base.php" {
		t.Error("failed")
	} else {
		t.Log("ok")
	}
}
