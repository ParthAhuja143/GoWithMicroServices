package utils

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/ParthAhuja143/GoWithMicroServices/errors"
)

func GetIDFromURI(rw http.ResponseWriter, r *http.Request) (int, errors.HTTPError) {
	regex := regexp.MustCompile(`/([0-9]+)`)
	path := r.URL.Path

	uidArr := regex.FindAllStringSubmatch(path, -1)
	if len(uidArr) != 1 {
		return -1, errors.ErrInvalidURI
	}

	uid := uidArr[0][1]
	id, err := strconv.Atoi(uid)

	if err != nil {
		return -1, errors.ErrInvalidURI
	}

	return id, errors.NoErr
}

