package Camera

func CameraSerializeMap(camera Camera) map[string]interface{} {
	return map[string]interface{}{
		"id":                   camera.Uuid,
		"title":                camera.Title,
		"baseUrl":              camera.BaseUrl,
		"user":                 camera.User,
		"password":             camera.Password,
		"lastConnectionStatus": camera.LastConnectionStatus,
		"status":               camera.Status,
	}
}
