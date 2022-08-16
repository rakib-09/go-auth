package services

import (
	"github.com/golang-jwt/jwt"
	"go-auth/config"
	"go-auth/errors"
	"go-auth/types"
	"time"
)

type JwtTokenService struct {
}

func NewJwtToken() *JwtTokenService {
	return &JwtTokenService{}
}

func (jt *JwtTokenService) CreateToken(userId uint) (*types.JwtToken, error) {
	jwtConf := config.Jwt()
	token := &types.JwtToken{}
	var err error

	token.UserID = userId
	token.AccessExpiry = time.Now().Add(time.Minute * jwtConf.AccessTokenExpiry).Unix()
	token.RefreshExpiry = time.Now().Add(time.Minute * jwtConf.RefreshTokenExpiry).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["uid"] = userId
	atClaims["exp"] = token.AccessExpiry

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(jwtConf.AccessTokenSecret))
	if err != nil {
		return nil, errors.InvalidAccessToken
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["uid"] = userId
	rtClaims["exp"] = token.RefreshExpiry

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(jwtConf.RefreshTokenSecret))
	if err != nil {
		return nil, errors.InvalidAccessToken
	}

	return token, nil
}
