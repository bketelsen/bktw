package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS
var articles []Article
var homepage Homepage
var global Global
var categories []Category
var updateMutex sync.Mutex
var strapi string
var port string

type HomeContext struct {
	Articles   []Article
	Homepage   Homepage
	Global     Global
	Categories []Category
}

type PostContext struct {
	Article    Article
	Homepage   Homepage
	Global     Global
	Categories []Category
}
type CategoryContext struct {
	Homepage Homepage
	Global   Global
	Category Category
}

func BySlug(slug string) (Article, error) {
	for _, article := range articles {
		fmt.Println(slug, article.Slug)
		if article.Slug == slug {
			return article, nil
		}
	}
	return Article{}, errors.New("Article not found")
}

func CategoryBySlug(slug string) (Category, error) {
	for _, cat := range categories {
		fmt.Println(slug, cat.Slug)
		if cat.Slug == slug {
			return cat, nil
		}
	}
	return Category{}, errors.New("Category not found")
}
func getArticles() {
	updateMutex.Lock()
	defer updateMutex.Unlock()
	resp, err := http.Get(strapi + "/articles")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := json.Unmarshal(body, &articles); err != nil {
		panic(err)
	}
}
func getCategories() {
	updateMutex.Lock()
	defer updateMutex.Unlock()
	resp, err := http.Get(strapi + "/categories")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := json.Unmarshal(body, &categories); err != nil {
		panic(err)
	}
}

func getGlobal() {
	updateMutex.Lock()
	defer updateMutex.Unlock()
	resp, err := http.Get(strapi + "/global")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := json.Unmarshal(body, &global); err != nil {
		panic(err)
	}
}
func getHomepage() {
	updateMutex.Lock()
	defer updateMutex.Unlock()
	resp, err := http.Get(strapi + "/homepage")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := json.Unmarshal(body, &homepage); err != nil {
		panic(err)
	}
}
func main() {

	strapi = os.Getenv("STRAPI_SERVER")
	if strapi == "" {
		panic("STRAPI_SERVER not set. export STRAPI_SERVER=https://your.strapi.com")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)

	getArticles()
	getGlobal()
	getHomepage()
	getCategories()
	// example: /public/assets/images/example.png
	router.StaticFS("/public", http.FS(f))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", HomeContext{
			Articles:   articles,
			Global:     global,
			Homepage:   homepage,
			Categories: categories,
		})
	})

	router.GET("/posts/:slug", func(c *gin.Context) {
		slug := c.Params.ByName("slug")
		article, err := BySlug(slug)
		if err != nil {
			c.Error(err)
		}
		c.HTML(http.StatusOK, "article.tmpl", PostContext{
			Article:    article,
			Global:     global,
			Homepage:   homepage,
			Categories: categories,
		})
	})
	router.GET("/category/:slug", func(c *gin.Context) {
		slug := c.Params.ByName("slug")
		cat, err := CategoryBySlug(slug)
		if err != nil {
			c.Error(err)
		}
		c.HTML(http.StatusOK, "category.tmpl", CategoryContext{
			Global:   global,
			Homepage: homepage,
			Category: cat,
		})
	})
	router.POST("/content", func(c *gin.Context) {
		getArticles()
		getGlobal()
		getHomepage()
		getCategories()

		c.JSON(200, gin.H{"content": "Updated"})
	})
	router.GET("favicon.ico", func(c *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.Run(":" + port)
}
