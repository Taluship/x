package pclient

type HealthData struct {
	CanBackup bool `json:"canBackup"`
}

type HealthResponse struct {
	BaseResponse
	Data HealthData `json:"data"`
}

func (c *Client) Health() (*HealthResponse, error) {
	resp, err := c.get("/api/health")
	if err != nil {
		return nil, err
	}

	var data HealthResponse
	err = c.decode(resp, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
