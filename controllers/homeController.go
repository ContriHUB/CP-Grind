package controllers

import (
	"html/template"
	"time"

	"github.com/dhanrajchaurasia/CP-GRIND/initializers"
	"github.com/dhanrajchaurasia/CP-GRIND/models"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	msg := c.Cookies("message")
	cfProfile := template.HTML(c.Cookies("cfProfile"))
	c.ClearCookie("message")
	c.ClearCookie("cfProfile")
	data := fiber.Map{
		"User":       c.Cookies("username"),
		"CF_Profile": cfProfile,
		"Message":    msg,
	}
	return c.Render("home/index", data)
}
func LoginPage(c *fiber.Ctx) error {
	auth := c.Cookies("authorization_token")
	if initializers.IsValidToken(auth) {
		return c.Redirect("/")
	}
	return c.Render("home/login", fiber.Map{
		"Message": "Welcome to CP Grind!",
	})
}

func Signup(c *fiber.Ctx) error {
	fname := c.FormValue("fname")
	lname := c.FormValue("lname")
	username := c.FormValue("username")
	email := c.FormValue("email")
	pass := c.FormValue("password")
	cpass := c.FormValue("cpassword")

    // Check if any field is empty
	  if fname == "" || lname == "" || username == "" || email == "" || pass == "" || cpass == "" {
        return c.Render("home/login", fiber.Map{
            "Message": "Please fill all the fields!",
        })
    }
  //check the length of password
	  if len(pass) < 8 {
        return c.Render("home/login", fiber.Map{
            "Message": "Password must be at least 8 characters long!",
        })
    }

	 // Check email format 
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return c.Render("home/login", fiber.Map{
            "Message": "Invalid email address format!",
        })
    }
    
	  // Check password strength
    passwordStrength, improvementTips := getPasswordStrength(pass)

    if passwordStrength == "very weak" {
        return c.Render("home/login", fiber.Map{
            "Message": "Password is very weak. " + improvementTips,
        })
    } else if passwordStrength == "weak" {
        return c.Render("home/login", fiber.Map{
            "Message": "Password is weak. " + improvementTips,
        })
    } else if passwordStrength == "strong" {
        return c.Render("home/login", fiber.Map{
            "Message": "Password is strong. " + improvementTips,
        })
    } else if passwordStrength == "very strong" {
        return c.Render("home/login", fiber.Map{
            "Message": "Password is very strong. " + improvementTips,
        })
    }

	// Check if password and confirm password match
	if pass != cpass {
		return c.Render("home/login", fiber.Map{
			"Message": "Password didn't match!",
		})
	}

	user := models.User{
		FirstName:  fname,
		SecondName: lname,
		Username:   username,
		Email:      email,
		Password:   pass,
	}
	err := initializers.CreateNewUser(user)
	if err != nil {
		return c.Render("home/login", fiber.Map{
			"Message": err.Error(),
		})
	}
	return c.Render("home/login", fiber.Map{
		"Message": "Your account has been created successfully, Login to Continue!",
	})
}

func getPasswordStrength(password string) (string, string) {
    lengthRegex := regexp.MustCompile(`^.{8,}$`)
    uppercaseRegex := regexp.MustCompile(`[A-Z]`)
    lowercaseRegex := regexp.MustCompile(`[a-z]`)
    digitRegex := regexp.MustCompile(`[0-9]`)
    specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

    // Check each criteria and stored the feedback
    feedback := []string{}

    if !lengthRegex.MatchString(password) {
        feedback = append(feedback, "Password should be at least 8 characters.")
    }

    if !uppercaseRegex.MatchString(password) {
        feedback = append(feedback, "It should contain at least one uppercase letter.")
    }

    if !lowercaseRegex.MatchString(password) {
        feedback = append(feedback, "It should contain at least one lowercase letter.")
    }

    if !digitRegex.MatchString(password) {
        feedback = append(feedback, "It should contain at least one digit.")
    }

    if !specialCharRegex.MatchString(password) {
        feedback = append(feedback, "It should contain at least one special character.")
    }

    // Determined password strength based on the number of criteria met
    strength := ""
    if len(feedback) == 0 {
        strength = "very strong"
    } else if len(feedback) <= 1 {
        strength = "strong"
    } else if len(feedback) <= 3 {
        strength = "weak"
    } else {
        strength = "very weak"
    }

    return strength, strings.Join(feedback, " ")
}

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	token, err := initializers.IsUserPresent(username, password)
	if err != nil {
		return c.Render("home/login", fiber.Map{
			"Message": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:  "authorization_token",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "username",
		Value: username,
	})
	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "authorization_token",
		Value:   "",
		Expires: time.Now(),
	})
	return c.Redirect("/login")
}

func NotFound(c *fiber.Ctx) error {
	return c.Render("home/404", fiber.Map{})
}
