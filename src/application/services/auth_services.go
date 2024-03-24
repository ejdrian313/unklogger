package services

import "unklogger.com/src/domain/repositories"

type AuthServices struct {
	repo repositories.AuthReposiory
}

func NewAuthService(repo repositories.AuthReposiory) *AuthServices {
	return &AuthServices{repo: repo}
}

func (s *AuthServices) Login(username, password string) (string, error) {
	return s.repo.Login(username, password)
}

func (s *AuthServices) Register(username, password string) (string, error) {
	return s.repo.Register(username, password)
}
