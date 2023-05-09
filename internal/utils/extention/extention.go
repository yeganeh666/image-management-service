package extention

import (
	"fmt"
	"net/http"
	"strings"
)

func GetFileExtension(resp *http.Response) string {
	// Detect the content type of the response body
	contentType := resp.Header.Get("Content-Type")

	// Extract the MIME type from the content type
	mimeType := strings.Split(contentType, "/")[1]

	return fmt.Sprintf(".%v", mimeType)
}
