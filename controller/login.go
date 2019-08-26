package controller

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-expenseapp-api/model"
	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// LoginRequest Structure of the API request parameters
type LoginRequest struct {
	// Username Username of the user
	Username string `json:"username"`
	// Password Password of the user
	Password string `json:"password"`
}

// Login Controller for login
func Login(c *gin.Context) {
	var request LoginRequest
	err := c.BindJSON(&request)

	if err != nil{
		c.JSON(400, model.Response(model.BadRequestError, "Invalid Input", nil))
		return
	}

	auth, err := model.IsAuthenticatedUser(request.Username, request.Password)
	
	if err != nil{
		c.JSON(500, model.Response(model.InternalServerError, err.Error(), nil))
		return
	}

	if !auth {
		c.JSON(200, model.Response(model.Failed, "Invalid Username or Password!", nil))
		return
	}
	
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(500, model.Response(model.InternalServerError, err.Error(), nil))
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	data := map[string]string{"token" : tokenString}
	c.JSON(200, model.Response(model.Success, "Successfully logged in!", data))
}