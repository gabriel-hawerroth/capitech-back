package dto

type SaveSearchLogDTO struct {
	FieldKey   string `json:"field_key"`
	FieldValue string `json:"field_value"`
}

type SaveSearchLogWithUserDTO struct {
	UserId     int    `json:"user_id"`
	FieldKey   string `json:"field_key"`
	FieldValue string `json:"field_value"`
}
