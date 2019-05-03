package Elgamal

// package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9")

type (
	Sum struct {
		A float64 `json:"a" validate:"required|eq=0"`
		B float64 `json:"b" validate:"required|eq=0"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest Api by Go")
	})
	e.GET("/getServerTime", getServerTime)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/sumAB", sumAB)
	e.Logger.Fatal(e.Start(":7777")) // port number
}

func getServerTime(c echo.Context) error {
	t := time.Now().Format("02-01-2006 03:04:05")
	return c.String(http.StatusOK, "Time : "+t)
}

func sumAB(c echo.Context) (err error) {
	ab := new(Sum)
	err = c.Bind(ab)
	if err != nil {
		return c.JSON(http.StatusOK, "A or B is not number")
	}
	if err = c.Validate(ab); err != nil {
		return c.JSON(http.StatusOK, "A or B is error")
	}
	sum := ab.A + ab.B
	return c.JSON(http.StatusOK, sum)
}

// err = validateJSONSchema(schemaReceipt, string(receipt))
// if err != nil {
// 	return receiptAsStruct, errors.New(methodName + " " + err.Error())
// }

// const schemaReceipt string = `{
// 	"$schema": "http://json-schema.org/draft-04/schema#",
// 	"title": "receipt",
// 	"description": "receipt",
// 	"type": "object",
// 	"properties": {
// 		"receiptNo": {
// 			"description": "receiptNo",
// 			"type": "string"
// 		},
// 		"treatment": {
// 			"description": "treatment",
// 			"type": "array"
// 			"item": {
// 				"description": "treatment",
// 				"type": "object"
// 				"properties": {
// 					"item": {
// 						"description": "item",
// 						"type": "string"
// 					},
// 					"quantity": {
// 						"description": "quantity",
// 						"type": "number"
// 					},
// 					"price": {
// 						"description": "price",
// 						"type": "number"
// 					},
// 					"amount": {
// 						"description": "amount",
// 						"type": "number"
// 					}
// 				},
// 				"required": ["item", "price", "amount"],
// 			}
// 		},
// 		"total": {
// 			"description": "total",
// 			"type": "number"
// 		}
// 	},
// 	"required": ["receiptNo", "treatment", "total"],
// 	"additionalProperties": false
// }`
