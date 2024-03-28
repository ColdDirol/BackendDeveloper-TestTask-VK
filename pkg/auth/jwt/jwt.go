package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var SecretKey []byte
var Salt string

type JWT struct {
	Token string `json:"token"`
}

type Claims struct {
	Login     string `json:"login"`
	ExpiresAt int64  `json:"exp"`
}

func hmacSha256(data string, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Sha256EncodeWithSalt(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(Salt))
	hasher.Write([]byte(input))
	encodedBytes := hasher.Sum(nil)
	encodedString := hex.EncodeToString(encodedBytes)
	return encodedString
}

func CreateToken(login string) (*JWT, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // Срок действия токена: 5 минут

	claims := &Claims{
		Login:     login,
		ExpiresAt: expirationTime.Unix(),
	}

	claimsJson, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	header := base64.StdEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := base64.StdEncoding.EncodeToString(claimsJson)

	token := fmt.Sprintf("%s.%s", header, payload)
	signature := hmacSha256(token, SecretKey)

	return &JWT{
		Token: fmt.Sprintf("%s.%s", token, signature),
	}, nil
}

func VerifyToken(token string) (bool, *Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, nil, fmt.Errorf("invalid token")
	}

	payload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, nil, err
	}

	var claims Claims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return false, nil, err
	}

	signature := hmacSha256(parts[0]+"."+parts[1], SecretKey)
	if signature != parts[2] {
		return false, nil, fmt.Errorf("invalid signature")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return false, nil, fmt.Errorf("token has expired")
	}

	return true, &claims, nil
}

//func RefreshToken(token string) (*JWT, error) {
//	claims, err := VerifyToken(token)
//	if err != nil {
//		return nil, err
//	}
//	newToken, err := CreateToken(claims.Login)
//	if err != nil {
//		return nil, err
//	}
//	return newToken, nil
//}
