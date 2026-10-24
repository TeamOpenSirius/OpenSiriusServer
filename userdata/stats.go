package userdata

type Stats struct {
	ObjectCount  int
	NilSlotCount int
	UnionCounts  map[int64]int
}

func Summarize(payload []any) Stats {
	stats := Stats{ObjectCount: len(payload), UnionCounts: map[int64]int{}}
	for _, raw := range payload {
		if raw == nil {
			stats.NilSlotCount++
			continue
		}
		unionID, _, ok := SplitDataObject(raw)
		if ok {
			stats.UnionCounts[unionID]++
		}
	}
	return stats
}
