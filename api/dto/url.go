package dto

type UrlShortenRequestDTO struct {
	Url         string `json:"url"`
	CustomShort string `json:"custom_short"`
}

type UrlShortenResponseDTO struct {
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}
