package responses

import (
	"encoding/json"
	"fmt"
)

type (
	Basic struct {
		Code string `json:"code"`
		Msg  string `json:"msg,omitempty"`
	}
)

type resCode Basic

func (b *Basic) UnmarshalJSON(bf []byte) error {
	var r resCode
	err := json.Unmarshal(bf, &r)
	if err != nil {
		return err
	}
	b = (*Basic)(&r)
	if b.Code != "0" {
		return fmt.Errorf("recevied error:%+v", b)
	}
	return nil
}
