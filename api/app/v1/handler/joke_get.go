package handler

import (
	"context"
	"io/ioutil"
	"strconv"
	"time"

	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/models"
	"jokes-bapak2-api/app/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func TodayJoke(c *fiber.Ctx) error {
	// check from redis if today's joke already exists
	// send the joke if exists
	// get a new joke if it's not, then send it.
	var joke models.Today
	err := redis.MGet(context.Background(), "today:link", "today:date", "today:image", "today:contentType").Scan(&joke)
	if err != nil {
		return err
	}

	eq, err := utils.IsToday(joke.Date)
	if err != nil {
		return err
	}

	if eq {
		c.Set("Content-Type", joke.ContentType)
		return c.Status(fiber.StatusOK).Send([]byte(joke.Image))
	} else {
		var link string
		err := db.QueryRow(context.Background(), "SELECT link FROM jokesbapak2 ORDER BY random() LIMIT 1").Scan(&link)
		if err != nil {
			return err
		}

		response, err := client.Get(link, nil)
		if err != nil {
			return err
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		now := time.Now().UTC().Format(time.RFC3339)
		err = redis.MSet(context.Background(), map[string]interface{}{
			"today:link":        link,
			"today:date":        now,
			"today:image":       string(data),
			"today:contentType": response.Header.Get("content-type"),
		}).Err()
		if err != nil {
			return err
		}

		c.Set("Content-Type", response.Header.Get("content-type"))
		return c.Status(fiber.StatusOK).Send(data)
	}

}

func SingleJoke(c *fiber.Ctx) error {
	checkCache, err := core.CheckJokesCache(memory)
	if err != nil {
		return err
	}

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(db)
		if err != nil {
			return err
		}
		err = memory.Set("jokes", jokes)
		if err != nil {
			return err
		}
	}

	link, err := core.GetRandomJokeFromCache(memory)
	if err != nil {
		return err
	}

	// Get image data
	response, err := client.Get(link, nil)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	c.Set("Content-Type", response.Header.Get("content-type"))
	return c.Status(fiber.StatusOK).Send(data)

}

func JokeByID(c *fiber.Ctx) error {
	checkCache, err := core.CheckJokesCache(memory)
	if err != nil {
		return err
	}

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(db)
		if err != nil {
			return err
		}
		err = memory.Set("jokes", jokes)
		if err != nil {
			return err
		}
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	link, err := core.GetCachedJokeByID(memory, id)
	if err != nil {
		return err
	}

	if link == "" {
		return c.Status(fiber.StatusNotFound).Send([]byte("Requested ID was not found."))
	}

	// Get image data
	response, err := client.Get(link, nil)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	c.Set("Content-Type", response.Header.Get("content-type"))
	return c.Status(fiber.StatusOK).Send(data)
}
