package Driver

import "encoding/json"

type Wrapper struct {
	data map[string]interface{}
}

func BuildWrapper(data []byte) *Wrapper {
	W := Wrapper{}
	json.Unmarshal(data, &W.data)
	return &W
}

func (w *Wrapper) AddField(name string, value interface{}) *Wrapper {
	w.data[name] = value

	return w
}

func (w *Wrapper) Value() map[string]interface{} {
	return w.data
}
func (w *Wrapper) ToBytes() []byte {

	b, err := json.Marshal(w.data)
	if err != nil {
		panic(" Error.................")
	}

	return b
}

func (w *Wrapper) ToJson() string {
	return string(w.ToBytes())
}
