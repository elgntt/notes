package dto

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const CustomDate = "02-01-2006" // число-месяц-год

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format(CustomDate))), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(CustomDate, strings.Trim(string(data), "\""))
	if err != nil {
		return errors.New(`невалидный формат времени, ожидается вид: число-месяц-год`)
	}

	d.Time = t

	return nil
}
