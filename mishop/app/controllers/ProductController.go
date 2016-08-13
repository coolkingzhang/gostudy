package controllers

import (
	"fmt"
	"mishop/app/models"
	"net/http"

	//	l4g "github.com/alecthomas/log4go"
	//	"mishop/app/cache"
	//	"crypto/md5"
	"crypto/sha1"
	//	"encoding/hex"
	"mishop/app/controllers/vo"
	"mishop/app/core"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v4"
)

type ProductController struct {
}

func (self ProductController) Welcome(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func (self ProductController) Json(c *gin.Context) {
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

func (self ProductController) Json2(c *gin.Context) {
	core.SessionGet(c, "userId")
	fmt.Print("bee hot test")
}

func (self ProductController) Ljson(c *gin.Context) {
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

func (self ProductController) Fjson(c *gin.Context) {
	var form vo.LoginVo
	if c.Bind(&form) == nil {
		if form.User == "manu" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func (self ProductController) Test(c *gin.Context) {

	var user models.User
	//	user.Name = "gin"
	//	user.Age = 30
	//	user.Salt = "5"
	//	user.Passwd = "123456"
	//	model.UserModel{}.Insert(user)

	user = models.UserModel{}.GetUser(1)
	fmt.Print(user.Name)

	//	h := md5.New()
	//	h.Write([]byte("123456")) // 需要加密的字符串为 123456
	//	cipherStr := h.Sum(nil)
	//	fmt.Println(cipherStr)
	//	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果

	s := "sha1 this string"
	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	c.XML(http.StatusOK, user)
}

func (self ProductController) Redis(c *gin.Context) {
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
