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
		Horoscope = "With the moon in the dining room of Cassiopeia, " +
			"avoid making any financial decisions today."
	case "taurus":
		Horoscope = "Despite Orion roaming the meadows of Capricorn, " +
			"the pen remains mightier than the sword. Still, the stars say you should " +
			"definitely not go to work without your sword."
	case "gemini":
		Horoscope = "Even though Saturn is retrograde, don't be afraid " +
			"to rock the boat today. Unless you're actually in a boat, in which case " +
			"you should just sit down and shut up."
	case "cancer":
		Horoscope = "As Libra scampers through the tulip fields of Zeus, " +
			"the stars suggest you ignore all the warning signs and have a great day."
	case "leo":
		Horoscope = "You wouldn't think it would be possible with the Pleiades " +
			"in the powder room of Sagittarius, but it's happening today anyway. " +
			"Good luck with that."
	case "virgo":
		Horoscope = "Neptune, relaxing in a hammock at Cygnus's place, " +
			"casually mentioned he knows what you did last summer."
	case "libra":
		Horoscope = "Hercules has nothing to do with the stars, but " +
			"he stopped by to suggest you stay in bed with the covers pulled up " +
			"to your ears until Pisces enters the back yard of Jupiter."
	case "scorpio":
		Horoscope = "While Venus enjoys a mimosa on Bacchus's front porch, " +
			"the stars say this is a perfect time to save money on car insurance."
	case "sagittarius":
		Horoscope = "As long as Taurus is swimming laps in the pool of " +
			"Olympus, the stars aren't giving back your passport."
	case "capricorn":
		Horoscope = "Mercury said to tell you not to call the cops, or else. " +
			"The stars assume you know what that's all about."
	case "aquarius":
		Horoscope = "Yes, you're an Aquarius, the water sign and all that, " +
			"but the stars emphatically said not to get on any boat of any size " +
			"until Mars starts wearing age-appropriate bathing suits."
	case "pisces":
		Horoscope = "Leda, enjoying some well-deserved downtime in the " +
			"guest house of Scorpio, says the wolf at your door would probably " +
			"make a great pet."
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

	fmt.Printf("Planetary Motion horoscope server running on port %d\n", portNumber)
	portSpec := ":" + strconv.Itoa(portNumber)
	err = app.Listen(portSpec)
	if err != nil {
		fmt.Print("Error loading server - is the port in use?")
		return
	}
}
