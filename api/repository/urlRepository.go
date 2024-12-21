package repository

import (
	"errors"
	"time"

	"github.com/shreekumar2901/url-shortener/database"
	"github.com/shreekumar2901/url-shortener/domain"
	"github.com/shreekumar2901/url-shortener/dto"
)

func GetUrlByShort(short string) (domain.Urls, error) {
	db := database.Db.DB
	var url domain.Urls

	err := db.Where("short = ?", short).First(&url).Error

	if err != nil {
		return domain.Urls{}, errors.New("url for this short does not exist")
	}

	return url, nil
}

func SaveShortenedUrl(urlDTO dto.UrlShortenRequestDTO) (domain.Urls, error) {
	db := database.Db.DB
	url := domain.Urls{
		Short:  urlDTO.CustomShort,
		Url:    urlDTO.Url,
		Expiry: time.Now().Local().Add(48 * time.Hour), // Setting expity as 48 hours
	}

	if err := db.Save(&url).Error; err != nil {
		return url, errors.New("some error occurred! please try again")
	}

	return url, nil
}

func GetShortByUrl(url string) (string, error) {
	db := database.Db.DB

	var urlDomain domain.Urls

	err := db.Where("url = ?", url).First(&urlDomain).Error

	if err != nil {
		return "", errors.New("short does not exist for given url")
	}

	return urlDomain.Short, nil
}
