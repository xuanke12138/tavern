package login

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey string
	//SigningKey *ecdsa.PrivateKey
	//SignOutKey *ecdsa.PublicKey
}

var(
	key []byte = []byte("Hello World！This is secret!")
)

type Claim struct{
	Name string
	Password string
	jwt.StandardClaims
}


func CreatToken(name string,password string) (string,error){

	nowTime:=time.Now()
	expireTime:=nowTime.Add(60*time.Hour)


	// 产生json web token

	claims:=Claim{
		Name:     name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "wym",
		},
	}


	Claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Claims.SignedString(key)
	return token,err

	//？？？？？？
	//神奇的错误用上边三行就可以成功，用下面三行就会报错
	//es256是非对称加密（有格式要求），hs256是对称加密
	//Claims := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	//token , err := Claims.SignedString(key)
	//return token,err
}



// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string)(*Claim,string){

	var claims Claim

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})

	if err != nil {
		return &Claim{},"token超时"
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*Claim); ok && setToken.Valid {
			return key, "nil"
		}
	}
	return nil, "没有token"




	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	//tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
	//	return key, nil
	//})
	//
	//if err != nil{
	//	return nil,err
	//}
	//if tokenClaims == nil {
	//	return nil,errors.New("Token 无效")
	//}
	//
	//
	//if tokenClaims!=nil{
	//	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
	//	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	//	if claims,ok:=tokenClaims.Claims.(*Claim);ok&&tokenClaims.Valid{
	//		return claims,nil
	//	}
	//}
	//
	//return nil,err

}

