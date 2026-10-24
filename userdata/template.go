package userdata

import (
	"encoding/json"
	"os"
)

type UserDataTemplateRegistry struct {
	Templates map[int64][][]any
}

func LoadTemplateRegistry(seedPath string) (*UserDataTemplateRegistry, error) {
	reg := &UserDataTemplateRegistry{
		Templates: make(map[int64][][]any),
	}

	raw, err := os.ReadFile(seedPath)
	if err != nil {
		if os.IsNotExist(err) {
			return reg, nil
		}
		return nil, err
	}

	var decoded capturedUserData
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil, err
	}

	payload := NormalizeCapturedPayload(decoded.Payload)
	for _, raw := range payload {
		unionID, data, ok := SplitDataObject(raw)
		if !ok {
			continue
		}
		reg.Templates[unionID] = append(reg.Templates[unionID], cloneArray(data))
	}
	return reg, nil
}

func (r *UserDataTemplateRegistry) GetTemplate(unionID int64) ([]any, bool) {
	templates, ok := r.Templates[unionID]
	if !ok || len(templates) == 0 {
		return nil, false
	}
	return cloneArray(templates[0]), true
}

func (r *UserDataTemplateRegistry) HasTemplate(unionID int64) bool {
	_, ok := r.Templates[unionID]
	return ok
}

func (r *UserDataTemplateRegistry) TemplateCount() int {
	return len(r.Templates)
}

func cloneArray(src []any) []any {
	if src == nil {
		return nil
	}
	dst := make([]any, len(src))
	for i, v := range src {
		switch val := v.(type) {
		case []any:
			dst[i] = cloneArray(val)
		case map[string]any:
			dst[i] = cloneMap(val)
		default:
			dst[i] = v
		}
	}
	return dst
}

func cloneMap(src map[string]any) map[string]any {
	dst := make(map[string]any, len(src))
	for k, v := range src {
		switch val := v.(type) {
		case []any:
			dst[k] = cloneArray(val)
		case map[string]any:
			dst[k] = cloneMap(val)
		default:
			dst[k] = v
		}
	}
	return dst
}
