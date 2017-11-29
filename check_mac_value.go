package ecpay

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
)

// CheckMacValue ...
func CheckMacValue(str string) string {
	encoded := dotNetURLEncode(str)
	lower := strings.ToLower(encoded)
	shabyte := sha256.Sum256([]byte(lower))
	result := fmt.Sprintf("%x", shabyte)
	return strings.ToUpper(result)
}

func dotNetURLEncode(str string) string {
	replacer := strings.NewReplacer("%2521", "!", "%252a", "*", "%2528", "(", "%2529", ")")
	return replacer.Replace(url.QueryEscape(str))
}
