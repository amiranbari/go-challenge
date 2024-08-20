package queryparam

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func GetFilterParams(c echo.Context) map[string]any {
	queryParams := make(map[string]any)
	for key, values := range c.QueryParams() {
		if len(values) > 0 && strings.HasPrefix(key, "filter_") {
			queryParams[strings.TrimPrefix(key, "filter_")] = values[0]
		}
	}

	return queryParams
}
