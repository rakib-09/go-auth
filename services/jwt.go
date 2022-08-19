package services

import (
	"github.com/golang-jwt/jwt"
	"go-auth/config"
	"go-auth/const"
	"go-auth/types"
	"time"
)

type JwtTokenService struct {
}

func NewJwtSvc() *JwtTokenService {
	return &JwtTokenService{}
}

func (jt *JwtTokenService) CreateToken(user *types.UserResp) (*types.JwtToken, error) {
	jwtConf := config.Jwt()
	token := &types.JwtToken{}
	var err error

	token.UserID = user.ID
	token.AccessExpiry = time.Now().Add(time.Minute * jwtConf.AccessTokenExpiry).Unix()
	token.RefreshExpiry = time.Now().Add(time.Minute * jwtConf.RefreshTokenExpiry).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["uid"] = user.ID
	atClaims["name"] = user.Name
	atClaims["exp"] = token.AccessExpiry

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(jwtConf.AccessTokenSecret))
	if err != nil {
		return nil, _const.InvalidAccessToken
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["uid"] = user.ID
	rtClaims["name"] = user.Name
	rtClaims["exp"] = token.RefreshExpiry

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(jwtConf.RefreshTokenSecret))
	if err != nil {
		return nil, _const.InvalidAccessToken
	}

	return token, nil
}
