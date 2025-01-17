package collection

// CopyMap Copy map
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}

func MapDeleteAll(m map[string]interface{}, keys []string) {
	for k, _ := range m {
		delete(m, k)
	}
}
