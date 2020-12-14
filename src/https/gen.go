package https

import (
	"fmt"
	"strings"
)

func GenUrl(uri string, ps map[string]string) string {
	var paramsList []string

	for k, v := range ps {
		paramsList = append(paramsList, fmt.Sprintf("%s=%s", k, v))
	}

	var url = uri
	if !strings.Contains(uri, "?") {
		url = fmt.Sprintf("%s?", url)
	}

	url = fmt.Sprintf("%s%s", url, strings.Join(paramsList, "&"))
	return url
}
