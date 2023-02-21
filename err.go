package persistor

import "errors"

// ErrUndefinedName happens when trying to persist data with an undefined name.
var ErrUndefinedName = errors.New("undefined name")
