package models

type Position struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type GameState struct {
    Snake Position `json:"snake"`
    Apple Position `json:"apple"`
    Score int      `json:"score"`
}
