package apiwire

type RegisterPayload struct{ Name string }

type AuthenticatePayload struct {
	LoginToken              string
	GameVersion             int64
	ApkHash                 string
	ApkApplicationSignature string
	ApplicationVersion      string
}

func DecodeRegisterPayload(arr []any) RegisterPayload { return RegisterPayload{Name: StringAt(arr, 0)} }

func DecodeAuthenticatePayload(arr []any) AuthenticatePayload {
	return AuthenticatePayload{LoginToken: StringAt(arr, 0), GameVersion: Int64At(arr, 1), ApkHash: StringAt(arr, 2), ApkApplicationSignature: StringAt(arr, 3), ApplicationVersion: StringAt(arr, 4)}
}

func AccountRegistResult(token string, errorType int64) []any { return []any{token, errorType} }

func AuthenticateResult(token string, banLevel int64, warnedUntil any) []any {
	return []any{token, banLevel, warnedUntil}
}
