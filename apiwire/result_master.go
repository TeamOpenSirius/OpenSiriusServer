package apiwire

func MasterDataManifest(uri string, sasToken string, version string, publishTimestamp int64) []any {
	return []any{uri, sasToken, version, publishTimestamp}
}
