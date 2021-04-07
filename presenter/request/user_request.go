package request

type UserRegister struct {
	Email     string `json:"email" validate:"required"`
	Name      string `json:"userName" validate:"required"`
	BirthDate string `json:"birthDate" validate:"required "`
	Password  string `json:"password" validate:"required"`
	Gender    int    `json:"gender" validate:"required, min=1,max=2"`
}
