
// film_collection automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


type FilmCollection struct {
    Count int `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Results []Film `json:"results"`
}
