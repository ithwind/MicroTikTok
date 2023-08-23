package AIChat

import (
	"fmt"
	"testing"
)

func TestChat(t *testing.T) {
	chat := Chat("你好")

	fmt.Println(chat)
}
