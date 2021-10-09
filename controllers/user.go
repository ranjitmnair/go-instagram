package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"

	"net/http"


	"github.com/ranjitmnair/go-instagram/models"
)

type userController struct{
	session *mgo.Session
}

func newUserController(session *mgo.Session) *userController{
	return &userController{session}
}

func (uc userController) getUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:= p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectHex(id)
	u:=models.User{}

	if err:=uc.session.DB("go-instagram").C("users").FindId(oid).One(&u); err!=nil{
		w.WriteHeader(404)
		return 
	}
	uj, err:=json.Marshal(u)
	if err!=nil{
		fmt.Println(err)
	}
w.Header().Set("Content-Type","application/json")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w,"%s\n",uj)

	
}


func (uc userController) createUser (w http.ResponseWriter, r *http.Request,  httprouter.Params){

	u:=models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id=bson.NewObjectId()
	u.password=bcrypt.GenerateFromPassword(u.password, bcrypt.DefaultCost)
	uc.session.DB("go-instagram").C("users").Insert(u)

	uj, err:= json.Marshal(u)

	if err!=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "%s\n",uj)

}




//get user &create user
