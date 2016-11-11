package models

import "strings"

// Team maps directly to the teams database table.
type Team struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	URLSlug  string `json:"url_slug"`
	Homepage string `json:"homepage"`
	IconURL  string `json:"icon_url"`
	Lead     User   `json:"lead"`
}

func (t *Team) String() string {
	return jsonString(t)
}

// NewTeam makes a new team generating the slug based on it's name.
func NewTeam(name, homepage, iconURL string) Team {
	slug := strings.ToLower(name)
	slug = strings.Replace(slug, " ", "-", -1)

	return Team{
		Name:     name,
		IconURL:  iconURL,
		Homepage: homepage,
		URLSlug:  slug,
	}
}
