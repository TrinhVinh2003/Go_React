package main

import (
	"database/sql"
	"fmt"
	"log"
	"project/vnexpress/crawls"
	"project/vnexpress/database"
	"project/vnexpress/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func main() {
// 	// fmt.Println("hello")
// 	// if db.DBerr != nil{
// 	// 	panic(db.DBerr.Error())
// 	// }
// 	// defer db.DBconnect.Close()

// 	// insert, err := db.DBconnect.Query("INSERT INTO test(ID,Name) VALUES(2,'VINH')")
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// defer insert.Close()

// 	// crawls.Crawl()

// }

type Article struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func init() {
	database.ConnectDB()
	crawls.Crawl()
}
func main() {

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Static("/static", "./static")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	app.Listen(":8000")

	// c := colly.NewCollector(colly.AllowedDomains("www.vnexpress.net", "vnexpress.net"))
	// // contentCollector := c.Clone()
	// var articles []Article

	// c.OnHTML("article.item-news", func(e *colly.HTMLElement) {
	// 	title := e.ChildText(".title-news a")
	// 	link := e.ChildAttr(".title-news a", "href")
	// 	image := e.ChildAttr("img[itemprop=contentUrl]", "src")
	// 	description := e.ChildText("p.description")
	// 	article := Article{Title: title, Link: link, Image: image, Description: description}

	// 	articles = append(articles, article)
	// 	// fmt.Println(link)

	// })
	// c.Visit("https://vnexpress.net/the-thao/euro-2024/tin-tuc")
	// // contentCollector.OnHTML("div.container", func(h *colly.HTMLElement) {
	// // 	description := h.DOM.Find("p.description").Text()
	// // 	// fmt.Println(description)
	// // 	// if description == "" {
	// // 	// 	log.Printf("ERROR ")
	// // 	// 	description = "không lấy được descript"
	// // 	// }
	// // 	fmt.Println(description)
	// // 	for i := range articles {

	// // 		articles[i].Description = description
	// // 		break

	// // 	}
	// // })

	// saveToMysql(articles)

	// http.HandleFunc("/api/vnexpress", func(w http.ResponseWriter, r *http.Request) {

	// 	jsonData, err := json.Marshal(articles)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(jsonData)
	// })

	// log.Fatal(http.ListenAndServe(":8080", nil))
}

// func saveToMysql(articles []Article) {
// 	// connect Mysql
// 	db, err := sql.Open("mysql", "root:@tcp(localhost)/test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// delete data cũ
// 	_, err = db.Exec("DELETE FROM articles")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// insert data mới
// 	insertDB, err := db.Prepare("INSERT INTO articles(title, link, image_url,content) VALUES(?, ?, ?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer insertDB.Close()

// 	for _, article := range articles {
// 		_, err := insertDB.Exec(article.Title, article.Link, article.Image, article.Description)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	fmt.Println("Success!")
// }
