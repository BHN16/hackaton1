package models

type User struct {
	Email string `json:"Email"`

	Password string `json:"Password"`

	Role string `json:"Role"`
}

type Token struct {
	Role string `json:"role"`

	Email string `json:"email"`

	TokenString string `json:"token"`
}

type UserAuthenticate struct {
	Email string `json: "Email"`

	Password string `json: "Password"`

	TokenString string `json: "TokenString"`
}

type TemporalUser struct {
	Name string `json:"Name"`

	Role string `json:"Role"`

	Email string `json:"Email"`

	Password string `json:"Password"`

	Codigo string `json:"Codigo"`
}
