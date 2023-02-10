package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/model"

	"github.com/linothomas14/exercise-course-api/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email string, password string) (string, error)
	CreateUser(user param.Register) (model.User, error)
	FindByEmail(email string) model.User
	IsDuplicateEmail(email string) bool
	GenerateToken(UserID int) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type authService struct {
	secretKey      string
	issuer         string
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) Login(email string, password string) (string, error) {
	var err error
	user := service.userRepository.FindByEmail(email)
	comparedPassword := comparePassword(user.Password, []byte(password))
	if comparedPassword != true {
		return "", err
	}
	token := service.GenerateToken(int(user.ID))

	return token, err
}

func (service *authService) CreateUser(user param.Register) (model.User, error) {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return userToCreate, err
	}
	res, err := service.userRepository.InsertUser(userToCreate)
	if err != nil {
		return userToCreate, err
	}
	return res, err
}

func (service *authService) FindByEmail(email string) model.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	user := service.userRepository.IsDuplicateEmail(email)
	if user != (model.User{}) {
		return true
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	return secretKey
}

func (j *authService) GenerateToken(UserID int) string {

	j.issuer = "lino"
	j.secretKey = getSecretKey()
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *authService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func extractClaim(tokenString string) string {
	return ""
}
