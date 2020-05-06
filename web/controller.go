package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var router *gin.Engine

var ch chan string

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	ch = make(chan string)
	go queueData("a")
	go queueData("b")
	go queueData("c")
	router := gin.Default()
	router.GET("/someGet:name/*action", getting)
	router.GET("/getParam", gettingParam)
	router.GET("/redirect", redirect)
	router.GET("/longAsync", longAsync)
	router.POST("/somePost", posting)
	router.POST("/upload", uploading)
	router.Any("/anyPerson", startPage)
	router.Any("/go", getGoroutineNum)
	router.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if c.BindQuery(&person) == nil {
		log.Println("====== Only Bind Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
}

// http://localhost:8080/longAsync get
func longAsync(c *gin.Context) {
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

// // http://localhost:8080/redirect get
func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
}

func uploading(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	//c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// http://localhost:8080/somePost post
// nick:huaxu
// message:222222
func posting(c *gin.Context) {
	id := c.Query("id")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	ua := c.GetHeader("User-Agent")
	ct := c.GetHeader("Content-Type")

	fmt.Printf("posting get header User-Agent=%s, Content-Type=%s", ua, ct)

	c.Header("User-Agent","Mozilla/5.0")
	c.Header("Content-Type","text/html; charset=utf-8")

	c.JSON(200, gin.H{
		"status"	:  "posted",
		"message"	: 	message,
		"nick"		:   nick,
		"id"		: 	id,
	})
}

// http://localhost:8080/getParam?firstname=huaxu&lastname=rookie
func gettingParam(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "Guest")
	lastName := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
}

func getGoroutineNum(c *gin.Context) {
	c.String(http.StatusOK, strconv.Itoa(runtime.NumGoroutine()))
}



// http://localhost:8080/someGethuaxu/23232
func getting(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	go func() {
		ch <- name
	}()
	c.String(http.StatusOK, message)
}

func queueData(chName string) {
	for {
		time.Sleep(1*time.Second)
		name := <- ch
		fmt.Println(chName + ":" + name)
		if len(name) == 10 {
			fmt.Println(chName + ": ok~")
		} else {
			go func() {
				ch <- name + "a"
			}()
		}
	}
}


