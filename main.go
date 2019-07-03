package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/peterbourgon/diskv"
)

var apikey = "1234"
var listen = "127.0.0.1:9991"

func main() {

	// Simplest transform function: put all the data files into the base dir.
	flatTransform := func(s string) []string { return []string{} }

	if err := os.Mkdir("cache", 0644); err != nil && !os.IsExist(err) {
		panic(err)
	}

	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	d := diskv.New(diskv.Options{
		BasePath:     "cache",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})

	e := echo.New()
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return strings.EqualFold(apikey, key), nil
		},
	}))

	e.POST("/search", func(c echo.Context) error {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}

		url, isset := m["image_url"].(string)
		if !isset || url == "" {
			return errors.New("missing url")
		}

		hash, isset := m["image_md5"].(string)
		if !isset || hash == "" {
			return errors.New("missing hash")
		}

		val, err := d.Read(hash)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}

		if os.IsNotExist(err) || len(val) == 0 {
			res, err := search(url, c)
			if err != nil {
				return err
			}

			val, err = json.Marshal(res)
			if err != nil {
				panic(err)
			}
			if err := d.Write(hash, val); err != nil {
				panic(err)
			}
		}

		if len(val) > 0 {
			return c.JSONBlob(http.StatusOK, val)
		}

		return c.NoContent(http.StatusNotFound)
	})

	e.Logger.Fatal(e.Start(listen))
}
