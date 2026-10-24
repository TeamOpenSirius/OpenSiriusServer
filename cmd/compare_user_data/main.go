package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type summary struct {
	TotalSize   int            `json:"total_size"`
	ObjectCount int            `json:"object_count"`
	UnionCounts map[string]int `json:"union_counts"`
	Warnings    []string       `json:"warnings"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: compare_user_data <official.summary.json> <seed.summary.json>")
		os.Exit(2)
	}

	a, err := readSummary(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	b, err := readSummary(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v\n", os.Args[2], err)
		os.Exit(1)
	}

	fmt.Printf("total_size: official=%d seed=%d delta=%d\n", a.TotalSize, b.TotalSize, b.TotalSize-a.TotalSize)
	fmt.Printf("object_count: official=%d seed=%d delta=%d\n", a.ObjectCount, b.ObjectCount, b.ObjectCount-a.ObjectCount)
	fmt.Println("union_counts:")

	keys := map[string]bool{}
	for key := range a.UnionCounts {
		keys[key] = true
	}
	for key := range b.UnionCounts {
		keys[key] = true
	}
	ordered := make([]string, 0, len(keys))
	for key := range keys {
		ordered = append(ordered, key)
	}
	sort.Slice(ordered, func(i, j int) bool { return ordered[i] < ordered[j] })
	for _, key := range ordered {
		if a.UnionCounts[key] != b.UnionCounts[key] {
			fmt.Printf("  union %s: official=%d seed=%d\n", key, a.UnionCounts[key], b.UnionCounts[key])
		}
	}

	if len(b.Warnings) > 0 {
		fmt.Println("seed warnings:")
		for _, warning := range b.Warnings {
			fmt.Printf("  - %s\n", warning)
		}
	}
}

func readSummary(path string) (summary, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return summary{}, err
	}
	var s summary
	if err := json.Unmarshal(raw, &s); err != nil {
		return summary{}, err
	}
	if s.UnionCounts == nil {
		s.UnionCounts = map[string]int{}
	}
	return s, nil
}
