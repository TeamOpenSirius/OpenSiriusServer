package apiwire

func EpisodeResult(title string, storyType int64, episodeOrder int64, episodeDetailAssetSource string) []any {
	return []any{title, int32(storyType), int32(episodeOrder), episodeDetailAssetSource}
}
