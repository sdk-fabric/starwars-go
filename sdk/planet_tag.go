
// PlanetTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type PlanetTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll 
func (client *PlanetTag) GetAll(name string) (PlanetCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["name"] = name

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/planets", pathParams))
    if err != nil {
        return PlanetCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return PlanetCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return PlanetCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return PlanetCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response PlanetCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return PlanetCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return PlanetCollection{}, errors.New("the server returned an unknown status code")
    }
}

// Get 
func (client *PlanetTag) Get(id string) (Planet, error) {
    pathParams := make(map[string]interface{})
    pathParams["id"] = id

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/planets/:id", pathParams))
    if err != nil {
        return Planet{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return Planet{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Planet{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Planet{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Planet
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Planet{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return Planet{}, errors.New("the server returned an unknown status code")
    }
}



func NewPlanetTag(httpClient *http.Client, parser *sdkgen.Parser) *PlanetTag {
	return &PlanetTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
