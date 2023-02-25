package base

/* Common Functions used by other modules */

import (
    "log"
)


type User struct {
    Key string 
}

type Player struct {
    Server     string
    ServerWide string
    Name       string
    Puuid      string
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
