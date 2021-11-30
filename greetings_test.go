package greetings

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	if res := Hello("zhangsan"); res != "Hi, zhangsan. Welcome!" {
		t.Fatal("fail")
	}

	if res := Hello("李四"); !strings.EqualFold(res, "Hi, 李四. Welcome!") {
		t.Fatal("fail")
	}
}
