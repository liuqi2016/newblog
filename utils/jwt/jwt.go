package jwt

import (
	"blog/utils/setting"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetToken 获取token
func GetToken(username, password string) (token string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})

	// Sign and get the complete encoded token as a string using the secret
	token, err = tokenObj.SignedString(jwtSecret)

	// fmt.Println(tokenString, err)
	return
}

// ParseToken 解析token
func ParseToken(tokenString string) (userInfo UserInfo, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	tokenObj, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		fmt.Println(claims["username"], claims["password"])
		userInfo.Username = claims["username"].(string)
		userInfo.Password = claims["password"].(string)
	}
	return
}
