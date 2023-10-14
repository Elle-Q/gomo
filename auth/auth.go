package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"leetroll/common/runtime"
	"leetroll/config"
	"strconv"
	"time"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(ID int) (*Tokens, error) {
	tokens := &Tokens{}
	tokens.AtExpires = time.Now().Add(time.Hour * 24).Unix()
	tokens.AccessUuid = uuid.New().String()

	tokens.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokens.RefreshUuid = uuid.New().String()

	var err error

	//create access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tokens.AccessUuid
	atClaims["user_id"] = ID
	atClaims["exp"] = tokens.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokens.AccessToken, err = at.SignedString([]byte(config.JWTConfig.AccessSecret))
	if err != nil {
		return nil, err
	}

	//create refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokens.AccessUuid
	rtClaims["user_id"] = ID
	rtClaims["exp"] = tokens.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokens.RefreshToken, err = rt.SignedString([]byte(config.JWTConfig.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func CreateAuth(userid uint64, tokens *Tokens) error {
	at := time.Unix(tokens.AtExpires, 0)
	rt := time.Unix(tokens.RtExpires, 0)
	now := time.Now()

	client := runtime.App.GetRedis()
	_, err := client.Set(tokens.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Result()
	if err != nil {
		return err
	}
	_, err = client.Set(tokens.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Result()
	if err != nil {
		return err
	}

	return nil
}

func DeleteAuth(uuid string) (int64, error) {
	client := runtime.App.GetRedis()
	deleted, err := client.Del(uuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
