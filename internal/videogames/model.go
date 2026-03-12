package videogames

type Videogame struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	ActivePlayers int     `json:"active_players"`
	Size          float64 `json:"size"`
	Rating        int     `json:"rating"`
	Downloads     int     `json:"downloads"`
}

type CreateVideogame struct {
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	ActivePlayers int     `json:"active_players"`
	Size          float64 `json:"size"`
	Rating        int     `json:"rating"`
	Downloads     int     `json:"downloads"`
}

type UpdateVideogame struct {
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	ActivePlayers int     `json:"active_players"`
	Size          float64 `json:"size"`
	Rating        int     `json:"rating"`
	Downloads     int     `json:"downloads"`
}
