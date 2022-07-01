package lib

func Response(status string, message string, obj interface{}) interface{} {
	res := map[string]interface{}{
		"status":  status,
		"message": message,
	}

	if obj != nil {
		res["data"] = obj
	}

	return res
}
