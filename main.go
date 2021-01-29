package main

import (
	"os"
	"io/ioutil"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/trueConf/pkg/model"
	log "github.com/sirupsen/logrus"
)

const storage = "storage.json"

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no env")
	}
}


func main() {
	port := os.Getenv("PORT")

	e := echo.New()

	e.POST("/user", createUser)
	e.GET("/user", getUserList)
	e.GET("/user/:id", getUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	e.Logger.Fatal(e.Start(":" + port))
}

func createUser(c echo.Context) error{
	storageData, err := ioutil.ReadFile(storage)
	if err != nil {
		log.Info("ioutil error ", err)
	}

	userList := model.UserList{}

	err = json.Unmarshal(storageData, &userList)
	if err != nil {
		log.Info("unmarshal error ", err)
	}

	user := model.User{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Info("POST Request error ", err)
	}

	user.AddId(userList.LastIndex + 1)
	userList.AddLastIndex(user.Id)

	userList.Users = append(userList.Users, user)
	dataOut, err := json.MarshalIndent(&userList,""," ")
	if err != nil {
		log.Info("Marshal Indent error ", err)
	}

	err = ioutil.WriteFile(storage, dataOut, 0)
	if err != nil {
		log.Info("Write File error ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error",})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",})
}



func getUserList(c echo.Context) error{

	storageData, err := ioutil.ReadFile(storage)
	if err != nil {
		log.Info("ioutil error ", err)
	}

	userList := model.UserList{}

	err = json.Unmarshal(storageData, &userList)
	if err != nil {
		log.Info("unmarshal error ", err)
	}


	return c.JSON(http.StatusOK, userList.Users)
}

func getUser(c echo.Context) error{

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Info("convert to int ", err)
	}

	storageData, err := ioutil.ReadFile(storage)
	if err != nil {
		log.Info("ioutil error ", err)
	}

	userList := model.UserList{}
	err = json.Unmarshal(storageData, &userList)
	if err != nil {
		log.Info("unmarshal error ", err)
	}
	index, err := userList.GetUserIndex(intId)
	if err != nil {
		log.Info("no index, ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error","message": "no this id"})
	}
	return c.JSON(http.StatusOK, userList.Users[index])
}


func updateUser(c echo.Context) error{

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Info("convert to int ", err)
	}

	user := model.User{} 
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Info("PUT Request error ", err)
	}

	storageData, err := ioutil.ReadFile(storage)
	if err != nil {
		log.Info("ioutil error ", err)
	}

	userList := model.UserList{}
	err = json.Unmarshal(storageData, &userList)
	if err != nil {
		log.Info("unmarshal error ", err)
	}


	index, err := userList.GetUserIndex(intId)
	if err != nil {
		log.Info("no index, ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error","message": "no this id"})
	}
	userList.Users[index].Name = user.Name
	dataOut, err := json.MarshalIndent(&userList,""," ")
	if err != nil {
		log.Info("Marshal Indent error ", err)
	}
	err = ioutil.WriteFile(storage, dataOut, 0)
	if err != nil {
		log.Info("Write File error ", err)
	}


	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",})
}


func deleteUser(c echo.Context) error{

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Info("convert to int ", err)
	}


	storageData, err := ioutil.ReadFile(storage)
	if err != nil {
		log.Info("ioutil error ", err)
	}
	userList := model.UserList{}

	err = json.Unmarshal(storageData, &userList)
	if err != nil {
		log.Info("unmarshal error ", err)
	}

	index, err := userList.GetUserIndex(intId)
	if err != nil {
		log.Info("no index, ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error","message": "no this id"})
	}

	userList.Users = append(userList.Users[:index], userList.Users[index+1:]...) 

	dataOut, err := json.MarshalIndent(&userList,""," ")
	if err != nil {
		log.Info("Marshal Indent error ", err)
	}

	err = ioutil.WriteFile(storage, dataOut, 0)
	if err != nil {
		log.Info("Write File error ", err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",})
}









