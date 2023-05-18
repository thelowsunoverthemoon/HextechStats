package main

import (
    "github.com/gin-gonic/gin"
    "HTStats/profile"
    . "HTStats/endpoint"
    . "HTStats/base"
)

func main() {

    // connect to database
    err := profile.ConnectDatabase()
    CheckErr(err)
    
    // router settings
    r := gin.Default()

    // to allow API requests by change COR settings
    r.Use(CORSMiddleware())
    
    v1 := r.Group("/api/v1")
    {
      v1.GET("profile", GetProfiles)
      v1.GET("profile/:name/:server", GetProfile)
      v1.POST("profile", AddProfile)
      v1.PUT("profile", UpdateProfile)
      v1.PUT("profile/delete", DeleteProfile)
      
    }
    
    r.Run()
   
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}





