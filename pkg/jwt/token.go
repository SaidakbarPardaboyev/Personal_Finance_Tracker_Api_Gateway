package jwt

import (
	"api_gateway/api/handlers/models"
	pb "api_gateway/genproto/users"

	"api_gateway/configs"
	"fmt"
	"net/smtp"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenarateJWTToken(user *pb.UserByEmail) (*models.Tokens, error) {
	accesToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claimsAccess := accesToken.Claims.(jwt.MapClaims)
	claimsAccess["user_id"] = user.Id
	claimsAccess["email"] = user.Email
	claimsAccess["full_name"] = user.FullName
	claimsAccess["user_role"] = user.UserRole
	claimsAccess["iat"] = time.Now().Unix()
	claimsAccess["exp"] = time.Now().Add(time.Hour).Unix()

	access, err := accesToken.SignedString([]byte(configs.Load().SigningKeyAccess))
	if err != nil {
		return nil, fmt.Errorf("error with generating access token: %s", err)
	}

	claimsRefresh := refreshToken.Claims.(jwt.MapClaims)
	claimsAccess["user_id"] = user.Id
	claimsAccess["email"] = user.Email
	claimsAccess["full_name"] = user.FullName
	claimsAccess["user_role"] = user.UserRole
	claimsRefresh["iat"] = time.Now().Unix()
	claimsRefresh["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := accesToken.SignedString([]byte(configs.Load().SigningKeyRefresh))
	if err != nil {
		return nil, fmt.Errorf("error with generating refresh token: %s", err)
	}

	return &models.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    int(time.Now().Add(time.Hour).Unix()),
	}, nil
}

func GenarateAccessToken(refreshToken string) (
	*models.Tokens, error) {
	accesToken := jwt.New(jwt.SigningMethodHS256)

	claims, err := ExtractClaims(refreshToken, true)
	if err != nil {
		return nil, err
	}
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	accesToken.Claims = claims
	access, err := accesToken.SignedString([]byte(configs.Load().SigningKeyAccess))
	return &models.Tokens{
		AccessToken:  access,
		RefreshToken: refreshToken,
		ExpiresIn:    int(time.Now().Add(time.Hour).Unix()),
	}, err
}

func ValidateToken(tokenStr string, isRefresh bool) (bool, error) {
	_, err := ExtractClaims(tokenStr, isRefresh)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenStr string, isRefresh bool) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(
		t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				t.Header["alg"])
		}
		if isRefresh {
			return []byte(configs.Load().SigningKeyRefresh), nil
		}
		return []byte(configs.Load().SigningKeyAccess), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func SendEmail(to, subject, body string) error {

	var (
		cfg                                *configs.Config
		from, password, smtpHost, smtpPort string
		auth                               smtp.Auth
		msg                                []byte
		err                                error
	)
	cfg = configs.Load()

	// set up authentication information
	from = cfg.Email
	password = cfg.Password
	smtpHost = "smtp.gmail.com"
	smtpPort = ":587"

	auth = smtp.PlainAuth("", from, password, smtpHost)

	// Set up email content
	msg = []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	if err = smtp.SendMail(smtpHost+smtpPort, auth, from, []string{to}, msg); err != nil {
		return fmt.Errorf("error while sending message to email: %v", err)
	}

	return nil
}
