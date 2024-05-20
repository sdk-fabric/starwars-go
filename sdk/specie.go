
// specie automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


import "time"

// A Species resource is a type of person or character within the Star Wars Universe.
type Specie struct {
    Name string `json:"name"`
    Classification string `json:"classification"`
    Designation string `json:"designation"`
    AverageHeight string `json:"average_height"`
    AverageLifespan string `json:"average_lifespan"`
    EyeColors string `json:"eye_colors"`
    HairColors string `json:"hair_colors"`
    SkinColors string `json:"skin_colors"`
    Language string `json:"language"`
    Homeworld string `json:"homeworld"`
    People []string `json:"people"`
    Films []string `json:"films"`
    Url string `json:"url"`
    Created time.Time `json:"created"`
    Edited time.Time `json:"edited"`
}
