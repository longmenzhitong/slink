package sids

import (
	"fmt"
	"slink/src/rds"
)

func Gen() (int64, error) {
	cmd := rds.Client.Incr("slink:sid")
	id, err := cmd.Result()
	if err != nil {
		return -1, fmt.Errorf("get result of %v error: %v", cmd, err)
	}
	return id, nil
}
