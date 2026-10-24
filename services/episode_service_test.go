package services

import (
	"context"
	"testing"
)

func TestEpisodeServiceFallbackShape(t *testing.T) {
	svc := NewEpisodeService(nil, nil, "")
	payload, err := svc.GetEpisodeDetail(context.Background(), 140271)
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) != 4 {
		t.Fatalf("EpisodeResult len = %d", len(payload))
	}
	if payload[0] != "Episode" {
		t.Fatalf("fallback title = %#v", payload[0])
	}
	if payload[3] != "episode-details/140271" {
		t.Fatalf("fallback asset source = %#v", payload[3])
	}
}
