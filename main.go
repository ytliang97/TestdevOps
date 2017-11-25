package main

import (
    "net/http"
     "strconv"
	 
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	
	 e.GET("/add/:num1/:num2", func(c echo.Context) error {
        num1Str := c.Param("num1")
        num2Str := c.Param("num2")
        num1, err := strconv.Atoi(num1Str)
        if err != nil {
            return c.String(http.StatusUnprocessableEntity, err.Error())
        }
        num2, err := strconv.Atoi(num2Str)
        if err != nil {
            return c.String(http.StatusUnprocessableEntity, err.Error())
        }
        result := Add(num1, num2)
        return c.String(http.StatusOK, strconv.Itoa(result))
    })

	
    e.Logger.Fatal(e.Start(":1323"))
	
}

func Add(num1, num2 int) int {
    return num1 + num2
}
