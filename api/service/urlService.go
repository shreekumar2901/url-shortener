package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/shreekumar2901/url-shortener/config"
	"github.com/shreekumar2901/url-shortener/dto"
	"github.com/shreekumar2901/url-shortener/helpers"
	"github.com/shreekumar2901/url-shortener/repository"
	"github.com/shreekumar2901/url-shortener/validator"
)

type UrlService struct{}

func (s *UrlService) isShortExist(short string) bool {

	_, err := repository.GetUrlByShort(short)

	return err == nil

}

func (s *UrlService) ShortenUrl(shortenUrlRequestDTO dto.UrlShortenRequestDTO) (dto.UrlShortenResponseDTO, error) {

	errorMsgs := validator.ShortenUrlValidator(&shortenUrlRequestDTO)
	var response dto.UrlShortenResponseDTO

	if len(errorMsgs["errors"]) > 0 {
		return response, errors.New(strings.Join(errorMsgs["errors"], ","))
	}

	isUrlExist := s.isShortExist(shortenUrlRequestDTO.CustomShort)

	if isUrlExist {
		return response, errors.New("can not user given custom short")
	}

	shortenUrlRequestDTO.Url = helpers.EnforeHTTP(shortenUrlRequestDTO.Url)

	customShort := shortenUrlRequestDTO.CustomShort

	if len(strings.TrimSpace(customShort)) == 0 {
		customShort = uuid.NewString()[:6]
	}

	shortenUrlRequestDTO.CustomShort = customShort

	url, err := repository.SaveShortenedUrl(shortenUrlRequestDTO)

	if err != nil {
		return response, err
	}

	response = dto.UrlShortenResponseDTO{
		Url:      url.Url,
		ShortUrl: config.Config("DOMAIN") + "/" + url.Short,
	}

	return response, nil

}

func (s *UrlService) ResolveUrl(short string) (string, error) {
	url, err := repository.GetUrlByShort(short)

	if err != nil {
		return "", err
	}

	return url.Url, nil
}

func (s *UrlService) ListUrls() ([]dto.UrlListResponseDTO, error) {
	urls, err := repository.GetAll()

	if err != nil {
		return []dto.UrlListResponseDTO{}, err
	}

	var response []dto.UrlListResponseDTO

	for _, url := range urls {
		element := dto.UrlListResponseDTO{
			Url:      url.Url,
			ShortUrl: config.Config("DOMAIN") + "/" + url.Short,
		}

		response = append(response, element)
	}

	return response, nil
}

func (s *UrlService) DeleteShortByUrl(url string) error {
	if err := repository.DeleteShortByUrl(url); err != nil {
		return err
	}

	return nil
}
