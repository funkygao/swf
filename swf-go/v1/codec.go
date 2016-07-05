package swfapi

type Encoder interface {
	Marshal(v interface{}) ([]byte, error)
}

type Decoder interface {
	Unmarshal(data []byte, v interface{}) error
}
