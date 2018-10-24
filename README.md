# adminarea

Get provinces, cities/districts, districts and villages from a country

## Installation

```cmd
go get github.com/ramadani/adminarea
```

## Usage

Available Country
- Indonesia (ID)

```go
import "github.com/ramadani/adminarea"

idAdminArea := adminarea.New("ID")

// get country
country := idAdminArea.GetCountry()

// get provinces
provinces, err := idAdminArea.GetProvinces()

// get regencies
regencies, err := idAdminArea.GetRegencies()

// get districts
districts, err := idAdminArea.GetDistricts()

// get villages
villages, err := idAdminArea.GetVillages()
```

## References

* [edwardsamuel/Wilayah-Administratif-Indonesia](https://github.com/edwardsamuel/Wilayah-Administratif-Indonesia)
* [Country Codes List](https://www.nationsonline.org/oneworld/country_code_list.htm)
* [jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata)