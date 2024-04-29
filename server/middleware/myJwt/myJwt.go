package myjwt;

import "github.com/golang-jwt/jwt"



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

func CreateNewTokens()(){

}

func CheckAndRefreshTokens()(){

}

func createAuthTokenString()(){

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