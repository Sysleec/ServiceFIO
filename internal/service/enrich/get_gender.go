package enrich

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type genderizeResponse struct {
	Gender string `json:"gender"`
}

func GetGender(ctx context.Context, name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ar genderizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return "", err
	}

	return ar.Gender, nil
}
