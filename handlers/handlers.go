package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/shortUrl/database"
	"github.com/zeimedee/shortUrl/models"
	"github.com/zeimedee/shortUrl/utils"
)

func Redirect(c *fiber.Ctx) error {
	link := new(models.Urls)
	short := c.Params("short")
	// database.DB.Db.Where("shorturl = ?", short).Find(&link)
	database.DB.Db.Raw("SELECT * FROM urls WHERE short_url = ?", short).Scan(&link)
	return c.Redirect(link.InitialUrl)

}

func ShortenUrl(c *fiber.Ctx) error {
	initialUrl := new(models.InitialUrl)
	shortUrl := new(models.ShortUrl)

	if err := c.BodyParser(initialUrl); err != nil {
		return c.Status(400).JSON("Bad Request")
	}
	host := c.BaseURL()
	t := utils.Shortener()
	shortUrl.Url = host + "/" + t
	entry := models.Urls{
		InitialUrl: initialUrl.Url,
		ShortUrl:   t,
	}
	database.DB.Db.Create(&entry)
	return c.Status(200).JSON(shortUrl.Url)
}

func AllBooks(c *fiber.Ctx) error {
	books := []models.Urls{}
	database.DB.Db.Find(&books)

	return c.Status(200).JSON(books)
}
