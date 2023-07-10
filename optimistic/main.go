package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"strconv"
	"strings"
)

func getHandler(c *fiber.Ctx) error {
	Horoscope := "Sorry, you need to tell me your sign."
	sign := strings.ToLower(c.Params("sign"))
	fmt.Printf("  Getting horoscope for %s\n", sign)

	switch sign {
	case "aries":
		Horoscope = "The stars suggest you say hello to your " +
			"friends and loved ones today."
	case "taurus":
		Horoscope = "The stars say it's time for a vacation!"
	case "gemini":
		Horoscope = "The stars say you should put your irrational " +
			"fears behind you and make the most of the day!"
	case "cancer":
		Horoscope = "The stars say something hilarious will happen to " +
			"you today. Laugh hard!"
	case "leo":
		Horoscope = "The stars say your enemies will forgive you today."
	case "virgo":
		Horoscope = "The stars say something will weigh heavily on your " +
			"mind. Stay positive!"
	case "libra":
		Horoscope = "The stars say to stand up for yourself - today is " +
			"no time to back down. You can do this!"
	case "scorpio":
		Horoscope = "The stars say you should be grateful for all the " +
			"wonderful friends in your life."
	case "sagittarius":
		Horoscope = "The stars say a financial windfall could be yours today."
	case "capricorn":
		Horoscope = "The stars say you should treat yourself to a lavish " +
			"dinner tonight. You've earned it!"
	case "aquarius":
		Horoscope = "The stars say you should make a generous, unexpected " +
			"gift to someone you barely know today."
	case "pisces":
		Horoscope = "The stars say to approach a confrontation with " +
			"kid gloves and everything will be fine."
	default:
		// Just use the default value for Horoscope
		Horoscope = "Sorry, you need to tell me your sign."
	}

	return c.JSON(&fiber.Map{
		"serviceName":   "Optimistic",
		"css":           "color: green;",
		"horoscopeText": Horoscope,
	})
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Authorization, Content-Length, X-Requested-With, Origin",
		AllowMethods:     "GET, POST, OPTIONS",
	}))
	app.Get("/horoscope/:sign", getHandler)

	portNumber := 3000
	var err error = nil
	portVar := os.Getenv("PORT")
	if len(portVar) > 0 {
		portNumber, err = strconv.Atoi(portVar)
	}
	if err != nil {
		portNumber = 3000
	}

	fmt.Printf("Optimistic horoscope server running on port %d\n", portNumber)
	portSpec := ":" + strconv.Itoa(portNumber)
	err = app.Listen(portSpec)
	if err != nil {
		fmt.Print("Error loading server - is the port in use?")
		return
	}
}
