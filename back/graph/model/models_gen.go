// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CompletedstatusInput struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginData struct {
	Token    string    `json:"token"`
	Userdata *UserData `json:"userdata"`
}

type NewTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type UserData struct {
	Key       string `json:"_key"`
	ID        string `json:"_id"`
	Rev       string `json:"rev"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
