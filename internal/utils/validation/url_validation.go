package validation

import "net/url"

// Check if a given URL is a valid image URL
func IsValidImageURL(imageUrl string) bool {
	_, err := url.ParseRequestURI(imageUrl)
	if err != nil {
		return false
	}
	return true
}
