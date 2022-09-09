package requests

import "encoding/json"

func UnmarshalData(data interface{}, out interface{}) error {
	if data == nil {
		return nil
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}
