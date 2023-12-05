/**
 * @Author: kwens
 * @Date: 2022/7/13 16:06
 * @Description:
 */

package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type PayLoad map[string]interface{}

func GenToken(payload interface{}, key []byte, opts ...TokenOption) (string, error) {
	var defaultOpt = defaultOption
	for _, opt := range opts {
		opt.apply(&defaultOpt)
	}
	exp := time.Now().Add(defaultOpt.Exp).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     exp,               // 过期时间
		"nbf":     time.Now().Unix(), // 生效时间
		"iat":     time.Now().Unix(), // 签发时间
		"payload": payload,
	})

	return token.SignedString(key)
}

func VerifyToken(token string, key []byte) (interface{}, error) {
	tokenData, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	payload := tokenData.Claims.(jwt.MapClaims)["payload"]
	return payload, nil
}
