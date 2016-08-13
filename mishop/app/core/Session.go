package core

import (
	"log"

	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	//	gsr "gopkg.in/boj/redistore.v1"
)

//session存储memcache
var signature []byte = []byte("tcl-kuyu2016")
var sessionName string = "gosession"
var sessionPrefix = "gosession_"
var sessionPath string = "/"
var sessionDomain string = "localhost"
var sessionMaxAge int = 1800

var StoreSession = gsm.NewMemcacheStore(MemcacheClient, sessionPrefix, signature)

//session存储redis
//var StoreSession, err = gsr.NewRediStore(50, "tcp", "10.73.128.85:6379", "", signature)

// cookies存储
var StoreCookie = sessions.NewCookieStore(signature)

// session
func SessionGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		session, err := StoreSession.Get(c.Request, sessionName)
		if err != nil {
			log.Print("Error getting session: %v", err)
		}
		SessionKeep(c)
		log.Print("session middle ware:")
		log.Print(session.ID)
	}
}

func SessionKeep(c *gin.Context) {
	session, _ := StoreSession.Get(c.Request, sessionName)
	session.Options.Path = sessionPath
	session.Options.Domain = sessionDomain
	session.Options.MaxAge = sessionMaxAge

	session.Save(c.Request, c.Writer)
}

func SessionGet(c *gin.Context, key string) {
	session, _ := StoreSession.Get(c.Request, sessionName)
	log.Print(session.Values[key])
	log.Print("get session ID: " + session.ID)
}

func SessionSet(c *gin.Context, key string, value string) {
	session, _ := StoreSession.Get(c.Request, sessionName)
	session.Options.Path = sessionPath
	session.Options.Domain = sessionDomain
	session.Options.MaxAge = sessionMaxAge

	session.Values[key] = value
	session.Save(c.Request, c.Writer)
	log.Print("set session ID: " + session.ID)
}

// cookies

func CookiesGet(c *gin.Context, key string) {
	cookie, _ := StoreCookie.Get(c.Request, sessionName)
	log.Print(cookie.Values[key])
	log.Print("get cookies ID: " + cookie.ID)
}

func CookiesSet(c *gin.Context, key string, value string) {
	cookie, _ := StoreCookie.Get(c.Request, sessionName)
	cookie.Options.Path = sessionPath
	cookie.Options.Domain = sessionDomain
	cookie.Options.MaxAge = sessionMaxAge

	cookie.Values[key] = value
	cookie.Save(c.Request, c.Writer)
	log.Print("set cookies ID: " + cookie.ID)
}
