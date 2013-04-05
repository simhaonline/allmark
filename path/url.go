package path

import (
	"strings"
)

func CombineUrlComponents(baseUrl string, components ...string) string {
	url := StripTrailingUrlDirectorySeperator(baseUrl)

	for _, component := range components {
		url += UrlDirectorySeperator + StripTrailingUrlDirectorySeperator(component)
	}

	return url
}

func StripTrailingUrlDirectorySeperator(urlComponent string) string {

	url := urlComponent

	for strings.LastIndex(urlComponent, UrlDirectorySeperator)+1 == len(url) {
		url = strings.TrimRight(url, UrlDirectorySeperator)
	}

	return url
}

func AddLeadingUrlDirectorySeperator(url string) string {
	newUrl := url

	for strings.Index(newUrl, UrlDirectorySeperator) == 0 {
		newUrl = strings.TrimLeft(newUrl, UrlDirectorySeperator)
	}

	return UrlDirectorySeperator + newUrl
}