package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"
	"gorm.io/gorm"
)

type PostElement struct {
	Author  string
	Title   string
	Content string
}
type Credentials struct {
	AdminLogin    string
	AdminPassword string
	DBLogin       string
	DBPassword    string
}

var credentials Credentials
var IsAdmin bool
var DB *gorm.DB
var Articles []Post

func initCredentials() {
	file, _ := os.Open("assets/admin_credentials.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "admin_login") {
			credentials.AdminLogin = strings.Split(line, "=")[1]
		} else if strings.Contains(line, "admin_password") {
			credentials.AdminPassword = strings.Split(line, "=")[1]
		} else if strings.Contains(line, "db_login") {
			credentials.DBLogin = strings.Split(line, "=")[1]
		} else if strings.Contains(line, "db_password") {
			credentials.DBPassword = strings.Split(line, "=")[1]
		}
	}
}

func initialization() (app *fiber.App, db *gorm.DB, err error) {
	engine := html.New("./views", ".html")
	app = fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layout",
	})
	app.Static("/", "./assets")

	// app.Use(limiter.New())

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 5 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
	}))

	initCredentials()
	err = createConnection()
	if err != nil {
		fmt.Println("db error:", err)
	}

	if err == nil {
		app.Get("/articles/", func(c *fiber.Ctx) error {
			fetchPosts()
			data_db := struct {
				Title string
				Items []Post
			}{
				Title: "Article redactor",
				Items: Articles,
			}
			return c.Render("index", data_db)
		})
		app.Get("/admin", func(c *fiber.Ctx) error {
			data_admin := struct {
				IsLoggedIn bool
			}{
				IsLoggedIn: IsAdmin,
			}
			return c.Render("login", data_admin)
		})
		app.Post("/admin", AdminEndpoint)
		app.Get("/show/:article_id", func(c *fiber.Ctx) error {
			var p Post
			DB.Where("id = ?", c.Params("article_id")).First(&p)
			c.Set("Content-type", "text/html; charset=UTF-8")
			return c.SendString(p.Content)
			// return c.Render("article", p)
		})
	}
	return
}

func main() {

	app, _, err := initialization()
	if err == nil {
		log.Fatal(app.Listen(":8888"))
	}

}

// var page int
// if c.Params("page") == "" {
// 	page = 0
// } else {
// 	page, _ = strconv.Atoi(c.Params("page"))
// }
// if page == 0 {
// data_db := struct {
// 	Title string
// 	Items []Post
// }{
// 	Title: "Article redactor",
// 	Items: Articles[:3],
// }
// return c.Render("index", data_db)
// } else {
// 	data_db := struct {
// 		Title string
// 		Items []Post
// 	}{
// 		Title: "Article redactor",
// 		Items: Articles[3*page : 3*(page+1)],
// 	}
// 	return c.Render("index", data_db)
// }
