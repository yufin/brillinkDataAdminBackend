package startup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/cmd/migrate"
	common "go-admin/common/models"
	"go-admin/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Ch = make(chan string)

func Startup(configYml string) {
	router := gin.Default()
	err := mime.AddExtensionType(".js", "text/javascript")
	if err != nil {
		return
	}
	err = mime.AddExtensionType(".css", "text/css; charset=UTF-8")
	if err != nil {
		return
	}
	err = mime.AddExtensionType(".html", "text/html; charset=UTF-8")
	if err != nil {
		return
	}
	router.Static("/startup", "./static/guide")
	router.NoRoute(func(c *gin.Context) {
		if c.Request.URL.Path == "/startup/home" || c.Request.URL.Path == "/startup/config" {
			accept := c.Request.Header.Get("Accept")
			flag := strings.Contains(accept, "text/html")
			if flag {
				content, err := ioutil.ReadFile("./static/guide/index.html")
				if (err) != nil {
					c.Writer.WriteHeader(404)
					c.Writer.WriteString("Not Found")
					return
				}
				c.Writer.WriteHeader(200)
				c.Writer.Header().Add("Accept", "text/html")
				_, err = c.Writer.Write(content)
				if err != nil {
					return
				}
				c.Writer.Flush()
			}
		}
	})
	router.POST("/api/v1/config", SetupConfig)
	router.POST("/api/v1/migrate", DatabaseMigrate)
	srvGuide := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srvGuide.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	err = OpenBrowser("http://localhost:8080/startup")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case msg := <-Ch:
			fmt.Println("已经完成:" + msg)
			if msg == "done2" {
				return
			} else if msg == "done2errors" {
				panic("数据库迁移失败")
			}
		}
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 不同平台启动指令不同
var commands = map[string]string{
	"windows": "explorer",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenBrowser(uri string) error {
	// runtime.GOOS获取当前平台
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Run()
}

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

func SetupConfig(c *gin.Context) {
	var config Config
	err := c.BindJSON(&config)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "errMessage": "配置文件生成失败" + err.Error()})
		return
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == "" {
		config.Port = "3306"
	}
	if config.User == "" {
		config.User = "root"
	}
	if config.Password == "" {
		config.Password = "123456"
	}
	if config.DatabaseName == "" {
		config.DatabaseName = "go-admin"
	}
	path := "./config/"
	fileName := "settings.yml"
	// 判断文件是否存在，如果存在就备份
	if pathExists(path + fileName) {
		_ = os.Rename(path+"settings.yml", path+"settings.yml.bak")
	}
	err = template.GenCodeForString(config, template.Yml, path, fileName, fileName)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "errMessage": "配置文件生成失败" + err.Error()})
	}
	c.JSON(200, gin.H{"success": true})
	Ch <- "done1"
}

func DatabaseMigrate(c *gin.Context) {
	var systemInfo common.SystemInfo
	err := c.BindJSON(&systemInfo)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "errMessage": "配置文件生成失败" + err.Error()})
		return
	}
	systemInfo.Default(systemInfo.SystemName, systemInfo.Username, systemInfo.Password)

	err = migrate.Run()
	if err != nil {
		c.JSON(200, gin.H{"success": false, "errMessage": "数据库迁移失败" + err.Error()})
		Ch <- "done2error"
	} else {
		c.JSON(200, gin.H{"success": true})
		Ch <- "done2"
	}
	// 安装成功后删除初始化标记文件
	err = os.Remove("./config/startup.txt")
	if err != nil {
		return
	}
}
