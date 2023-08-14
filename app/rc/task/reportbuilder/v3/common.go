package v3

import (
	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"time"
)

func GetFinFactorTimeline(content []byte) (time.Time, error) {
	tStr, err := jsonparser.GetString(content, "impExpEntReport", "lrbDetail", "[0]", "RQ")
	if err != nil {
		return time.Time{}, errors.Wrap(err, "get RQ error while iter lrbDetail")
	}
	t, err := time.Parse("2006-01-02", tStr)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "parse time error while iter lrbDetail")
	}
	return t, nil
}
