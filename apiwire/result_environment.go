package apiwire

func EnvironmentResult(applicationVersion string, assetVersion string, apiEndpoint string, maintenanceApiEndpoint string, newsApiEndpoint string, isMaintenance bool, masterDataURL string, staticContentURL string, assetURL string, isAppReview bool, photoContentURL string, multiRealTimeServerURL string, externalPaymentURL string) []any {
	return []any{applicationVersion, assetVersion, apiEndpoint, maintenanceApiEndpoint, newsApiEndpoint, isMaintenance, masterDataURL, staticContentURL, assetURL, isAppReview, photoContentURL, multiRealTimeServerURL, externalPaymentURL}
}
