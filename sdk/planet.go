
// planet automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


import "time"

// A Planet resource is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
type Planet struct {
    Name string `json:"name"`
    Diameter string `json:"diameter"`
    RotationPeriod string `json:"rotation_period"`
    OrbitalPeriod string `json:"orbital_period"`
    Gravity string `json:"gravity"`
    Population string `json:"population"`
    Climate string `json:"climate"`
    Terrain string `json:"terrain"`
    SurfaceWater string `json:"surface_water"`
    Residents []string `json:"residents"`
    Films []string `json:"films"`
    Url string `json:"url"`
    Created time.Time `json:"created"`
    Edited time.Time `json:"edited"`
}
