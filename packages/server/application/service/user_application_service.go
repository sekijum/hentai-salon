package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"server/domain/lib/util"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
	"time"

	mailpit "server/infrastructure/mailpit"

	"github.com/dgrijalva/jwt-go"
)

type UserApplicationService struct {
	userDatasource *datasource.UserDatasource
	mailpitClient  *mailpit.MailpitClient
}

func NewUserApplicationService(userDatasource *datasource.UserDatasource, mailpitClient *mailpit.MailpitClient) *UserApplicationService {
	return &UserApplicationService{userDatasource: userDatasource, mailpitClient: mailpitClient}
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
	Ctx    context.Context
	UserID int
	Body   request.UserUpdatePasswordRequest
}

func (svc *UserApplicationService) UpdatePassword(params UserApplicationServiceUpdatePasswordParams) error {
	user, err := svc.userDatasource.FindByID(datasource.UserDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
	})
	if err != nil {
		return err
	}

	if err := util.ComparePassword(user.EntUser.Password, params.Body.OldPassword); err != nil {
		return errors.New("現在のパスワードが一致しません")
	}

	hashedPassword, err := util.HashPassword(params.Body.NewPassword)
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

type UserApplicationServiceForgotPasswordParams struct {
	Ctx  context.Context
	Body request.UserForgotPasswordRequest
}

func (svc *UserApplicationService) ForgotPassword(params UserApplicationServiceForgotPasswordParams) error {
	user, err := svc.userDatasource.FindByEmail(datasource.UserDatasourceFindByEmailParams{
		Ctx:   params.Ctx,
		Email: params.Body.Email,
	})
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("ユーザーが見つかりません。")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return errors.New("秘密鍵が設定されていません")
	}

	resetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.EntUser.ID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // 有効期限24時間
	})

	tokenString, err := resetToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return fmt.Errorf("トークンの生成に失敗しました: %v", err)
	}

	clientURL := os.Getenv("CLIENT_URL")
	resetURL := fmt.Sprintf("%s/reset-password?token=%s", clientURL, url.QueryEscape(tokenString))

	emailSubject := "パスワードリセットのリクエスト"
	emailBody := fmt.Sprintf(`
%sさん

変態サロンをご利用いただきありがとうございます。
パスワードの再設定は以下のリンクからお願いします。
このリンクの有効期限は24時間です。

パスワードの再設定
%s

Webページを開く
%s

※このメールは返信しても届きません。`, user.EntUser.Name, resetURL, clientURL)

	err = svc.mailpitClient.SendEmail(params.Body.Email, emailSubject, emailBody)
	if err != nil {
		return fmt.Errorf("メールの送信に失敗しました: %v", err)
	}

	return nil
}

type TokenClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

type UserVerifyResetPasswordTokenRequestParams struct {
	Ctx  context.Context
	Body request.UserVerifyResetPasswordTokenRequest
}

func (svc *UserApplicationService) VerifyResetPasswordToken(params UserVerifyResetPasswordTokenRequestParams) error {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return errors.New("秘密鍵が設定されていません")
	}

	token, err := jwt.ParseWithClaims(params.Body.Token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("予期しない署名方法です")
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return errors.New("トークンの解析に失敗しました")
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return errors.New("トークンの有効期限が切れています")
		}
		return nil
	} else {
		return errors.New("トークンが無効です")
	}
}

type UserResetPasswordRequestParams struct {
	Ctx  context.Context
	Body request.UserResetPasswordRequest
}

func (svc *UserApplicationService) ResetPassword(params UserResetPasswordRequestParams) error {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return errors.New("秘密鍵が設定されていません")
	}

	token, err := jwt.ParseWithClaims(params.Body.Token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("予期しない署名方法です")
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return errors.New("トークンの解析に失敗しました")
	}

	var userID int
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return errors.New("トークンの有効期限が切れています")
		}
		userID = claims.UserID
	} else {
		return errors.New("トークンが無効です")
	}

	hashedPassword, err := util.HashPassword(params.Body.Password)
	if err != nil {
		return fmt.Errorf("パスワードのハッシュ化に失敗しました: %v", err)
	}

	_, err = svc.userDatasource.UpdatePassword(datasource.UserDatasourceUpdatePasswordParams{
		Ctx:      params.Ctx,
		UserID:   userID,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}
