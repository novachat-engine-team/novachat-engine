
https://stackoverflow.com/questions/42876418/geolite2-country-list-of-codes

struct CountryIP {
    Ip string
    Country String
}

mtproto.NearestDc:country

countryIP := CountryIP
nearestDc := &mtproto.NearestDc{}
nearestDc.Country = countryIP.Country

