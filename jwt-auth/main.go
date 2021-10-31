package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/novalagung/gubrak/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type M map[string]interface{}

var APPLICATION_NAME = "My Simple Jwt App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	Group    string `json:"group"`
}

func main() {
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareJWTAuthorization)

	mux.HandleFunc("/login", HandleLogin)
	mux.HandleFunc("/index", HandleIndex)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":8080"

	log.Println("starting server at", server.Addr)

	server.ListenAndServe()
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/login" {
			next.ServeHTTP(writer, request)
			return
		}

		authorizationHeader := request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(writer, "Invalid token", http.StatusBadRequest)
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}

func HandleIndex(rw http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	message := fmt.Sprintf("Hello %s (%s", userInfo["username"], userInfo["group"])
	log.Fatal(rw.Write([]byte(message)))
}
func HandleLogin(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "Unsupported http method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(rw, "Invalid username or password", http.StatusBadRequest)
		return
	}

	ok, userInfo := authenticateUser(username, password)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
		Group:    userInfo["group"].(string),
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(M{"token": signedToken})
	rw.Write(tokenString)
}

func authenticateUser(username, password string) (bool, M) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "users.json")
	buf, _ := ioutil.ReadFile(dbPath)

	data := make([]M, 0)
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return false, nil
	}

	res := gubrak.From(data).Find(func(each M) bool {
		return each["username"] == username && each["password"] == password
	}).Result()

	if res != nil {
		resM := res.(M)
		delete(resM, "password")
		return true, resM
	}

	return false, nil
}
