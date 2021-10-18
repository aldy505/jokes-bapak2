package joke

import (
	"errors"
	"io/ioutil"
	core "jokes-bapak2-api/app/core/joke"
	"jokes-bapak2-api/app/core/schema"
	"jokes-bapak2-api/app/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) TodayJoke(c *fiber.Ctx) error {
	// check from handler.Redis if today's joke already exists
	// send the joke if exists
	// get a new joke if it's not, then send it.
	var joke Today
	err := d.Redis.MGet(c.Context(), "today:link", "today:date", "today:image", "today:contentType").Scan(&joke)
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
	}

	link, err := core.GetRandomJokeFromDB(d.DB, c.Context())
	if err != nil {
		return err
	}

	response, err := d.HTTP.Get(link, nil)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	err = d.Redis.MSet(c.Context(), map[string]interface{}{
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

func (d *Dependencies) SingleJoke(c *fiber.Ctx) error {
	checkCache, err := core.CheckJokesCache(d.Memory)
	if err != nil {
		return err
	}

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(d.DB, c.Context())
		if err != nil {
			return err
		}

		err = d.Memory.Set("jokes", jokes)
		if err != nil {
			return err
		}
	}

	link, err := core.GetRandomJokeFromCache(d.Memory)
	if err != nil && !errors.Is(err, schema.ErrEmpty) {
		return err
	}

	// Get image data
	response, err := d.HTTP.Get(link, nil)
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

func (d *Dependencies) JokeByID(c *fiber.Ctx) error {
	checkCache, err := core.CheckJokesCache(d.Memory)
	if err != nil {
		return err
	}

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(d.DB, c.Context())
		if err != nil {
			return err
		}

		err = d.Memory.Set("jokes", jokes)
		if err != nil {
			return err
		}
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	link, err := core.GetCachedJokeByID(d.Memory, id)
	if err != nil {
		return err
	}

	if link == "" {
		return c.
			Status(fiber.StatusNotFound).
			Send([]byte("Requested ID was not found."))
	}

	// Get image data
	response, err := d.HTTP.Get(link, nil)
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
