package apiwire

func ArrayByKey(size int, values map[int]any) []any {
	out := make([]any, size)
	for key, value := range values {
		if key >= 0 && key < size {
			out[key] = value
		}
	}
	return out
}

func DataObject(unionID int64, arr []any) []any { return []any{unionID, arr} }

func BooleanResult(success bool) []any { return []any{success} }
