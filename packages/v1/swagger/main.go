package main

import (
	_types "docs/internal/types"
	_ "embed"
	"fmt"
	"lib/constants"
	"net/http"
)

//go:embed swagger.json
var swaggerJSON string

// @title Kibisis API
// @version 1.0
// @description The Kibisis API provides analytics queries and token data.

// @contact.name   Agora Labs Support
// @contact.email  support@agoralabs.sh

// @license.name  GNU AGPLv3+
// @license.url   https://github.com/kibis-is/api/blob/main/COPYING
func Main() *_types.Response {
	return &_types.Response{
		Body: swaggerJSON,
		Headers: _types.ResponseHeaders{
			CacheControl: fmt.Sprintf("public, max-age=%d", constants.TwentyFourHoursInSeconds),
			ContentType:  "application/json",
		},
		StatusCode: http.StatusOK,
	}
}
