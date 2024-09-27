package hashing

import (
	"crypto/sha256"
	"fmt"
)

//we may change hashing algorythm so it makes sense to put it in it's own package.
func Make(str string) string{
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
