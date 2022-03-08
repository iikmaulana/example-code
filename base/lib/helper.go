package lib

import (
	"github.com/iikmaulana/example-code/base/models"
	"regexp"
)

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func PaginateFromQuery(totalData int64, page int, limit int) (res models.FormMetaData, offset int) {

	offset = (page * limit) - limit
	cLimit := int64(limit)

	lastPage := totalData / cLimit
	if totalData%cLimit != 0 {
		lastPage += 1
	}

	res.TotalDataPerpage = 1
	res.TotalPage = int(lastPage)
	res.TotalData = totalData
	res.To = page + 1
	res.From = page - 1

	if res.From < res.TotalDataPerpage {
		res.From = 1
	}

	if res.To > res.TotalPage {
		res.To = res.TotalPage
	}

	return res, offset
}

func Paginate(totalData int64, page int, limit int) (res models.FormMetaData) {
	cLimit := int64(limit)

	lastPage := totalData / cLimit
	if totalData%cLimit != 0 {
		lastPage += 1
	}

	res.TotalDataPerpage = 1
	res.TotalPage = int(lastPage)
	res.TotalData = totalData
	res.To = page + 1
	res.From = page - 1

	if res.From < res.TotalDataPerpage {
		res.From = 1
	}

	if res.To > res.TotalPage {
		res.To = res.TotalPage
	}

	return res
}
