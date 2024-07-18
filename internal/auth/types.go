package auth

type registerDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password_strength"`
}

type loginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password_strength"`
}

type session struct {
  SID    string `json:"sid"`
  IP     string `json:"ip"`
  Login  string `json:"login"`
  Expiry string `json:"expiry"`
  UA     string `json:"ua"`
}
