package repositories

type AuthReposiory interface {
	Register(email, password string) (string, error)
	Login(email, password string) (string, error)
}
