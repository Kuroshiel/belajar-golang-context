package belajargolangcontext

import (
	"context"
	"fmt"
	"testing"
)

// Membuat Context

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
