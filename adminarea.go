package adminarea

import (
	"log"

	idres "github.com/ramadani/adminarea/resources/indonesia"
	"github.com/ramadani/adminarea/src"
	"github.com/ramadani/adminarea/src/indonesia"
)

// AdminArea to get administrative areas of country
type AdminArea struct {
	province src.ProvinceContract
	regency  src.RegencyContract
	district src.DistrictContract
	village  src.VillageContract
}

// GetProvinces get provinces by country
func (aa *AdminArea) GetProvinces() ([]*src.Province, error) {
	return aa.province.All()
}

// GetRegencies get regencies by country
func (aa *AdminArea) GetRegencies() ([]*src.Regency, error) {
	return aa.regency.All()
}

// GetDistricts get districts by country
func (aa *AdminArea) GetDistricts() ([]*src.District, error) {
	return aa.district.All()
}

// GetVillages get villages by country
func (aa *AdminArea) GetVillages() ([]*src.Village, error) {
	return aa.village.All()
}

// New instance of AdminArea
func New(countryID string) *AdminArea {
	var province src.ProvinceContract
	var regency src.RegencyContract
	var district src.DistrictContract
	var village src.VillageContract

	countryRepo := src.NewCountryRepository()
	country, err := countryRepo.FindByID(countryID)
	if err != nil {
		log.Fatal(err)
	}

	switch country.ID {
	case "ID":
		prContents, err := idres.Asset("provinces.csv")
		if err != nil {
			log.Fatal(err)
		}

		rgContents, err := idres.Asset("regencies.csv")
		if err != nil {
			log.Fatal(err)
		}

		dsContents, err := idres.Asset("districts.csv")
		if err != nil {
			log.Fatal(err)
		}

		vlContents, err := idres.Asset("villages.csv")
		if err != nil {
			log.Fatal(err)
		}

		province = indonesia.NewProvince(country.ID, prContents)
		regency = indonesia.NewRegency(country.ID, rgContents)
		district = indonesia.NewDistrict(country.ID, dsContents)
		village = indonesia.NewVillage(country.ID, vlContents)
	}

	return &AdminArea{
		province,
		regency,
		district,
		village,
	}
}
