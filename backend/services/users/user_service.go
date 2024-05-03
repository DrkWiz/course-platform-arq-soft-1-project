package users

import "backend/domain/users"

func Login(request users.LoginRequest) users.LoginResponse {
	// Validar contra la DB
	return users.LoginResponse{
		Token: "asdasdasd",
	}
}
