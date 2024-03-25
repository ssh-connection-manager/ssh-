package file

import (
	"encoding/json"
)

type Connections struct {
	Connects []Connect `json:"connects"`
}

type Connect struct {
	Alias    string `json:"alias"`
	Login    string `json:"login"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (c *Connections) UnmarshalJSON(data []byte) error {
	type connections Connections
	var result connections

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	c.Connects = result.Connects
	return nil
}

func (c *Connections) GetConnectionsAlias() []string {
	var result []string

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	return result
}
