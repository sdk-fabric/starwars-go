
# starwars-go

This [SDK](https://github.com/sdk-fabric/starwars-go) is managed by the [SDK Fabric](https://sdk-fabric.org/) project, a global infrastructure to
automatically generate SDKs for every API.

You can find more information about this SDK at [TypeHub](https://typehub.cloud/):
https://app.typehub.cloud/d/sdkfabric/starwars

## Usage

```go
import (
	"github.com/sdk-fabric/starwars-go/sdk"
)

var client, _ = sdk.Build("[access_token]");

// Get all the people.
response, err := client.People().Getall("search")

// Get a specific people.
response, err := client.People().Get("id")

// Get all the films.
response, err := client.Film().Getall("search")

// Get a specific film.
response, err := client.Film().Get("id")

// Get all the starships.
response, err := client.Starship().Getall("search")

// Get a specific starship.
response, err := client.Starship().Get("id")

// Get all the species.
response, err := client.Species().Getall("search")

// Get a specific species.
response, err := client.Species().Get("id")

// Get all the vehicles.
response, err := client.Vehicle().Getall("search")

// Get a specific vehicle.
response, err := client.Vehicle().Get("id")

// Get all the planets.
response, err := client.Planet().Getall("search")

// Get a specific planet.
response, err := client.Planet().Get("id")
```
