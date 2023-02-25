package data

/* Functions used for match analysis */

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
    . "HTStats/base"
)

type Data struct {
    Minions  []int
    Level    []int
    Xp       []int
    Gold     []int
    DmgDealt []int
    DmgTaken []int
}

type DataVal struct {
    Minions  int
    Level    int
    Xp       int
    Gold     int
    DmgDealt int
    DmgTaken int
}

type MatchData struct {
    CurrentGold int `json:"currentGold"`
    DamageStats struct {
        TotalDamageDone            int `json:"totalDamageDone"`
        TotalDamageDoneToChampions int `json:"totalDamageDoneToChampions"`
        TotalDamageTaken           int `json:"totalDamageTaken"`
    } `json:"damageStats"`
    GoldPerSecond       int `json:"goldPerSecond"`
    JungleMinionsKilled int `json:"jungleMinionsKilled"`
    Level               int `json:"level"`
    MinionsKilled       int `json:"minionsKilled"`
    ParticipantID       int `json:"participantId"`
    TotalGold           int `json:"totalGold"`
    Xp                  int `json:"xp"`
}

type Match struct {
    Metadata struct {
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		Frames        []struct {
			ParticipantFrames struct {
				Num1 MatchData `json:"1"`
				Num2 MatchData  `json:"2"`
				Num3 MatchData  `json:"3"`
				Num4 MatchData  `json:"4"`
				Num5 MatchData  `json:"5"`
				Num6 MatchData  `json:"6"`
				Num7 MatchData  `json:"7"`
				Num8 MatchData  `json:"8"`
				Num9 MatchData  `json:"9"`
				Num10 MatchData  `json:"10"`
			} `json:"participantFrames"`
			Timestamp int `json:"timestamp"`
		} `json:"frames"`
	} `json:"info"`
}

// Your API key here
var USER = User{Key : ""}
var CLIENT = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := CLIENT.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func GetDate() string {
    return time.Now().Format("2006-01-02,15:04")
}

func GetData(name string, server string) string {
    player := Player{Puuid : "", Name : name, Server : server, ServerWide : "americas"}
    
    getPlayerPUUID(&player)
    
    matches := getMatchList(&player)
    
    // start goroutine for each match...
    ch := make(chan Data)
    for _, m := range matches {
        go getMatch(m, &player, ch)
    }
    
    // analyze them all back together since dont want sequential http requests
    var data []Data
    for range matches {
        data = append(data, <-ch)
    }
    
    str := getAvgData(data)
    return str
}

func addAvg(num *int, sum *int, add int) {
    *sum += add
    *num++
}

func getAvgPoint(avg *Data, data []Data, i int) {
    num := DataVal{}
    sum := DataVal{}
    // loop through averages for matches vertically
    for _, entry := range data {
        if (len(entry.Minions) > i) {
            addAvg(&num.Minions, &sum.Minions, entry.Minions[i])
            addAvg(&num.Xp, &sum.Xp, entry.Xp[i])
            addAvg(&num.DmgDealt, &sum.DmgDealt, entry.DmgDealt[i])
            addAvg(&num.DmgTaken, &sum.DmgTaken, entry.DmgTaken[i])
            addAvg(&num.Gold, &sum.Gold, entry.Gold[i])
            addAvg(&num.Level, &sum.Level, entry.Level[i])
        }
    }
    
    avg.Minions = append(avg.Minions, sum.Minions / num.Minions)
    avg.Xp = append(avg.Xp, sum.Xp / num.Xp)
    avg.DmgDealt = append(avg.DmgDealt, sum.DmgDealt / num.DmgDealt)
    avg.DmgTaken = append(avg.DmgTaken, sum.DmgTaken / num.DmgTaken)
    avg.Gold = append(avg.Gold, sum.Gold / num.Gold)
    avg.Level = append(avg.Level, sum.Level / num.Level)

}

func getAvg(data []Data, gtr int) Data {
    avg := Data{}
    for i := 0; i < gtr; i++ {
        getAvgPoint(&avg, data, i)
    }
    return avg
}

func getAvgData(data []Data) string {
    gtr := 0
    for _, entry := range data {
        // Minions field b/c all same length
        if len(entry.Minions) > gtr { 
            gtr = len(entry.Minions)
        }
    }
    
    avg := getAvg(data, gtr)

    return dataToJSON(avg)
    
}

func dataToJSON(data Data) string {
    j, err := json.Marshal(data)
    if err != nil {
        CheckErr(err)
    }
     
    return string(j)
}

func getPlayerPUUID(player *Player) {

    str := fmt.Sprintf(
        "https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s",
        player.Server, player.Name, USER.Key)
    err := getJson(str, player)
    
    CheckErr(err)

}

func getMatchList(player *Player) []string {
    str := fmt.Sprintf(
        "https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?type=ranked&start=0&count=10&api_key=%s", 
        player.ServerWide, player.Puuid, USER.Key)

    var matches []string
    err := getJson(str, &matches)
    CheckErr(err)

    return matches
}

func getMatch(match string, player *Player, ch chan Data) {
    str := fmt.Sprintf(
        "https://%s.api.riotgames.com/lol/match/v5/matches/%s/timeline?api_key=%s",
        player.ServerWide, match, USER.Key)
        
    var data Match
    err := getJson(str, &data)
    CheckErr(err)
    
    ch <- analyzeMatch(player.Puuid, &data)

}

func getPlayerNum(puuid string, data *Match) int {
    for i, u := range data.Metadata.Participants {
        if (u == puuid) {
            return i
        }
    }
    return -1
}

func analyzeMatch(puuid string, data *Match) Data {

    var matchData Data
    
    playerNum := getPlayerNum(puuid, data)
    if (playerNum == -1) {
        return matchData
    }

    // only want stats for our player
    for _, d := range data.Info.Frames {
        // 10 players in game
        analyze := [10](MatchData){
            d.ParticipantFrames.Num1, d.ParticipantFrames.Num2,
            d.ParticipantFrames.Num3, d.ParticipantFrames.Num4,
            d.ParticipantFrames.Num5, d.ParticipantFrames.Num6,
            d.ParticipantFrames.Num7, d.ParticipantFrames.Num8,
            d.ParticipantFrames.Num9, d.ParticipantFrames.Num10}
        
        matchData.Minions = append(matchData.Minions, analyze[playerNum].MinionsKilled);
        matchData.Xp = append(matchData.Xp, analyze[playerNum].Xp);
        matchData.Gold = append(matchData.Gold, analyze[playerNum].CurrentGold);
        matchData.Level = append(matchData.Level, analyze[playerNum].Level);
        matchData.DmgDealt = append(matchData.DmgDealt, analyze[playerNum].DamageStats.TotalDamageDone);
        matchData.DmgTaken = append(matchData.DmgTaken, analyze[playerNum].DamageStats.TotalDamageTaken);
        
    }
    
    return matchData
}
