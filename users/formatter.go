package users

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url"`
}

func FormatUser(users Users, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         users.ID,
		Name:       users.Name,
		Occupation: users.Occupation,
		Email:      users.Email,
		Token:      token,
		ImageURL:   users.AvatarFileName,
	}

	return formatter
}
