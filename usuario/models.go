package usuario

type Usuario struct {
	Usuario_id       string `json:"usuario_id"`
	Usuario_nombre   string `json:"nombre"`
	Usuario_email    string `json:"email"`
	Usuario_password string `json:"password,omitempty"`
	Usuario_avatar   string `json:"avatar,omitempty"`
}
