package lib

func Response(status string, message string, obj interface{}, accessToken *string, refreshToken *string) interface{} {
	res := map[string]interface{}{
		"status":  status,
		"message": message,
	}

	if obj != nil {
		res["data"] = obj
	}

	if accessToken != nil {
		res["access_token"] = accessToken
	}

	if refreshToken != nil {
		res["refresh_token"] = refreshToken
	}

	return res
}
