package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.themoviedb.org/3"

type Client struct {
    apiKey string
}

func NewClient(apiKey string) *Client {
    return &Client{apiKey: apiKey}
}

func (c *Client) FetchMovieGenres() ([]Genres, error) {
    url := fmt.Sprintf("%s/genre/movie/list?api_key=%s",apiURL,c.apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err

    }
    defer resp.Body.Close()
    var genres struct {
        Genres []Genres `json:"genres"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&genres); err != nil {
        return nil, err
    }
    return genres.Genres, nil


}

func (c *Client) FetchTrendingMovies() ([]Movie, error) {
    url := fmt.Sprintf("%s/trending/movie/day?api_key=%s",apiURL,c.apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Results []Movie `json:"results"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return result.Results, nil
}

func (c *Client) FetchPopularMovies() ([]Movie, error) {
    url := fmt.Sprintf("%s/movie/popular?api_key=%s", apiURL, c.apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Results []Movie `json:"results"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return result.Results, nil
}

func (c *Client) FetchMovieDetails(id int) (*Movie, error) {
    url := fmt.Sprintf("%s/movie/%d?api_key=%s", apiURL, id, c.apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var movie Movie
    if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
        return nil, err
    }
    return &movie, nil
}

func (c *Client) FetchMovieCast(id int) (*Credits, error) {
    url := fmt.Sprintf("%s/movie/%d/credits?api_key=%s", apiURL, id, c.apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var credit Credits
    if err := json.NewDecoder(resp.Body).Decode(&credit); err != nil {
        return nil, err
    }
    return &credit, nil
}

func (c *Client) FetchTrendingTv() ([]TV , error) {
    url := fmt.Sprintf("%s/trending/tv/day?api_key=%s",apiURL,c.apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Results []TV `json:"results"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return result.Results, nil
}