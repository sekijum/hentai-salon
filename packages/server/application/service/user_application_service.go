package service

import (
	"context"
	"errors"
	"os"
	"server/domain/lib/util"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserApplicationService struct {
	userDatasource *datasource.UserDatasource
}

func NewUserApplicationService(userDatasource *datasource.UserDatasource) *UserApplicationService {
	return &UserApplicationService{userDatasource: userDatasource}
}

type UserApplicationServiceFindByIDParams struct {
	Ctx    context.Context
	UserID int
	Qs     request.UserFindByIdRequest
}

func (svc *UserApplicationService) FindByID(params UserApplicationServiceFindByIDParams) (any, error) {
	user, err := svc.userDatasource.FindByID(datasource.UserDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
		Limit:  params.Qs.Limit,
		Offset: params.Qs.Offset,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewUserResource(resource.NewUserResourceParams{User: user, Limit: params.Qs.Limit, Offset: params.Qs.Offset})

	return dto, nil
}

type UserApplicationServiceSignupParams struct {
	Ctx  context.Context
	Body request.UserSignupRequest
}

func (svc *UserApplicationService) Signup(params UserApplicationServiceSignupParams) (string, error) {
	isDuplicated, err := svc.userDatasource.IsEmailDuplicated(datasource.UserDatasourceIsEmailDuplicatedParams{
		Ctx:   params.Ctx,
		Email: params.Body.Email,
	})
	if err != nil {
		return "", err
	}
	if isDuplicated {
		return "", errors.New("このメールアドレスは既に使用されています。")
	}

	hashedPassword, err := util.HashPassword(params.Body.Password)
	if err != nil {
		return "", err
	}

	user := &model.User{
		EntUser: &ent.User{
			Name:        params.Body.Name,
			Email:       params.Body.Email,
			Password:    hashedPassword,
			AvatarURL:   params.Body.AvatarURL,
			ProfileLink: params.Body.ProfileLink,
			Status:      int(model.UserStatusActive),
			Role:        int(model.UserRoleMember),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	user, err = svc.userDatasource.Create(datasource.UserDatasourceCreateParams{
		Ctx:  params.Ctx,
		User: user,
	})
	if err != nil {
		return "", err
	}

	err = util.ComparePassword(user.EntUser.Password, params.Body.Password)
	if err != nil {
		return "", errors.New("認証情報が無効です")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return "", errors.New("秘密鍵が設定されていません")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.EntUser.ID,
		"exp":    time.Now().AddDate(0, 1, 0).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", errors.New("トークンの生成に失敗しました")
	}

	return tokenString, nil
}

type UserApplicationServiceSigninParams struct {
	Ctx             context.Context
	Email, Password string
}

func (svc *UserApplicationService) Signin(params UserApplicationServiceSigninParams) (string, error) {
	user, err := svc.userDatasource.FindByEmail(datasource.UserDatasourceFindByEmailParams{
		Ctx:   params.Ctx,
		Email: params.Email,
	})
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("ユーザーが登録されていません。")
	}

	err = util.ComparePassword(user.EntUser.Password, params.Password)
	if err != nil {
		return "", errors.New("パスワードが一致しません。")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return "", errors.New("秘密鍵が設定されていません")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.EntUser.ID,
		"exp":    time.Now().AddDate(0, 1, 0).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", errors.New("トークンの生成に失敗しました")
	}

	return tokenString, nil
}

type UserApplicationGetAuthenticatedUserParams struct {
	Ctx         context.Context
	TokenString string
}

func (svc *UserApplicationService) GetAuthenticatedUser(params UserApplicationGetAuthenticatedUserParams) (*resource.UserResource, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("秘密鍵が設定されていません")
	}

	token, err := jwt.Parse(params.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("予期しない署名方法です")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("トークンの解析に失敗しました")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userID"].(float64))
		user, err := svc.userDatasource.FindByID(datasource.UserDatasourceFindByIDParams{
			Ctx:    params.Ctx,
			UserID: userID,
		})
		if err != nil {
			return nil, errors.New("ユーザーの取得に失敗しました")
		}

		dto := resource.NewUserResource(resource.NewUserResourceParams{User: user})

		return dto, nil
	} else {
		return nil, errors.New("トークンが無効です")
	}
}

type UserApplicationServiceUpdateParams struct {
	Ctx    context.Context
	UserID int
	Body   request.UserUpdateRequest
}

func (svc *UserApplicationService) Update(params UserApplicationServiceUpdateParams) error {
	isDuplicated, err := svc.userDatasource.IsEmailDuplicated(datasource.UserDatasourceIsEmailDuplicatedParams{
		Ctx:       params.Ctx,
		Email:     params.Body.Email,
		ExcludeID: &params.UserID,
	})
	if err != nil {
		return err
	}
	if isDuplicated {
		return errors.New("このメールアドレスは既に使用されています。")
	}

	user := model.User{
		EntUser: &ent.User{
			ID:          params.UserID,
			Name:        params.Body.Name,
			Email:       params.Body.Email,
			AvatarURL:   params.Body.AvatarURL,
			ProfileLink: params.Body.ProfileLink,
		},
	}

	_, err = svc.userDatasource.Update(datasource.UserDatasourceUpdateParams{
		Ctx:  params.Ctx,
		User: user,
	})
	if err != nil {
		return err
	}

	return nil
}

type UserApplicationServiceUpdatePasswordParams struct {
	Ctx                      context.Context
	UserID                   int
	OldPassword, NewPassword string
}

func (svc *UserApplicationService) UpdatePassword(params UserApplicationServiceUpdatePasswordParams) error {
	user, err := svc.userDatasource.FindByID(datasource.UserDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
	})
	if err != nil {
		return err
	}

	if err := util.ComparePassword(user.EntUser.Password, params.OldPassword); err != nil {
		return errors.New("古いパスワードが一致しません")
	}

	hashedPassword, err := util.HashPassword(params.NewPassword)
	if err != nil {
		return err
	}

	_, err = svc.userDatasource.UpdatePassword(datasource.UserDatasourceUpdatePasswordParams{
		Ctx:      params.Ctx,
		UserID:   params.UserID,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}
