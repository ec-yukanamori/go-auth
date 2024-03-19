package main

type (
	generateTokenInput struct {
		AppURI string   `json:"app_uri" validate:"required,uri"`
		UserID string   `json:"user_id" validate:"required"`
		Roles  []string `json:"roles"`
	}
)
