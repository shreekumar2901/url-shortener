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

func (s *UrlService) ShortenUrl(requestDTO dto.UrlShortenRequestDTO) (dto.UrlShortenResponseDTO, error) {

	errorMsgs := validator.ShortenUrlValidator(&requestDTO)
	var response dto.UrlShortenResponseDTO

	if len(errorMsgs["errors"]) > 0 {
		return response, errors.New(strings.Join(errorMsgs["errors"], ","))
	}

	isUrlExist := s.isShortExist(requestDTO.CustomShort)

	if isUrlExist {
		return response, errors.New("can not user given custom short")
	}

	requestDTO.Url = helpers.EnforeHTTP(requestDTO.Url)

	customShort := requestDTO.CustomShort

	if len(strings.TrimSpace(customShort)) == 0 {
		customShort = uuid.NewString()[:6]
	}

	requestDTO.CustomShort = customShort

	url, err := repository.SaveShortenedUrl(requestDTO)

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
