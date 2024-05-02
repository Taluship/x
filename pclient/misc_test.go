package pclient

import (
	"reflect"
	"testing"
)

func TestClient_Health(t *testing.T) {
	expectedResp := &HealthResponse{
		BaseResponse: BaseResponse{
			Code:    200,
			Message: "API is healthy.",
		},
		Data: HealthData{
			CanBackup: true,
		},
	}
	c := NewClient()

	healthResp, err := c.Health()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(healthResp, expectedResp) {
		t.Fatalf("unexpected response: %+v, expected: %+v", healthResp, expectedResp)
	}
}
