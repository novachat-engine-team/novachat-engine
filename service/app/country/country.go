package country

import (
	"bufio"
	"bytes"
	"novachat_engine/mtproto"
	"os"
	"strings"
)

var countryList []*mtproto.Help_Country

func InstallCountryList(filename string) {
	err := parseCountry(filename)
	if err != nil {
		panic(err)
	}
}

func GetCountryList() *mtproto.Help_CountriesList {
	return mtproto.NewTLHelpCountriesList(&mtproto.Help_CountriesList{
		Countries: countryList,
		Hash:      0,
	}).To_Help_CountriesList()
}

type raw struct {
	CountryCode string
	Patterns    string
	ISO         string
	DefaultName string
	Name        string
	Len         string
}

func makeRaw(l []string) (raw, bool) {
	r := raw{}
	if len(l) == 0 {
		return r, false
	}
	switch len(l) {
	case 6:
		fallthrough
	case 5:
		r.Len = l[4]
		fallthrough
	case 4:
		r.Patterns = l[3]
		fallthrough
	case 3:
		r.CountryCode = l[0]
		r.ISO = l[1]
		r.DefaultName = l[2]
		r.Name = l[2]
	default:
		panic(len(l))
	}
	return r, true
}

func parseCountry(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(bytes.NewReader(data))

	countryList = make([]*mtproto.Help_Country, 0, 32)
	for {
		lineData, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		_ = prefix
		line := string(lineData)
		l := strings.Split(line, ";")
		value, ok := makeRaw(l)
		if !ok {
			continue
		}

		var codeList []*mtproto.Help_CountryCode
		code := mtproto.Help_CountryCode{
			CountryCode: value.CountryCode,
			Prefixes:    nil,
			Patterns:    nil,
		}

		codeList = []*mtproto.Help_CountryCode{mtproto.NewTLHelpCountryCode(&code).To_Help_CountryCode()}
		if len(value.Patterns) > 0 {

			ll := strings.Split(value.Patterns, " ")
			code.Patterns = make([]string, 0, len(ll))
			for _, v := range ll {
				if v == value.CountryCode {
					code.Prefixes = []string{value.CountryCode}
					continue
				}

				code.Patterns = append(code.Patterns, v)
			}
		} else {
			code.Prefixes = []string{value.CountryCode}
		}

		country := mtproto.Help_Country{
			Hidden:       false,
			Iso2:         value.ISO,
			DefaultName:  value.DefaultName,
			Name:         value.Name,
			CountryCodes: codeList,
		}
		countryList = append(countryList, mtproto.NewTLHelpCountry(&country).To_Help_Country())
	}

	return nil
}
