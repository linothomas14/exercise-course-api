package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email string, password string, role string) (string, error)

	IsDuplicateEmail(email string) bool
	GenerateToken(UserID int, role string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type authService struct {
	secretKey       string
	issuer          string
	userRepository  repository.UserRepository
	adminRepository repository.AdminRepository
}

type actor struct {
	id       uint32
	email    string
	password string
}

func NewAuthService(userRep repository.UserRepository, adminRep repository.AdminRepository) AuthService {
	return &authService{
		userRepository:  userRep,
		adminRepository: adminRep,
	}
}

func (service *authService) Login(email string, password string, role string) (string, error) {
	var err error

	var actorStruct actor

	if role == "admin" {
		admin := service.adminRepository.FindByEmail(email)
		actorStruct = parseActor(admin.ID, admin.Email, admin.Password)

	} else {
		user := service.userRepository.FindByEmail(email)
		actorStruct = parseActor(user.ID, user.Email, user.Password)

	}

	comparedPassword := comparePassword(actorStruct.password, []byte(password))

	if comparedPassword != true {
		return "", err
	}
	token := service.GenerateToken(int(actorStruct.id), role)

	return token, err
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

func (j *authService) GenerateToken(UserID int, role string) string {

	j.issuer = "lino"
	j.secretKey = getSecretKey()
	claims := &jwt.MapClaims{
		"user_id":   UserID,
		"role":      role,
		"ExpiresAt": time.Now().AddDate(1, 0, 0).Unix(),
		"Issuer":    j.issuer,
		"IssuedAt":  time.Now().Unix(),
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

func parseActor(id uint32, email string, password string) actor {
	var actorStruct actor
	actorStruct.id = id
	actorStruct.email = email
	actorStruct.password = password

	return actorStruct
}
