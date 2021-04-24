package dto

type UserDTO struct {
	ID        string `db:"id"`
	UserName  string `db:"user_name"`
	Email     string `db:"email"`
	BirthDate string `db:"birth_date"`
}
