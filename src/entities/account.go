package entities

type Account struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	State      int    `json:"state"`
	Created_at string `json:"created_at"`
}

type AccountWithToken struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	State      int    `json:"state"`
	Token      string `json:"token"`
	Created_at string `json:"created_at"`
}
