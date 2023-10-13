package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"html/template"
	"github.com/gofiber/fiber/v2"
	"github.com/dhanrajchaurasia/CP-GRIND/initializers"
	"github.com/dhanrajchaurasia/CP-GRIND/models"
)

// func UserProfile(c *fiber.Ctx) error {
// 	return c.Render("profile/index", fiber.Map{})
// }

func UserProfile(c *fiber.Ctx) error {
	msg := c.Cookies("message")
	cfProfile := template.HTML(c.Cookies("cfProfile"))
	c.ClearCookie("message")
	c.ClearCookie("cfProfile")
	data := fiber.Map{
		"User":       c.Cookies("username"),
		"CF_Profile": cfProfile,
		"Message":    msg,
	}
	return c.Render("profile/index", data)
}

func GetCFProfile(c *fiber.Ctx) error {
	url := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%v", c.FormValue("cf-handle"))
	response, err := http.Get(url)
	htmlBody := "Could not get the data!"
	c.Cookie(&fiber.Cookie{Name: "message", Value: "", Expires: time.Now()})
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/")
	}
	defer response.Body.Close()
	jsonData, _ := io.ReadAll(response.Body)
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/")
	}
	result := data["result"].([]interface{})
	if len(result) > 0 {
		user := result[0].(map[string]interface{})
		rank := user["rank"].(string)
		rating := user["rating"].(float64)
		htmlBody = fmt.Sprintf("CF Rank: %v<br>CF Rating: %v", rank, int(rating))
	}
	c.Cookie(&fiber.Cookie{
		Name:  "cfProfile",
		Value: htmlBody,
	})
	return c.Redirect("/")
}


func GetCodeforces(c *fiber.Ctx) error {
	url := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%v", c.FormValue("cf-handle"))
	response, err := http.Get(url)
	htmlBody := "Could not get the data!"
	c.Cookie(&fiber.Cookie{Name: "message", Value: "", Expires: time.Now()})
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/profile")
	}
	defer response.Body.Close()
	jsonData, _ := io.ReadAll(response.Body)
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/profile")
	}
	result := data["result"].([]interface{})
	if len(result) > 0 {
		user := result[0].(map[string]interface{})
		rank := user["rank"].(string)
		Handle := user["handle"].(string)
		rating := user["rating"].(float64)
		contributions := user["contribution"].(float64)
		lastOnline :=  user["lastOnlineTimeSeconds"].(float64)
		friendCount := user["friendOfCount"].(float64)
		maxRating := user["maxRating"].(float64)
		maxRank := user["maxRank"].(string)
		htmlBody = fmt.Sprintf("CF Rank: %v<br>CF Rating: %v<br>Contributions: %v<br>Last online: %v<br>Friends: %v<br>Max Rating: %v<br>Max Rank: %v", rank, int(rating), int(contributions), int(lastOnline), int(friendCount), int(maxRating), maxRank)

		//database
		profile := models.Profile{
			Handle: Handle,
			Rank: rank,
			Rating: rating,
			MaxRank: maxRank,
			MaxRating: maxRating,
		}
		// c.Cookie(&fiber.Cookie{
		// 	Name:  "cfProfile",
		// 	Value: htmlBody,
		// })
		initializers.CreateNewProfile(profile)
	// if err != nil {
	// 	return c.Render("profile/index", fiber.Map{
	// 		"Message": err.Error(),
	// 	})
	// }
	}


	c.Cookie(&fiber.Cookie{
		Name:  "cfProfile",
		Value: htmlBody,
	})
	return c.Redirect("/profile")
}

//atcoder


func GetAtcoderProfile(c *fiber.Ctx) error {
	url := fmt.Sprintf("https://kenkoooo.com/atcoder/atcoder-api/v3/user/ac_rank?user=%v", c.FormValue("cf-handle"))
	response, err := http.Get(url)
	htmlBody := "Could not get the data!"
	c.Cookie(&fiber.Cookie{Name: "message", Value: "", Expires: time.Now()})
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/profile")
	}
	defer response.Body.Close()
	jsonData, _ := io.ReadAll(response.Body)
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/profile")
	}
	// result := data["result"].([]interface{})
	if len(data) > 0 {
		rank := data["rank"].(float64)
		Handle := c.FormValue("cf-handle")
		count := data["count"].(float64)
		htmlBody = fmt.Sprintf("CF Rank: %v<br>Sumbissions: %v<br> %v", int(rank), int(count), Handle)

		//database
		atprofile := models.ATProfile{
			Handle: Handle,
			Rank: rank,
			Sumbissions: count,
		}
		initializers.AddATProfile(atprofile)
	}


	c.Cookie(&fiber.Cookie{
		Name:  "cfProfile",
		Value: htmlBody,
	})
	return c.Redirect("/profile")
}
