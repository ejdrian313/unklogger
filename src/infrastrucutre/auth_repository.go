package infrastrucutre

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) *GormAuthRepository {
	return &GormAuthRepository{db: db}
}

func (repo *GormAuthRepository) Register(username, password string) (string, error) {
	userExists := repo.db.Where("username = ?", username).First(&UserAuth{})
	if userExists.Error == nil {
		return "", errors.New("User already exists")
	}
	user := User{
		Name: username,
	}
	repo.db.Create(&user)
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	repo.db.Create(&UserAuth{
		UserID:    user.ID,
		Username:  username,
		Password:  string(bcryptPassword),
		CreatedAt: time.Now(),
	})

	tokenString, err := generateToken(username, user.ID)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (repo *GormAuthRepository) Login(username, password string) (string, error) {
	var userAuth UserAuth
	userExists := repo.db.Where("username = ?", username).First(&userAuth)
	if userExists.Error != nil {
		return "", userExists.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(password))
	if err != nil {
		return "", err
	}
	tokenString, err := generateToken(username, userAuth.UserID)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generateToken(username string, ID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"ID":       ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
