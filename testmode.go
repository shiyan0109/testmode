package testmode

import (
    "fmt"
)

func Hi(name string) string {
    return fmt.Sprintf("Hi, %s", name)
	return fmt.Sprintf("Hi v2, %s", name)
}