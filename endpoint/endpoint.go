package endpoint

/* Functions called when API endpoints are queries */

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "HTStats/profile"
    . "HTStats/data"
    . "HTStats/base"
)


type ProfList struct {
    Prof []ProfId `json:"list"`
}

type ProfId struct {
    Server string    `json:"server"`
    Name   string `json:"name"`
}

func GetProfiles(c *gin.Context) {

    profiles, err := profile.GetProfiles()
    CheckErr(err)

    if profiles == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
        return
    } else {
        c.JSON(http.StatusOK, gin.H{"data": profiles})
    }
}

func GetProfile(c *gin.Context) {

    server := c.Param("server")
    name := c.Param("name")
    
    data, err := profile.GetProfile(name, server)
    CheckErr(err)

    if data == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": data})
    }
}

func AddProfile(c *gin.Context) {

    var json ProfId

    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    date := GetDate()
    data := GetData(json.Name, json.Server)
    success, err := profile.AddProfile(json.Name, json.Server, date, data)

    if success {
        c.JSON(http.StatusOK, gin.H{"data" : data, "message": "Success"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
    }
}

func UpdateProfile(c *gin.Context) {
    var json ProfList

    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    date := GetDate()
    
    for _, entry := range json.Prof {
        data := GetData(entry.Name, entry.Server)
        success, err := profile.UpdateProfile(entry.Name, entry.Server, date, data)

        if success {
            c.JSON(http.StatusOK, gin.H{"message": "Success"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
        }
    }
}

func DeleteProfile(c *gin.Context) {
    var json ProfList

    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for _, entry := range json.Prof {
        success, err := profile.DeleteProfile(entry.Name, entry.Server)
        if success {
            c.JSON(http.StatusOK, gin.H{"message": "Success"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
        }
    }
    
}
