package errors

import (
	"fmt"
)

var ErrProductNotFound error = fmt.Errorf("product Not found")
var ErrInvalidURI error = fmt.Errorf("invalid uri")
var ErrUnmarshal error = fmt.Errorf("can't unmarshal given json")