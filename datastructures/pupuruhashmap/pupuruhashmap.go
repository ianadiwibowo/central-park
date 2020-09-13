// PupuruHashmap is reinventing the wheel of a hashmap
package pupuruhashmap

type PupuruHashmapInterface interface {
	Set(key string, value string) error
	Get(key string) (value string)
}

const (
	hashmapSize = 10
)

type PupuruHashmap struct {
	Data [hashmapSize]*PupuruHashmapNode
}

type PupuruHashmapNode struct {
	Key   string
	Value string
	Next  *PupuruHashmapNode
}

func NewPupuruHashmap() *PupuruHashmap {
	return &PupuruHashmap{}
}

func (h *PupuruHashmap) Set(key string, value string) {
	index := mapKeyToIndex(key)
	newNode := &PupuruHashmapNode{key, value, nil}

	if h.Data[index] == nil {
		h.Data[index] = newNode
		return
	}

	currentNode := h.Data[index]
	for {
		if currentNode.Key == key {
			currentNode.Value = value
			return
		} else if currentNode.Next == nil {
			currentNode.Next = newNode
			return
		}
		currentNode = currentNode.Next
	}
}

func (h *PupuruHashmap) Get(key string) (value string) {
	index := mapKeyToIndex(key)

	if h.Data[index] == nil {
		return ""
	}

	currentNode := h.Data[index]
	for {
		if currentNode.Key == key {
			return currentNode.Value
		}
		currentNode = currentNode.Next
	}
}

// In this fake hash function, we just take the first chart of the key
// and convert it into int. In real life, this will use a real hash function
// like SipHash (or MD5, SHA1).
func mapKeyToIndex(key string) int {
	hashCode := int(key[0])
	return hashCode % hashmapSize
}
