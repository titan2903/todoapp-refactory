package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

type UserFetchFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}

func FormatFetchUser(user User) UserFetchFormatter {
	formatter := UserFetchFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
	}

	return formatter
}