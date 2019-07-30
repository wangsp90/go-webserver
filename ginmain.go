package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	path = "config"
)

type httpcfg struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Path     string `json:"path"`
}

type UserInfo struct {
	User   string `form:"user"`
	Passwd string `form:"passwd"`
}

func main() {
	var (
		tmpb []byte
		err  error
		cfg  httpcfg
	)

	if tmpb, err = ioutil.ReadFile(path); err != nil {
		log.Println(err)
		return
	}

	if err = json.Unmarshal(tmpb, &cfg); err != nil {
		log.Println(err)
		return
	}

	// r1 := gin.Default()
	// r1.StaticFS("/", http.Dir(cfg.Path))
	// go r1.Run(":8789")

	// r2 := gin.Default()
	// r2.GET("/test", index)
	// r2.LoadHTMLGlob("template/*.html")
	// r2.Run(cfg.Port)
	r := gin.Default()
	r.StaticFS("/share", http.Dir(cfg.Path))
	r.LoadHTMLGlob("template/*.html")
	r.GET("/login", Login)
	r.Run(cfg.Port)
}

func Login(c *gin.Context) {
	var (
		u   UserInfo
		err error
	)

	if err = c.Bind(&u); err != nil {
		log.Println(err)
		time.Sleep(time.Second * 200)
		return
	}
	log.Println(u)
	if u.User != "Wangsp" || u.Passwd != "123456" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	} else {
		c.HTML(http.StatusOK, "upload.html", nil)
	}

	// name := c.PostForm("name")
	// log.Println(name)
	// file, header, err := c.Request.FormFile("upload")
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
	// }
	// filename := header.Filename

	// log.Println(file, err, filename)

	// out, err := os.Create(filename)
	// defer out.Close()
	// io.Copy(out, file)
	// c.String(http.StatusCreated, "upload successful")
}
