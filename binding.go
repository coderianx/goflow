package goflow

import (
	"encoding/json"
	"errors"
	"io"
)

func (c *Ctx) ShouldBindJSON(dst any) error {
	err := json.NewDecoder(c.Request.Body).Decode(dst)

	if errors.Is(err, io.EOF) {
		return errors.New("goflow: request body is empty")
	}

	return err
}
