package cmderror

import "errors"

var LoggerNotFoundError = errors.New("could not get logger from context")
