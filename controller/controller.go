package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
	"unicode"
    
	"github.com/gin-gonic/gin"
)

func PingV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ValidatePasswordV1(c *gin.Context) {
	password := c.PostForm("password")
	var IsLength, IsUpper, IsLower, IsNumber, IsSpecial bool
	var Errors []string
	var strength int = 0
    
	if len(password) >= 8 {
		IsLength = true
	}
	
	for _, v := range password {
		switch {
		case unicode.IsNumber(v):
			IsNumber = true

		case unicode.IsUpper(v):
			IsUpper = true

		case unicode.IsLower(v):
			IsLower = true

		case unicode.IsPunct(v) || unicode.IsSymbol(v):
			IsSpecial = true

		}
	}	

	if IsLength {
		strength = 1
	} else {
		Errors = append(Errors, "less than 8")
	} 

	if IsNumber {
		strength += 1
	} else {
         Errors = append(Errors, "Number is missing")
	}

	if IsLower {
		strength += 1
	} else {
		Errors = append(Errors, "Lower is missing")
	}

	if IsUpper {
		strength += 1
	} else {
		Errors = append(Errors, "Upper is missing")
	}

	if IsSpecial {
		strength += 1
	} else {
		Errors = append(Errors, "Special is missing")
	}
	
	
	if strength >= 4 {
		fmt.Println("strong password")
		c.JSON(http.StatusBadRequest, gin.H{
			"strength": strength,
			"valid":    "false",
			"Errors": Errors,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"strength": strength,
			"valid":    "true",
		})
	}	
	fmt.Println(Errors)
	fmt.Println(strength)
	fmt.Println(IsLength, IsLower, IsNumber,IsUpper, IsSpecial)
}

func GetHashV1(c *gin.Context) {
	password := c.Query("password")
	hashed := md5.Sum([]byte(password))
	current_time := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"hash":      hex.EncodeToString(hashed[:]),
		"timestamp": current_time,
	})
}

