package helper

import (
	Entity "entity"
	"encoding/json"
)

func OutputJson(code int, msg string, data interface{}) string {
	Resp := Entity.Response{code, msg, data}
	str, err := json.Marshal(Resp)
	if err != nil {
		return ""
	}
	return string(str)
}
