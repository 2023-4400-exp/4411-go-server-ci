package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/", Hello())
    e.GET("/api/", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
    	return c.String(http.StatusOK, "4411: 室町だと母には二度会うけれども、父には一度も会わないものは？ ver.2")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "学生番号", "message": "一言"})
    }
}
