package api

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var Client *fasthttp.Client

func Initialize() {
	Client = &fasthttp.Client{
		MaxConnsPerHost: 20000,
	}
}

func GetRequest(url string) (res *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	res = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(url)
	err = Client.Do(req, res)

	if res.StatusCode() != fasthttp.StatusOK {
		fmt.Println("Status not found")
		return
	}
	//defer fasthttp.ReleaseResponse(res) // Only when you are done with body!

	return res, err
}

func PostUrl(c *fiber.Ctx) error {
	domain := UrlDomain{}
	domainInfo := Domain{}
	response := Response{}
	err := c.BodyParser(&domain)
	if err != nil {
		log.Println(err)
		return c.Status(400).SendString("Invalid request body")
	}

	var url string
	url = "https://rdap.verisign.com/com/v1/domain/" + domain.DomainParam

	res, err := GetRequest(url)
	if err != nil {

		return c.Status(500).SendString("Error")
	}

	err = json.Unmarshal(res.Body(), &domainInfo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	//eventDate alıncak
	//günümüz tarihinden çıkartılıp gün olarak döndürülecek

	events := domainInfo.Events
	eventDate := events[0].EventDate

	difference := int(time.Now().Sub(eventDate).Hours() / 24)
	response.EventDate = eventDate
	response.FinalResult = difference

	log.Println(response)
	return c.JSON(domain)

}
