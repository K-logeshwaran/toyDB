package Driver

// import (
// 	"encoding/json"
// 	"log"
// )

// type Wrapper map[string]interface{}

// func BuildWrapper(Data []byte) *Wrapper {
// 	W := Wrapper{}
// 	json.Unmarshal(Data, &W)
// 	return &W
// }

// func (w *Wrapper) AddField(name string, value interface{}) *Wrapper {
// 	w.Data[name] = value

// 	return w
// }

// func (w *Wrapper) Value() map[string]interface{} {
// 	return w
// }
// func (w *Wrapper) ToBytes() []byte {

// 	b, err := json.Marshal(w.Data)
// 	if err != nil {
// 		panic(" Error.................")
// 	}

// 	return b
// }

// func (w *Wrapper) ToJson() string {
// 	return string(w.ToBytes())
// }

// func WrapperArrayToBytes(W []Wrapper) []byte {
// 	x, err := json.Marshal(W)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return x
// }

import (
	"encoding/json"
	"log"
)

type Wrapper map[string]interface{}

func BuildWrapper(Data []byte) *Wrapper {
	var W Wrapper // Initialize the Wrapper properly
	json.Unmarshal(Data, &W)
	return &W
}

func (w *Wrapper) AddField(name string, value interface{}) *Wrapper {
	if (*w)[name] != nil {
		return w
	}
	(*w)[name] = value // Use w directly as it's a map

	return w
}

func (w *Wrapper) Update(name string, value interface{}) *Wrapper {

	(*w)[name] = value // Use w directly as it's a map

	return w
}

func (w *Wrapper) Value() map[string]interface{} {
	return *w
}

func (w *Wrapper) ToBytes() []byte {
	b, err := json.Marshal(w)
	if err != nil {
		panic("Error while marshaling Wrapper to JSON")
	}
	return b
}

func (w *Wrapper) ToJson() string {
	return string(w.ToBytes())
}

func WrapperArrayToBytes(W []Wrapper) []byte {
	x, err := json.Marshal(W)
	if err != nil {
		log.Fatal(err)
	}
	return x
}
