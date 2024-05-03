package users

import usersDomain "backend/domain/users"

func Login(request usersDomain.LoginRequest) usersDomain.LoginResponse {

	//validate with db

	return usersDomain.LoginResponse{Token: "nicotroll123"}
}
