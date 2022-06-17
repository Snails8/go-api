package domain

type User struct {
	Id   int32
	Name string
}

func NewUser(id int32, name string) (User, error) {
	userErrors := UserErrors{}
	userErrors = validation(userErrors, id, name)

	if len(userErrors.Items) > 0 {
		return User{}, userErrors
	}

	return User{
		Id: id,
		Name: name,
	}, nil
}

func validation(userErrors UserErrors, id int32, name string) UserErrors {
	// 文字数ではなく文字列 判定
	if len(name) > 50 && len(name) < 0 {
		userErrors.Items = append(userErrors.Items, ErrName)
	}

	return userErrors
}
