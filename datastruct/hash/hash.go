package hash

// the implementation of hash table.

type (
	// Hash hash table struct.
	Hash struct {
		record Record
	}

	// Record hash records to save.
	Record map[string]map[string][]byte
)

// New create a new hash datastruct.
func New() *Hash {
	return &Hash{make(Record)}
}

// HSet设置存储在键到值处的散列字段。如果key不存在，则创建一个包含散列的新key。
// 如果字段在散列中已经存在，则会覆盖它。
func (h *Hash) HSet(key string, field string, value []byte) (res int) {
	if !h.exist(key) {
		h.record[key] = make(map[string][]byte)
	}

	if h.record[key][field] != nil {
		//如果该字段存在，则覆盖它。
		h.record[key][field] = value
	} else {
		// 如果此字段不存在，则创建。
		h.record[key][field] = value
		res = 1
	}
	return
}

func (h *Hash) HSetNx(key string, field string, value []byte) int {
	if !h.exist(key) {
		h.record[key] = make(map[string][]byte)
	}

	if _, exist := h.record[key][field]; !exist {
		h.record[key][field] = value
		return 1
	}
	return 0
}

// 返回与存储在键的散列中的字段相关联的值。
func (h *Hash) HGet(key, field string) []byte {
	if !h.exist(key) {
		return nil
	}

	return h.record[key][field]
}

// HGetAll返回存储在key中的所有字段和哈希值。
// 在返回值中，每个字段名后面都跟着它的值，因此应答的长度是散列大小的两倍。
func (h *Hash) HGetAll(key string) (res [][]byte) {
	if !h.exist(key) {
		return
	}

	for k, v := range h.record[key] {
		res = append(res, []byte(k), v)
	}
	return
}

// HDel从存储在key的散列中删除指定的字段。在此散列中不存在的指定字段将被忽略。
// 如果key不存在，它将被视为空散列，该命令返回false。
func (h *Hash) HDel(key, field string) int {
	if !h.exist(key) {
		return 0
	}

	if _, exist := h.record[key][field]; exist {
		delete(h.record[key], field)
		return 1
	}
	return 0
}

// 如果键存在于散列中则返回。
func (h *Hash) HKeyExists(key string) bool {
	return h.exist(key)
}

// 如果field是存储在key中的散列中已有的字段，则返回。
func (h *Hash) HExists(key, field string) (ok bool) {
	if !h.exist(key) {
		return
	}

	if _, exist := h.record[key][field]; exist {
		ok = true
	}
	return
}

// 返回存储在key的散列中包含的字段的数量。
func (h *Hash) HLen(key string) int {
	if !h.exist(key) {
		return 0
	}
	return len(h.record[key])
}

// 返回存储在key的散列中的所有字段名。
func (h *Hash) HKeys(key string) (val []string) {
	if !h.exist(key) {
		return
	}

	for k := range h.record[key] {
		val = append(val, k)
	}
	return
}

// HVals返回存储在key的散列中的所有值。
func (h *Hash) HVals(key string) (val [][]byte) {
	if !h.exist(key) {
		return
	}

	for _, v := range h.record[key] {
		val = append(val, v)
	}
	return
}

func (h *Hash) HClear(key string) {
	if !h.exist(key) {
		return
	}
	delete(h.record, key)
}

func (h *Hash) exist(key string) bool {
	_, exist := h.record[key]
	return exist
}
