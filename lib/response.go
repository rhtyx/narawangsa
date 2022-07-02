package lib

func Response(status string, message string, obj interface{}, accessToken *string) interface{} {
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

	return res
}
