type Videogame struct {
	id int `json:"id"`
	name string `json:"name"`
	category string `json:"category"`
	active_players int `json:"active_players"`
	size float `json:"size"`
	rating int `json:"rating"`
	downloads int `json:"downloads"`
}

type CreateVideogame struct {
	name string `json:"name"`
	category string `json:"category"`
	active_players int `json:"active_players"`
	size float `json:"size"`
	rating int `json:"rating"`
	downloads int `json:"downloads"`
}

type UpdateVideogame struct {
	name string `json:"name"`
	category string `json:"category"`
	active_players int `json:"active_players"`
	size float `json:"size"`
	rating int `json:"rating"`
	downloads int `json:"downloads"`
}