package service

import (
	"context"
	"errors"
	"os"
	"server/domain/model"
	"server/infrastructure/datasource"
	request "server/presentation/request"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Signup(ctx context.Context, body request.UserSignupRequest) (string, error)
	Signin(ctx context.Context, email, password string) (string, error)
	GetAuthenticatedUser(ctx context.Context, tokenString string) (*model.User, error)
}

type userService struct {
	userDatasource datasource.UserDatasource
}

func NewUserService(userDatasource datasource.UserDatasource) UserService {
	return &userService{userDatasource: userDatasource}
}

func (svc *userService) Signup(ctx context.Context, body request.UserSignupRequest) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &model.User{
		Name:        body.Name,
		Email:       body.Email,
		Password:    string(hashedPassword),
		DisplayName: body.DisplayName,
		AvatarUrl:   body.AvatarUrl,
		Status:      0, // Active
		Role:        0, // Member
	}

	_, err = svc.userDatasource.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return svc.Signin(ctx, body.Email, body.Password)
}

func (svc *userService) Signin(ctx context.Context, email, password string) (string, error) {
	user, err := svc.userDatasource.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("認証情報が無効です")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("秘密鍵が設定されていません")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("トークンの生成に失敗しました")
	}

	return tokenString, nil
}

func (svc *userService) GetAuthenticatedUser(ctx context.Context, tokenString string) (*model.User, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("秘密鍵が設定されていません")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("予期しない署名方法です")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("トークンの解析に失敗しました")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["user_id"].(float64))
		user, err := svc.userDatasource.GetUserByID(ctx, userID)
		if err != nil {
			return nil, errors.New("ユーザーの取得に失敗しました")
		}
		return user, nil
	} else {
		return nil, errors.New("トークンが無効です")
	}
}
