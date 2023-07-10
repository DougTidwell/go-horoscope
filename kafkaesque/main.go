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
		Horoscope = "You will be arrested today. You will never be told " +
			"the nature of the charges against you. " +
			"Your life will be bleak and miserable."
	case "taurus":
		Horoscope = "You will feel your father's disapproval keenly today. " +
			"Pretty much like every other day. Your life will be bleak and miserable."
	case "gemini":
		Horoscope = "After tomorrow, let's just say you'll be wearing shoes " +
			"six at a time. Shop accordingly. Your life will be bleak and miserable."
	case "cancer":
		Horoscope = "You will write the person you love, telling them your " +
			"relationship will never happen. They will never write back. "
	case "leo":
		Horoscope = "The excitement of finding a new form of disappointment " +
			"will fade quickly. Your life will be bleak and miserable."
	case "virgo":
		Horoscope = "You will have a sudden insight that will bring you " +
			"lasting peace and contentment. Just kidding. " +
			"Your life will be bleak and miserable."
	case "libra":
		Horoscope = "Your lawyer will assure you that this will all be " +
			"over soon. He is correct, but not in a good way. " +
			"Your life will be bleak and miserable."
	case "scorpio":
		Horoscope = "What happens today won't be any more bearable if " +
			"the stars tell you about it beforehand, so let's leave it at that. " +
			"Your life will be bleak and miserable."
	case "sagittarius":
		Horoscope = "Not everybody would keep pursuing an unobtainable goal, " +
			"but you seem to have a real flair for it. " +
			"Your life will be bleak and miserable."
	case "capricorn":
		Horoscope = "You will start the day with a really good cappuccino. " +
			"Like, seriously good. But it'll all be downhill from there. " +
			"Your life will be bleak and miserable."
	case "aquarius":
		Horoscope = "You don't have a dog, but even if you did, it " +
			"wouldn't like you. Your life will be bleak and miserable."
	case "pisces":
		Horoscope = "Your city will have an amazing zoo, but you'll never " +
			"live to see it. Your life will be bleak and miserable."
	default:
		// Just use the default value for Horoscope
		Horoscope = "Sorry, you need to tell me your sign."
	}
	return c.JSON(&fiber.Map{
		"serviceName":   "Kafkaesque",
		"css":           "color: red;",
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

	fmt.Printf("Kafkaesque horoscope server running on port %d\n", portNumber)
	portSpec := ":" + strconv.Itoa(portNumber)
	err = app.Listen(portSpec)
	if err != nil {
		fmt.Print("Error loading server - is the port in use?")
		return
	}
}
