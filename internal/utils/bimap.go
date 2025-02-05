package utils

type BiMap struct {
	Forward map[string]string
	Reverse map[string]string
}

func NewBiMap() *BiMap {
	return &BiMap{
		Forward: make(map[string]string),
		Reverse: make(map[string]string),
	}
}
