package myjwt

import (
	"csrf-project/db/models"

	"github.com/golang-jwt/jwt"
)



const (
	privKeyPath = "keys/app.rsa"
    pubKeyPath = "keys/app.rsa.pub"
)


func InitJWT()error{
   signBytes, err := ioutil.ReadFile(privKeyPath);

   if err != nil {
	return err;
   }

   signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes);
   if err!= nil {
	return err;
   }

   verifyBytes, err := ioutil.ReadFile(pubKeyPath);
   if err!= nil {
	return err;
   }

   verifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(verifyBytes);

   if err!= nil {
	return err;
   }
   return nil

}

func CreateNewTokens(uuid string, role string)(authTokenString,refreshToken string){

   // generate the csrf secret!
    csrfSecret, err := models.GenerateCSRFSecret();
    if err!= nil {
      return;
    }

    // generating the refresh token!
    refreshToken, err = StringcreateRefreshTokenString(uuid,role,csrfSecret);

    if err!= nil {
      return;
    }

    // generating the auth token!
    authTokenString, err = createAuthTokenString(uuid, role, csrfSecret);

    if err!= nil {
      return;
    }

    return;



}

func CheckAndRefreshTokens(oldAuthTokenString string, oldRefreshTokenString string, oldCsrfSecret string) (newAuthTokenString, newRefreshTokenString, newCsrfSecret string, err error) {

	if oldCsrfSecret == "" {
		log.Println("No CSRF token!")
		err = errors.New("Unauthorized")
		return
	}

	authToken, err := jwt.ParseWithClaims(oldAuthTokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	authTokenClaims, ok := authToken.Claims.(*models.TokenClaims)
	if !ok {
		return
	}
	if oldCsrfSecret != authTokenClaims.Csrf {
		log.Println("CSRF token doesn't match jwt!")
		err = errors.New("Unauthorized")
		return
	}

	if authToken.Valid {
		log.Println("Auth token is valid")

		newCsrfSecret = authTokenClaims.Csrf

		newRefreshTokenString, err = updateRefreshTokenExp(oldRefreshTokenString)
		newAuthTokenString = oldAuthTokenString
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		log.Println("Auth token is not valid")
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			log.Println("Auth token is expired")

			newAuthTokenString, newCsrfSecret, err = updateAuthTokenString(oldRefreshTokenString, oldAuthTokenString)
			if err != nil {
				return
			}

			newRefreshTokenString, err = updateRefreshTokenExp(oldRefreshTokenString)
			if err != nil {
				return
			}

			newRefreshTokenString, err = updateRefreshTokenCsrf(newRefreshTokenString, newCsrfSecret)
			return
		} else {
			log.Println("Error in auth token")
			err = errors.New("Error in auth token")
			return
		}
	} else {
		log.Println("Error in auth token")
		err = errors.New("Error in auth token")
		return
	}

	err = errors.New("Unauthorized")
	return
}

func createAuthTokenString(uuid string, role string, csrfSecret string) (authTokenString string, err error) {
	authTokenExp := time.Now().Add(models.AuthTokenValidTime).Unix()
	authClaims := models.TokenClaims{
		jwt.StandardClaims{
			Subject:   uuid,
			ExpiresAt: authTokenExp,
		},
		role,
		csrfSecret,
	}

	authJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), authClaims)

	authTokenString, err = authJwt.SignedString(signKey)
	return
}



func createRefreshTokenString()(){

}

func updateRefreshTokenExp()(){

}

func updateAuthTokenString() error{

}


func RevokeRefreshToken()error{

}


func updateRefreshTokenCsrf()(){

}

func GrabUUID()(){

}