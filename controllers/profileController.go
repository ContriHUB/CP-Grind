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
    "github.com/enescakir/emoji"
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
	IfPresent := initializers.IsProfilePresent(c.FormValue("cf-handle"))
	// var info models.Profile
	// var err 
	if IfPresent > 0 {
		info := initializers.CFProfile(c.FormValue("cf-handle"))
		Handle := info.Handle
		rank := info.Rank
		rating := info.Rating
		maxRank := info.MaxRank
		maxRating := info.MaxRating
		htmlBody := fmt.Sprintf("%v<br>%v CF Rank: %v<br>%v CF Rating: %v<br>%v Max Rating: %v<br>%v Max Rank: %v", Handle, emoji.ManTechnologist, rank, emoji.SmilingFaceWithSunglasses, int(rating), emoji.SportsMedal, int(maxRating), emoji.Trophy, maxRank)
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		
	}
	if IfPresent == 0 {
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
		htmlBody = fmt.Sprintf("%v<br>%v CF Rank: %v<br>%v CF Rating: %v<br>%v Contributions: %v<br>Last online: %v<br>%v Friends: %v<br>%v Max Rating: %v<br>%v Max Rank: %v", Handle, emoji.ManTechnologist, rank, emoji.SmilingFaceWithSunglasses, int(rating), emoji.Star, int(contributions), int(lastOnline/31536000),emoji.GlowingStar, int(friendCount),emoji.SportsMedal, int(maxRating), emoji.Trophy, maxRank)

		//database
		profile := models.Profile{
			Handle: Handle,
			Rank: rank,
			Rating: rating,
			MaxRank: maxRank,
			MaxRating: maxRating,
			Email: c.FormValue("cf-email"),
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
	
	}
	return c.Redirect("/profile")
}

//atcoder


func GetAtcoderProfile(c *fiber.Ctx) error {
	IfPresent := initializers.IsAtProfilePresent(c.FormValue("cf-handle"))
	// var info models.Profile
	// var err 
	if IfPresent > 0 {
		info := initializers.ATProfile(c.FormValue("cf-handle"))
		Handle := info.Handle
		rank := info.Rank
		Sumbissions := info.Sumbissions
		htmlBody := fmt.Sprintf("%v<br>%v AT Rank: %v<br>%v Sumbissions: %v",Handle, emoji.Trophy, int(rank),emoji.Star, int(Sumbissions))
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		
	}
	if IfPresent == 0 {
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
		htmlBody = fmt.Sprintf("%v<br>%v AT Rank: %v<br>%v Sumbissions: %v",Handle, emoji.Trophy, int(rank),emoji.Star, int(count))

		//database
		atprofile := models.ATProfile{
			Handle: Handle,
			Rank: rank,
			Sumbissions: count,
			Email: c.FormValue("at-email"),
		}
		initializers.AddATProfile(atprofile)
	}


	c.Cookie(&fiber.Cookie{
		Name:  "cfProfile",
		Value: htmlBody,
	})
	}
	return c.Redirect("/profile")
}