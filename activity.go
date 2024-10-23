package app

import (
	"context"
	"fmt"
)

func ComposeTest(ctx context.Context, msg string) (string, error) {
	test_res := fmt.Sprintf("Test %s!", msg)
	return test_res, nil
}

