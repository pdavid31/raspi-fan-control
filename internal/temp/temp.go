package temp

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	temperatureFile = "/sys/class/thermal/thermal_zone0/temp"
)

func Get() (float64, error) {
	content, err := ioutil.ReadFile(temperatureFile)
	if err != nil {
		return 0, err
	}

	str := string(content)
	temperature, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return 0, err
	}

	return temperature / 1000, nil
}
