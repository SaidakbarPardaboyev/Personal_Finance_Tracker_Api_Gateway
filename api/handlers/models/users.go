package models

type Response struct {
	StatisCode  int
	Description string
	Data        interface{}
}

type UserInfoFromToken struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	UserRole  string `json:"user_role"`
	CreatedAt string `json:"created_at"`
}

// type User struct {
// 	Id             string `json:"id"`
// 	Username       string `json:"username"`
// 	Email          string `json:"email"`
// 	FullName       string `json:"full_name"`
// 	NativeLanguage string `json:"native_language"`
// 	CreatedAt      string `json:"created_at"`
// }

type UpdateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type UpdateProfileResponce struct {
// 	Id             string `json:"id"`
// 	Username       string `json:"username"`
// 	Email          string `json:"email"`
// 	FullName       string `json:"full_name"`
// 	NativeLanguage string `json:"native_language"`
// 	UpdatedAt      string `json:"updated_at"`
// }

type ChangePassword struct {
	NewPassword     string `json:"new_password"`
	CurrentPassword string `json:"current_password"`
}

type ChangeUserRole struct {
	NewUserRole     string `json:"new_user_role"`
}

// // type ForgotPasswordReq struct {
// // 	Email string `json:"email"`
// // }

// // type ResetPasswordReq struct {
// // 	Code        string `json:"code"`
// // 	NewPassword string `json:"new_password"`
// // 	UserId      string `json:"user_id"`
// // 	Email       string `json:"email"`
// // }

// type PrimaryKey struct {
// 	Id string `json:""`
// }

// type Message struct {
// 	Message string `json:"message"`
// }
