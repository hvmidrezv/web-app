package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName string
	LastName  string
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test endpoint is working.",
	})
	return
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
	return
}

func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
	return
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserByUsername",
		"id":     username,
	})
	return
}

func (h *TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     id,
	})
	return
}
func (h *TestHandler) AddUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
		"id":     id,
	})
	return
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {

	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"id":     userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {

	header := header{}
	err := c.BindHeader(&header)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})
}
func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	})
}
func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	})
}
func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "Uri Binder",
		"ids":    id,
		"name":   name,
	})
}
func (h *TestHandler) BodyBinder(c *gin.Context) {

	p := personData{}
	c.BindJSON(&p)

	c.JSON(http.StatusOK, gin.H{
		"result":            "Uri Binder",
		"person first name": p.FirstName,
		"person last name":  p.LastName,
	})
}
func (h *TestHandler) FormBinder(c *gin.Context) {

	p := personData{}
	c.Bind(&p)

	c.JSON(http.StatusOK, gin.H{
		"result":            "Uri Binder",
		"person first name": p.FirstName,
		"person last name":  p.LastName,
	})
}
