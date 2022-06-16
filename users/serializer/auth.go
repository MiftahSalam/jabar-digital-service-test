package serializer

type AuthDtoResponse struct {
	Is_valid bool   `json:"is_valid"`
	Expired  string `json:"expired_at"`
	Username string `json:"username"`
}
