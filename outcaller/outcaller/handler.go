package function

import (
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	resp, _ := http.Get("http://google.com/")
	return fmt.Sprintf("Outcaller status: %s", resp.Status)
}
