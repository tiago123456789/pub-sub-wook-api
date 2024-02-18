package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tiago123456789/pub-sub-wook-api/internal/config"
	"github.com/tiago123456789/pub-sub-wook-api/internal/model"
	"github.com/tiago123456789/pub-sub-wook-api/internal/repository"

	"github.com/joho/godotenv"
)

type URLSubscribed struct {
	Headers map[string]string      `json:"headers"`
	Method  string                 `json:"method"`
	Data    map[string]interface{} `json:"data"`
	Url     string                 `json:"url"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.Start()
	if err != nil {
		log.Fatal(err)
	}

	urlProducerRepository := repository.NewUrlProducerRepository(db)

	queue := config.NewQueue()
	queueUrl := os.Getenv("QUEUE_URL")

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Post("/urls-producers", func(c *fiber.Ctx) error {
		urlProducer := model.UrlProducer{
			Enabled: true,
			Key:     uuid.NewString(),
		}
		urlProducerCreated, err := urlProducerRepository.Create(urlProducer)
		if err != nil {
			log.Fatal(err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Interval server error",
			})
		}

		return c.Status(200).JSON(urlProducerCreated)
	})

	app.Post("/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")
		key := c.Get("key")
		// tokenConverted, _ := strconv.Atoi(token)
		// urlProducer, _ := urlProducerRepository.GetById(tokenConverted)
		// if urlProducer.Id == 0 || urlProducer.Key != key {
		// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		// 		"errors": "Invalid URL.",
		// 	})
		// }

		// payload := new(model.Payload)
		// json.Unmarshal(c.Body(), payload)

		// payload.Token = token
		dataJson := `{ "key":"` + key + `", "token":` + token + `,`

		event := c.Query("event")
		if len(event) > 0 {
			dataJson += `"event":` + c.Query("event") + `,`
			// payload.Event = c.Query("event")
		}

		dataJson += ` "payload": ` + string(c.Body()) + "}"
		// dataJson, _ := json.Marshal(payload)

		// fmt.Println(dataJson)
		err = queue.SendMessage(config.Message{
			Queue:   queueUrl,
			Message: string(dataJson),
		})

		if err != nil {
			log.Fatal(err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Interval server error",
			})
		}

		return c.SendStatus(202)
	})

	app.Listen(":3000")
}
