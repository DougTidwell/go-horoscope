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
			"friends and loved ones while you still can."
	case "taurus":
		Horoscope = "The stars suggest you leave town until " +
			"the whole thing blows over."
	case "gemini":
		Horoscope = "The stars stand corrected: turns out your " +
			"irrational fears are completely justified."
	case "cancer":
		Horoscope = "The stars say something hilarious will happen to you " +
			"today. Keep in mind that the stars have a really mean sense " +
			"of humor."
	case "leo":
		Horoscope = "The stars say your enemies will forgive you today. " +
			"They'll still press charges, but it won't be anything personal."
	case "virgo":
		Horoscope = "The stars say a steamroller figures prominently " +
			"in your future. You shouldn't worry, as long as you drive a " +
			"steamroller for a living."
	case "libra":
		Horoscope = "The stars say to stand up for yourself - today " +
			"is no time to back down. They recommend standing up to " +
			"someone small and unarmed."
	case "scorpio":
		Horoscope = "The stars say you should be grateful for all the " +
			"wonderful friends in your life. All your friends are imaginary, " +
			"so that's probably not a big deal."
	case "sagittarius":
		Horoscope = "The stars say a financial windfall could be yours " +
			"today. The stars also rolled their eyes and said, \"Yeah, right.\""
	case "capricorn":
		Horoscope = "The stars say you should treat yourself to a lavish " +
			"dinner tonight. Their exact words: \"As if it were your last meal.\""
	case "aquarius":
		Horoscope = "The stars say you should make a generous, " +
			"unexpected gift to someone you barely know today. It probably won't " +
			"get you out of hot water with Human Resources, but it couldn't hurt."
	case "pisces":
		Horoscope = "The stars say to approach a confrontation with " +
			"kid gloves and everything will be fine...for your opponent, who's " +
			"wearing brass knuckles."
	default:
		// Just use the default value for Horoscope
		Horoscope = "Sorry, you need to tell me your sign."
	}

	return c.JSON(&fiber.Map{
		"serviceName":   "Ominous",
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

	fmt.Printf("Ominous horoscope server running on port %d\n", portNumber)
	portSpec := ":" + strconv.Itoa(portNumber)
	err = app.Listen(portSpec)
	if err != nil {
		fmt.Print("Error loading server - is the port in use?")
		return
	}
}
