
// film automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


import "time"

// A Film is a single film.
type Film struct {
    Title string `json:"title"`
    EpisodeId int `json:"episode_id"`
    OpeningCrawl string `json:"opening_crawl"`
    Director string `json:"director"`
    Producer string `json:"producer"`
    ReleaseDate time.Time `json:"release_date"`
    Species []string `json:"species"`
    Starships []string `json:"starships"`
    Vehicles []string `json:"vehicles"`
    Characters []string `json:"characters"`
    Planets []string `json:"planets"`
    Url string `json:"url"`
    Created time.Time `json:"created"`
    Edited time.Time `json:"edited"`
}
