package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/platform/cache"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/platform/database"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/utils"
	"github.com/gofiber/fiber/v2"
)

type Today struct {
	link string `redis:"link"`
	date string `redis:"date"`
}

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
var db = database.New()
var redis = cache.New()

func TodayJoke(c *fiber.Ctx) error {
	// check from redis if today's joke already exists
	// send the joke if exists
	// get a new joke if it's not, then send it.
	var joke Today
	err := redis.MGet(context.Background(), "today:link", "today:date").Scan(&joke)
	if err != nil {
		return err
	}

	eq, err := utils.IsToday(joke.date)
	if err != nil {
		return err
	}

	if eq {
		c.Attachment(joke.link)
		return c.SendStatus(200)
	} else {
		var link string
		err := db.QueryRow(context.Background(), "SELECT link FROM jokesbapak2 WHERE random() < 0.01 LIMIT 1").Scan(&link)
		if err != nil {
			return err
		}
		now := strconv.Itoa(int(time.Now().Unix()))
		err = redis.MSet(context.Background(), "today:link", link, "today:date", now).Err()
		if err != nil {
			return err
		}
		c.Attachment(link)
		return c.SendStatus(200)
	}

}

func SingleJoke(c *fiber.Ctx) error {
	// get a joke from db
	// fetch the image url
	// send the image as proxied file
	var link string
	err := db.QueryRow(context.Background(), "SELECT link FROM jokesbapak2 WHERE random() < 0.01 LIMIT 1").Scan(&link)
	if err != nil {
		return err
	}
	c.Attachment(link)
	return c.SendStatus(200)
}

func JokeByID(c *fiber.Ctx) error {
	// get a joke from db by id
	// fetch image url
	// send the image as proxied file
	var link string
	err := db.QueryRow(context.Background(), "SELECT link FROM jokesbapak2 WHERE id = ?", c.Params("id")).Scan(link)
	if err != nil {
		return err
	}
	if link == "" {
		return c.Status(404).Send([]byte("Requested ID was not found."))
	}
	c.Attachment(link)
	return c.SendStatus(200)
}
