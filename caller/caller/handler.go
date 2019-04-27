package function

import (
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req http.Request) ([]byte, error) {
	resp, err := http.Get(req.Header.Get("X-URL"))
	if err != nil {
		return []byte(err.Error()), err
	}
	defer resp.Body.Close()

	return []byte(fmt.Sprintf("Caller status: %s", resp.Status)), nil
}
