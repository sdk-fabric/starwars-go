
// PeopleTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type PeopleTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll Get all the people
func (client *PeopleTag) GetAll(search string) (PeopleCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/people", pathParams))
    if err != nil {
        return PeopleCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return PeopleCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return PeopleCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return PeopleCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response PeopleCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return PeopleCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return PeopleCollection{}, errors.New("the server returned an unknown status code")
    }
}

// Get Get a specific people
func (client *PeopleTag) Get(id string) (People, error) {
    pathParams := make(map[string]interface{})
    pathParams["id"] = id

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/people/:id", pathParams))
    if err != nil {
        return People{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return People{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return People{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return People{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response People
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return People{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return People{}, errors.New("the server returned an unknown status code")
    }
}



func NewPeopleTag(httpClient *http.Client, parser *sdkgen.Parser) *PeopleTag {
	return &PeopleTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
