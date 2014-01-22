package hashtable

type HashTable struct {
	values []interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{
		values: make([]interface{}, 0),
	}
}

// Get returns the instance previously stored
// with the specified key. If no instance exists
// at that key, or nil was stored at the key, nil
// is returned.
func (ht HashTable) Get(key string) interface{} {
	i := getIntHash(key)

	if ht.containsHash(i) {
		return ht.values[i]
	} else {
		return nil
	}
}

// Set stores the specified instance at the specified key.
// If an instance already exists for the key, the instance
// is overwritten.
func (ht *HashTable) Set(key string, value interface{}) {
	i := getIntHash(key)

	if lastIndex := len(ht.values) - 1; i > lastIndex {
		newTail := make([]interface{}, i-lastIndex)
		ht.values = append(ht.values, newTail...)
	}

	ht.values[i] = value
}

func (ht HashTable) containsHash(i int) bool {
	return i < len(ht.values)
}

func getIntHash(s string) int {
	hash := 0
	for _, i := range s {
		hash = (31*hash + int(i)) % 10
	}

	return hash
}
