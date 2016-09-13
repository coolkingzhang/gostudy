package controllers

import (
	"fmt"
	"mishop/app/models"
	"net/http"

	//	l4g "github.com/alecthomas/log4go"
	//	"mishop/app/cache"
	//	"crypto/md5"
	//	"crypto/sha1"
	//	"encoding/hex"
	"mishop/app/controllers/vo"
	"mishop/app/core"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v4"
)

type TestController struct {
}

func (self TestController) Welcome(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func (self TestController) Json(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 320

	core.SessionSet(c, "userId", "This is a session value")
	core.SessionGet(c, "userId")

	//	core.CookiesSet(c)
	//	core.CookiesGet(c, "foo1")

	c.JSON(http.StatusOK, msg)

	//	l4g.Debug("the time is now :%s -- %s", "999", "sad")
}

func (self TestController) Json2(c *gin.Context) {
	core.SessionGet(c, "userId")
	fmt.Print("bee hot test")
}

func (self TestController) Ljson(c *gin.Context) {
	//	var form Login
	//	if c.Bind(&form) == nil {
	//		if form.User == "manu" && form.Password == "123" {
	//			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//		} else {
	//			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		}
	//	}

	var json vo.LoginVo
	if c.BindJSON(&json) == nil {
		if json.User == "manu" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func (self TestController) Fjson(c *gin.Context) {
	var form vo.LoginVo
	if c.Bind(&form) == nil {
		if form.User == "manu" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func (self TestController) Test(c *gin.Context) {
	var user models.User
	user = models.UserModel{}.GetUserById(1)
	fmt.Print(user.Name)
	c.JSON(http.StatusOK, user)
}

func (self TestController) Test2(c *gin.Context) {
	fmt.Print("web test 2")

	var msg struct {
		Name    string `json:"myjson"`
		Message string
		Number  int
	}
	msg.Name = "myname"
	msg.Message = "mymessage"
	msg.Number = 900

	c.JSON(http.StatusOK, msg)
}

func (self TestController) Test3(c *gin.Context) {
	core.LogInfo("web test 3333333333333333")

	var msg struct {
		Name    string `json:"from json33"`
		Message string
		Number  int
	}
	msg.Name = "my 99"
	msg.Message = "my 66"
	msg.Number = 33

	c.JSON(http.StatusOK, msg)

	//	core.Log4go()
	core.LogInfo("log testting")
	core.LogInfo("log4go")

}

func (self TestController) Redis(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	})
	//	pong, err := client.Ping().Result()
	//	fmt.Println(pong, err)

	err := client.Set("key", "my redis value", 0).Err()

	if err != nil {
		fmt.Print(err)
	}
	val, err := client.Get("key").Result()
	fmt.Print(val)
	fmt.Print(err)
	fmt.Print("This is a redis service")
}
