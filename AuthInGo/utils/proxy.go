package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	fmt.Println("Proxying to:", target.String())
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Rewrite = func(pr *httputil.ProxyRequest) {
		originalPath := pr.In.URL.Path
		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)
		pr.SetURL(target)
		pr.Out.URL.Path = target.Path + strippedPath

	}
	return proxy.ServeHTTP
}
