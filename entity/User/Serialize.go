package User

func UserSerializeMap(user User) map[string]interface{} {
	return map[string]interface{}{
		"uuid":      user.Uuid,
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
		"status":    user.Status,
	}
}
