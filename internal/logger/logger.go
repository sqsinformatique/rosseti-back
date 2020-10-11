package logger

import (
	"reflect"
	"strconv"
	"strings"

	// other

	"github.com/rs/zerolog"
)

func InitializeLogger(parent *zerolog.Logger, emptyStruct interface{}) (log zerolog.Logger) {
	res := false
	packageTypes := []string{"domains", "internal"}
	packageType := ""
	pkgPath := reflect.TypeOf(emptyStruct).PkgPath()
	pathElements := strings.Split(pkgPath, "/")

	i := 0

	for _, pathElement := range pathElements {
		for _, packageType = range packageTypes {
			if pathElement == packageType {
				log = parent.With().Str("type", pathElements[i]).Logger()
				log = log.With().Str("package", pathElements[i+1]).Logger()
				res = true

				break
			}
		}

		if res {
			break
		}
		i++
	}

	if packageType == "domains" {
		ver, _ := strconv.ParseInt(strings.TrimPrefix(pathElements[i+2], "v"), 10, 64)
		log = log.With().Int64("version", ver).Logger()
		i++
	}

	if i+3 <= len(pathElements) {
		log = log.With().Str("subsystem", pathElements[i+2]).Logger()
	}

	log.Info().Msg("Initializing...")

	return log
}
