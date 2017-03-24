package musicbrainz

type Result struct {
	Area     string  `json:"area,omitempty"`
	Name     string  `json:"name,omitempty"`
	SortName string  `json:"sort_name,omitempty"`
	Gid      string  `json:"gid,omitempty"`
	Type     string  `json:"type,omitempty"`
	ID       int     `json:"id,omitempty"`
	Begin    Begin   `json:"begin,omitempty"`
	Aliases  Aliases `json:"aliases,omitempty"`
	Tags     Tags    `json:"tags,omitempty"`
	Rating   Rating  `json:"rating,omitempty"`
}

type Alias struct {
	Name     string `json:"name"`
	SortName string `json:"sort_name"`
}

type Aliases []Alias

type Begin struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
}

type Tag struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

type Tags []Tag

type Rating struct {
	Count int `json:"count"`
	Value int `json:"value"`
}
