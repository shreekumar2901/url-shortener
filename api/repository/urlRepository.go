package repository

import (
	"errors"
	"time"

	"github.com/shreekumar2901/url-shortener/database"
	"github.com/shreekumar2901/url-shortener/domain"
	"github.com/shreekumar2901/url-shortener/dto"
	"gorm.io/gorm"
)

func DeleteShortByUrl(url string, userId string) error {
	db := database.Db.DB
	var record domain.Urls

	err := db.Where("url = ? AND user_id = ?", url, userId).First(&record).Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return errors.New("record not found")
		}

		return err
	}

	if err := db.Delete(&record).Error; err != nil {
		return err
	}
	return nil
}

func GetAll(userId string) ([]domain.Urls, error) {
	db := database.Db.DB

	var urls []domain.Urls

	if err := db.Where("user_id = ?", userId).Find(&urls).Error; err != nil {
		return urls, errors.New("some error occured when getting urls")
	}

	return urls, nil
}

func GetUrlByShort(short string, userId string) (domain.Urls, error) {
	db := database.Db.DB
	var url domain.Urls

	err := db.Where("short = ? AND user_id = ?", short, userId).First(&url).Error

	if err != nil {
		return domain.Urls{}, errors.New("url for this short does not exist")
	}

	return url, nil
}

func SaveShortenedUrl(urlDTO dto.UrlShortenRequestDTO, userId string) (domain.Urls, error) {
	db := database.Db.DB
	url := domain.Urls{
		Short:  urlDTO.CustomShort,
		Url:    urlDTO.Url,
		Expiry: time.Now().Local().Add(48 * time.Hour), // Setting expity as 48 hours
		UserID: userId,
	}

	if err := db.Save(&url).Error; err != nil {
		return url, errors.New("some error occurred! please try again")
	}

	return url, nil
}

func GetShortByUrl(url string, userId string) (string, error) {
	db := database.Db.DB

	var urlDomain domain.Urls

	err := db.Where("url = ? AND user_id = ?", url, userId).First(&urlDomain).Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return "", errors.New("short does not exist for given url")
		}

		return "", err
	}

	return urlDomain.Short, nil
}
