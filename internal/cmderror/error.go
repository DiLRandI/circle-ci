package cmderror

import "errors"

// LoggerNotFoundError is returned when the logger is not found in the context
var LoggerNotFoundError = errors.New("could not get logger from context")

// CircleCiAPIKeyNotFoundError is returned when the Circle CI API key is not found in the environment
var CircleCiAPIKeyNotFoundError = errors.New("CIRCLE_CI_API_KEY is not set")
