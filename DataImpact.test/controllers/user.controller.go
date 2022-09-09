package controllers

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	
	"dataimpact/test/golang/models"
	"dataimpact/test/golang/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/tidwall/gjson"
	//"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	UserService services.UserService
}


func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}


func HashPassword(password string) string{
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!=nil{
		log.Panic(err)
	}
	return string(hash)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var users []interface{}
	var user models.User
	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&users)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	password := HashPassword(user.Password)
	user.Password = password
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.UserService.GetUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}



func VerifyPassword(userPassword string, providedPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = fmt.Sprintf("id or password is incorrect")
		println(userPassword)
		println(providedPassword)
		println([]byte(providedPassword))
		println([]byte(userPassword))
		check=false
	}
	return check, msg
}




func (uc *UserController) Login(w http.ResponseWriter, r *http.Request){
	var user models.User
	var foundUser models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}
	

	
	data, err := json.Marshal(user)
	jsondata := data
	id := gjson.Get(string(jsondata), "id")
	password := gjson.Get(string(jsondata), "password")
	if err != nil{
		panic(err)
	}
	

	err1, err2 := uc.UserService.LoginUser(id.String(), password.String())
	if err1 != nil || err2 != nil {
		w.Write([]byte("id or password is incorrect"))
		return 
	}
	users, err := uc.UserService.GetAll()
	all, err := json.Marshal(users)
	w.Write([]byte(all))
	//ctx.JSON(http.StatusOK, users)


	passwordIsValid, msg := VerifyPassword(foundUser.Password, password.String())
	if passwordIsValid != true{
		w.Write([]byte(msg))
		return 
	}


	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
}






func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.UserService.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/add/users", uc.CreateUser)
	rg.GET("/user/:id", uc.GetUser)
	rg.GET("/users/list", uc.GetAll)
	rg.PATCH("/user/:id", uc.UpdateUser)
	rg.DELETE("/delete/user/:id", uc.DeleteUser)

	/*mux := http.NewServeMux()
	mux.HandleFunc("/Login", uc.Login)
	http.ListenAndServe(":5000", mux)*/
}
