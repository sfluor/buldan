package main

import (
	"fmt"
	"maps"

	mrand "math/rand"
)

const minScore = 0.85
const minScoreExact = 0.95

type rawCountry struct {
	name string
	flag string
}

type Country struct {
	Name           string
	NormalizedName string
	Flag           string
}

type Countries struct {
	byName map[string]Country
}

func countriesStartingWith(char byte) Countries {
	byName := countries[char]

	return Countries{byName: maps.Clone(byName)}
}

func (c *Countries) remaining() int {
	return len(c.byName)
}

func (c *Countries) guess(guess string) (Country, string, bool) {
	normalizedQuery := normalize(guess)

	// countryScore holds a country name and its similarity score
	type countryScore struct {
		key     string
		country Country
		score   float64
	}

	match := countryScore{}

	for name, country := range c.byName {
		// Technically no need to renormalize
		normalizedCountry := normalize(name)
		score := jaroWinkler(normalizedQuery, normalizedCountry)
		if match.score < score {
			match = countryScore{key: name, country: country, score: score}
		}
	}

	valid := match.score >= minScore
	guessStr := guess
	if valid {
		delete(c.byName, match.key)

		if match.score < minScoreExact {
			guessStr = fmt.Sprintf("%s (%s)", guess, match.country.Name)
		}
	}

	return match.country, guessStr, valid
}

// Letter -> Normalized Name -> Country
var countries = map[byte]map[string]Country{}

func newLetters() []byte {
	letters := []byte{}
	for l := range countries {
		letters = append(letters, l)
	}

	mrand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	// TODO: remove me once we finished testing
	letters[0] = 'y'
	return letters
}

func init() {
	for _, country := range rawCountries {
		normalized := normalize(country.name)
		firstLetter := normalized[0]

		if _, ok := countries[firstLetter]; !ok {
			countries[firstLetter] = map[string]Country{}
		}

		if prev, ok := countries[firstLetter][normalized]; ok {
			panic(fmt.Sprintf("Found duplicate countries: %+v and %+v", prev, country))
		}

		countries[firstLetter][normalized] = Country{
			Name:           country.name,
			Flag:           country.flag,
			NormalizedName: normalized,
		}
	}
}

// Dumped from: https://flagpedia.net/emoji#google_vignette
// TODO: special chars
var rawCountries = []rawCountry{
	{
		name: "Afghanistan",
		flag: "🇦🇫",
	},
	{
		name: "Åland Islands",
		flag: "🇦🇽",
	},
	{
		name: "Albania",
		flag: "🇦🇱",
	},
	{
		name: "Algeria",
		flag: "🇩🇿",
	},
	{
		name: "American Samoa",
		flag: "🇦🇸",
	},
	{
		name: "Andorra",
		flag: "🇦🇩",
	},
	{
		name: "Angola",
		flag: "🇦🇴",
	},
	{
		name: "Anguilla",
		flag: "🇦🇮",
	},
	{
		name: "Antarctica",
		flag: "🇦🇶",
	},
	{
		name: "Antigua and Barbuda",
		flag: "🇦🇬",
	},
	{
		name: "Argentina",
		flag: "🇦🇷",
	},
	{
		name: "Armenia",
		flag: "🇦🇲",
	},
	{
		name: "Aruba",
		flag: "🇦🇼",
	},
	{
		name: "Australia",
		flag: "🇦🇺",
	},
	{
		name: "Austria",
		flag: "🇦🇹",
	},
	{
		name: "Azerbaijan",
		flag: "🇦🇿",
	},
	{
		name: "Bahamas",
		flag: "🇧🇸",
	},
	{
		name: "Bahrain",
		flag: "🇧🇭",
	},
	{
		name: "Bangladesh",
		flag: "🇧🇩",
	},
	{
		name: "Barbados",
		flag: "🇧🇧",
	},
	{
		name: "Belarus",
		flag: "🇧🇾",
	},
	{
		name: "Belgium",
		flag: "🇧🇪",
	},
	{
		name: "Belize",
		flag: "🇧🇿",
	},
	{
		name: "Benin",
		flag: "🇧🇯",
	},
	{
		name: "Bermuda",
		flag: "🇧🇲",
	},
	{
		name: "Bhutan",
		flag: "🇧🇹",
	},
	{
		name: "Bolivia",
		flag: "🇧🇴",
	},
	{
		name: "Bosnia and Herzegovina",
		flag: "🇧🇦",
	},
	{
		name: "Botswana",
		flag: "🇧🇼",
	},
	{
		name: "Bouvet Island",
		flag: "🇧🇻",
	},
	{
		name: "Brazil",
		flag: "🇧🇷",
	},
	{
		name: "British Indian Ocean Territory",
		flag: "🇮🇴",
	},
	{
		name: "Brunei",
		flag: "🇧🇳",
	},
	{
		name: "Bulgaria",
		flag: "🇧🇬",
	},
	{
		name: "Burkina Faso",
		flag: "🇧🇫",
	},
	{
		name: "Burundi",
		flag: "🇧🇮",
	},
	{
		name: "Cambodia",
		flag: "🇰🇭",
	},
	{
		name: "Cameroon",
		flag: "🇨🇲",
	},
	{
		name: "Canada",
		flag: "🇨🇦",
	},
	{
		name: "Cape Verde",
		flag: "🇨🇻",
	},
	{
		name: "Caribbean Netherlands",
		flag: "🇧🇶",
	},
	{
		name: "Cayman Islands",
		flag: "🇰🇾",
	},
	{
		name: "Central African Republic",
		flag: "🇨🇫",
	},
	{
		name: "Chad",
		flag: "🇹🇩",
	},
	{
		name: "Chile",
		flag: "🇨🇱",
	},
	{
		name: "China",
		flag: "🇨🇳",
	},
	{
		name: "Christmas Island",
		flag: "🇨🇽",
	},
	{
		name: "Cocos Islands",
		flag: "🇨🇨",
	},
	{
		name: "Colombia",
		flag: "🇨🇴",
	},
	{
		name: "Comoros",
		flag: "🇰🇲",
	},
	{
		name: "Republic of the Congo",
		flag: "🇨🇬",
	},
	{
		name: "DR Congo",
		flag: "🇨🇩",
	},
	{
		name: "Cook Islands",
		flag: "🇨🇰",
	},
	{
		name: "Costa Rica",
		flag: "🇨🇷",
	},
	{
		name: "Côte d'Ivoire",
		flag: "🇨🇮",
	},
	{
		name: "Croatia",
		flag: "🇭🇷",
	},
	{
		name: "Cuba",
		flag: "🇨🇺",
	},
	{
		name: "Curaçao",
		flag: "🇨🇼",
	},
	{
		name: "Cyprus",
		flag: "🇨🇾",
	},
	{
		name: "Czechia",
		flag: "🇨🇿",
	},
	{
		name: "Denmark",
		flag: "🇩🇰",
	},
	{
		name: "Djibouti",
		flag: "🇩🇯",
	},
	{
		name: "Dominica",
		flag: "🇩🇲",
	},
	{
		name: "Dominican Republic",
		flag: "🇩🇴",
	},
	{
		name: "Ecuador",
		flag: "🇪🇨",
	},
	{
		name: "Egypt",
		flag: "🇪🇬",
	},
	{
		name: "El Salvador",
		flag: "🇸🇻",
	},
	{
		name: "England",
		flag: "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	},
	{
		name: "Equatorial Guinea",
		flag: "🇬🇶",
	},
	{
		name: "Eritrea",
		flag: "🇪🇷",
	},
	{
		name: "Estonia",
		flag: "🇪🇪",
	},
	{
		name: "Eswatini",
		flag: "🇸🇿",
	},
	{
		name: "Ethiopia",
		flag: "🇪🇹",
	},
	{
		name: "Falkland Islands",
		flag: "🇫🇰",
	},
	{
		name: "Faroe Islands",
		flag: "🇫🇴",
	},
	{
		name: "Fiji",
		flag: "🇫🇯",
	},
	{
		name: "Finland",
		flag: "🇫🇮",
	},
	{
		name: "France",
		flag: "🇫🇷",
	},
	{
		name: "French Guiana",
		flag: "🇬🇫",
	},
	{
		name: "French Polynesia",
		flag: "🇵🇫",
	},
	{
		name: "French Southern and Antarctic Lands",
		flag: "🇹🇫",
	},
	{
		name: "Gabon",
		flag: "🇬🇦",
	},
	{
		name: "Gambia",
		flag: "🇬🇲",
	},
	{
		name: "Georgia",
		flag: "🇬🇪",
	},
	{
		name: "Germany",
		flag: "🇩🇪",
	},
	{
		name: "Ghana",
		flag: "🇬🇭",
	},
	{
		name: "Gibraltar",
		flag: "🇬🇮",
	},
	{
		name: "Greece",
		flag: "🇬🇷",
	},
	{
		name: "Greenland",
		flag: "🇬🇱",
	},
	{
		name: "Grenada",
		flag: "🇬🇩",
	},
	{
		name: "Guadeloupe",
		flag: "🇬🇵",
	},
	{
		name: "Guam",
		flag: "🇬🇺",
	},
	{
		name: "Guatemala",
		flag: "🇬🇹",
	},
	{
		name: "Guernsey",
		flag: "🇬🇬",
	},
	{
		name: "Guinea",
		flag: "🇬🇳",
	},
	{
		name: "Guinea-Bissau",
		flag: "🇬🇼",
	},
	{
		name: "Guyana",
		flag: "🇬🇾",
	},
	{
		name: "Haiti",
		flag: "🇭🇹",
	},
	{
		name: "Heard Island and McDonald Islands",
		flag: "🇭🇲",
	},
	{
		name: "Honduras",
		flag: "🇭🇳",
	},
	{
		name: "Hong Kong",
		flag: "🇭🇰",
	},
	{
		name: "Hungary",
		flag: "🇭🇺",
	},
	{
		name: "Iceland",
		flag: "🇮🇸",
	},
	{
		name: "India",
		flag: "🇮🇳",
	},
	{
		name: "Indonesia",
		flag: "🇮🇩",
	},
	{
		name: "Iran",
		flag: "🇮🇷",
	},
	{
		name: "Iraq",
		flag: "🇮🇶",
	},
	{
		name: "Ireland",
		flag: "🇮🇪",
	},
	{
		name: "Isle of Man",
		flag: "🇮🇲",
	},
	{
		name: "Italy",
		flag: "🇮🇹",
	},
	{
		name: "Jamaica",
		flag: "🇯🇲",
	},
	{
		name: "Japan",
		flag: "🇯🇵",
	},
	{
		name: "Jersey",
		flag: "🇯🇪",
	},
	{
		name: "Jordan",
		flag: "🇯🇴",
	},
	{
		name: "Kazakhstan",
		flag: "🇰🇿",
	},
	{
		name: "Kenya",
		flag: "🇰🇪",
	},
	{
		name: "Kiribati",
		flag: "🇰🇮",
	},
	{
		name: "North Korea",
		flag: "🇰🇵",
	},
	{
		name: "South Korea",
		flag: "🇰🇷",
	},
	{
		name: "Kosovo",
		flag: "🇽🇰",
	},
	{
		name: "Kuwait",
		flag: "🇰🇼",
	},
	{
		name: "Kyrgyzstan",
		flag: "🇰🇬",
	},
	{
		name: "Laos",
		flag: "🇱🇦",
	},
	{
		name: "Latvia",
		flag: "🇱🇻",
	},
	{
		name: "Lebanon",
		flag: "🇱🇧",
	},
	{
		name: "Lesotho",
		flag: "🇱🇸",
	},
	{
		name: "Liberia",
		flag: "🇱🇷",
	},
	{
		name: "Libya",
		flag: "🇱🇾",
	},
	{
		name: "Liechtenstein",
		flag: "🇱🇮",
	},
	{
		name: "Lithuania",
		flag: "🇱🇹",
	},
	{
		name: "Luxembourg",
		flag: "🇱🇺",
	},
	{
		name: "Macau",
		flag: "🇲🇴",
	},
	{
		name: "Madagascar",
		flag: "🇲🇬",
	},
	{
		name: "Malawi",
		flag: "🇲🇼",
	},
	{
		name: "Malaysia",
		flag: "🇲🇾",
	},
	{
		name: "Maldives",
		flag: "🇲🇻",
	},
	{
		name: "Mali",
		flag: "🇲🇱",
	},
	{
		name: "Malta",
		flag: "🇲🇹",
	},
	{
		name: "Marshall Islands",
		flag: "🇲🇭",
	},
	{
		name: "Martinique",
		flag: "🇲🇶",
	},
	{
		name: "Mauritania",
		flag: "🇲🇷",
	},
	{
		name: "Mauritius",
		flag: "🇲🇺",
	},
	{
		name: "Mayotte",
		flag: "🇾🇹",
	},
	{
		name: "Mexico",
		flag: "🇲🇽",
	},
	{
		name: "Micronesia",
		flag: "🇫🇲",
	},
	{
		name: "Moldova",
		flag: "🇲🇩",
	},
	{
		name: "Monaco",
		flag: "🇲🇨",
	},
	{
		name: "Mongolia",
		flag: "🇲🇳",
	},
	{
		name: "Montenegro",
		flag: "🇲🇪",
	},
	{
		name: "Montserrat",
		flag: "🇲🇸",
	},
	{
		name: "Morocco",
		flag: "🇲🇦",
	},
	{
		name: "Mozambique",
		flag: "🇲🇿",
	},
	{
		name: "Myanmar",
		flag: "🇲🇲",
	},
	{
		name: "Namibia",
		flag: "🇳🇦",
	},
	{
		name: "Nauru",
		flag: "🇳🇷",
	},
	{
		name: "Nepal",
		flag: "🇳🇵",
	},
	{
		name: "Netherlands",
		flag: "🇳🇱",
	},
	{
		name: "New Caledonia",
		flag: "🇳🇨",
	},
	{
		name: "New Zealand",
		flag: "🇳🇿",
	},
	{
		name: "Nicaragua",
		flag: "🇳🇮",
	},
	{
		name: "Niger",
		flag: "🇳🇪",
	},
	{
		name: "Nigeria",
		flag: "🇳🇬",
	},
	{
		name: "Niue",
		flag: "🇳🇺",
	},
	{
		name: "Norfolk Island",
		flag: "🇳🇫",
	},
	{
		name: "North Macedonia",
		flag: "🇲🇰",
	},
	{
		name: "Northern Mariana Islands",
		flag: "🇲🇵",
	},
	{
		name: "Norway",
		flag: "🇳🇴",
	},
	{
		name: "Oman",
		flag: "🇴🇲",
	},
	{
		name: "Pakistan",
		flag: "🇵🇰",
	},
	{
		name: "Palau",
		flag: "🇵🇼",
	},
	{
		name: "Palestine",
		flag: "🇵🇸",
	},
	{
		name: "Panama",
		flag: "🇵🇦",
	},
	{
		name: "Papua New Guinea",
		flag: "🇵🇬",
	},
	{
		name: "Paraguay",
		flag: "🇵🇾",
	},
	{
		name: "Peru",
		flag: "🇵🇪",
	},
	{
		name: "Philippines",
		flag: "🇵🇭",
	},
	{
		name: "Pitcairn Islands",
		flag: "🇵🇳",
	},
	{
		name: "Poland",
		flag: "🇵🇱",
	},
	{
		name: "Portugal",
		flag: "🇵🇹",
	},
	{
		name: "Puerto Rico",
		flag: "🇵🇷",
	},
	{
		name: "Qatar",
		flag: "🇶🇦",
	},
	{
		name: "Réunion",
		flag: "🇷🇪",
	},
	{
		name: "Romania",
		flag: "🇷🇴",
	},
	{
		name: "Russia",
		flag: "🇷🇺",
	},
	{
		name: "Rwanda",
		flag: "🇷🇼",
	},
	{
		name: "Saint Barthélemy",
		flag: "🇧🇱",
	},
	{
		name: "Saint Helena, Ascension and Tristan da Cunha",
		flag: "🇸🇭",
	},
	{
		name: "Saint Kitts and Nevis",
		flag: "🇰🇳",
	},
	{
		name: "Saint Lucia",
		flag: "🇱🇨",
	},
	{
		name: "Saint Martin",
		flag: "🇲🇫",
	},
	{
		name: "Saint Pierre and Miquelon",
		flag: "🇵🇲",
	},
	{
		name: "Saint Vincent and the Grenadines",
		flag: "🇻🇨",
	},
	{
		name: "Samoa",
		flag: "🇼🇸",
	},
	{
		name: "San Marino",
		flag: "🇸🇲",
	},
	{
		name: "São Tomé and Príncipe",
		flag: "🇸🇹",
	},
	{
		name: "Saudi Arabia",
		flag: "🇸🇦",
	},
	{
		name: "Scotland",
		flag: "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
	},
	{
		name: "Senegal",
		flag: "🇸🇳",
	},
	{
		name: "Serbia",
		flag: "🇷🇸",
	},
	{
		name: "Seychelles",
		flag: "🇸🇨",
	},
	{
		name: "Sierra Leone",
		flag: "🇸🇱",
	},
	{
		name: "Singapore",
		flag: "🇸🇬",
	},
	{
		name: "Sint Maarten",
		flag: "🇸🇽",
	},
	{
		name: "Slovakia",
		flag: "🇸🇰",
	},
	{
		name: "Slovenia",
		flag: "🇸🇮",
	},
	{
		name: "Solomon Islands",
		flag: "🇸🇧",
	},
	{
		name: "Somalia",
		flag: "🇸🇴",
	},
	{
		name: "South Africa",
		flag: "🇿🇦",
	},
	{
		name: "South Georgia",
		flag: "🇬🇸",
	},
	{
		name: "South Sudan",
		flag: "🇸🇸",
	},
	{
		name: "Spain",
		flag: "🇪🇸",
	},
	{
		name: "Sri Lanka",
		flag: "🇱🇰",
	},
	{
		name: "Sudan",
		flag: "🇸🇩",
	},
	{
		name: "Suriname",
		flag: "🇸🇷",
	},
	{
		name: "Svalbard and Jan Mayen",
		flag: "🇸🇯",
	},
	{
		name: "Sweden",
		flag: "🇸🇪",
	},
	{
		name: "Switzerland",
		flag: "🇨🇭",
	},
	{
		name: "Syria",
		flag: "🇸🇾",
	},
	{
		name: "Taiwan",
		flag: "🇹🇼",
	},
	{
		name: "Tajikistan",
		flag: "🇹🇯",
	},
	{
		name: "Tanzania",
		flag: "🇹🇿",
	},
	{
		name: "Thailand",
		flag: "🇹🇭",
	},
	{
		name: "Timor-Leste",
		flag: "🇹🇱",
	},
	{
		name: "Togo",
		flag: "🇹🇬",
	},
	{
		name: "Tokelau",
		flag: "🇹🇰",
	},
	{
		name: "Tonga",
		flag: "🇹🇴",
	},
	{
		name: "Trinidad and Tobago",
		flag: "🇹🇹",
	},
	{
		name: "Tunisia",
		flag: "🇹🇳",
	},
	{
		name: "Turkey",
		flag: "🇹🇷",
	},
	{
		name: "Turkmenistan",
		flag: "🇹🇲",
	},
	{
		name: "Turks and Caicos Islands",
		flag: "🇹🇨",
	},
	{
		name: "Tuvalu",
		flag: "🇹🇻",
	},
	{
		name: "Uganda",
		flag: "🇺🇬",
	},
	{
		name: "Ukraine",
		flag: "🇺🇦",
	},
	{
		name: "United Arab Emirates",
		flag: "🇦🇪",
	},
	{
		name: "United Kingdom",
		flag: "🇬🇧",
	},
	{
		name: "United States",
		flag: "🇺🇸",
	},
	{
		name: "United States Minor Outlying Islands",
		flag: "🇺🇲",
	},
	{
		name: "Uruguay",
		flag: "🇺🇾",
	},
	{
		name: "Uzbekistan",
		flag: "🇺🇿",
	},
	{
		name: "Vanuatu",
		flag: "🇻🇺",
	},
	{
		name: "Vatican",
		flag: "🇻🇦",
	},
	{
		name: "Venezuela",
		flag: "🇻🇪",
	},
	{
		name: "Vietnam",
		flag: "🇻🇳",
	},
	{
		name: "British Virgin Islands",
		flag: "🇻🇬",
	},
	{
		name: "United States Virgin Islands",
		flag: "🇻🇮",
	},
	{
		name: "Wales",
		flag: "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
	},
	{
		name: "Wallis and Futuna",
		flag: "🇼🇫",
	},
	{
		name: "Western Sahara",
		flag: "🇪🇭",
	},
	{
		name: "Yemen",
		flag: "🇾🇪",
	},
	{
		name: "Zambia",
		flag: "🇿🇲",
	},
	{
		name: "Zimbabwe",
		flag: "🇿🇼",
	},
}