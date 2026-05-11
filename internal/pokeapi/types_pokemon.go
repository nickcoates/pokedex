package pokeapi

type RespPokemon struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"base_experience"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Types      []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}
