package main

import (
	"strings"
)

type Country struct {
	Name string
	Flag string
}

func countriesStartingWith(char byte) (map[string]Country, error) {
    // TODO normalize country names
	res := make(map[string]Country)

	for _, country := range countries {
        countryName := strings.TrimSpace(strings.ToLower(country.Name))
		if countryName[0] == char {
			res[countryName] = country
		}
	}

	return res, nil
}

// Dumped from: https://flagpedia.net/emoji#google_vignette
// TODO: special chars
var countries = []Country{
	{
		Name: "Afghanistan",
		Flag: "🇦🇫",
	},
	{
		Name: "Åland Islands",
		Flag: "🇦🇽",
	},
	{
		Name: "Albania",
		Flag: "🇦🇱",
	},
	{
		Name: "Algeria",
		Flag: "🇩🇿",
	},
	{
		Name: "American Samoa",
		Flag: "🇦🇸",
	},
	{
		Name: "Andorra",
		Flag: "🇦🇩",
	},
	{
		Name: "Angola",
		Flag: "🇦🇴",
	},
	{
		Name: "Anguilla",
		Flag: "🇦🇮",
	},
	{
		Name: "Antarctica",
		Flag: "🇦🇶",
	},
	{
		Name: "Antigua and Barbuda",
		Flag: "🇦🇬",
	},
	{
		Name: "Argentina",
		Flag: "🇦🇷",
	},
	{
		Name: "Armenia",
		Flag: "🇦🇲",
	},
	{
		Name: "Aruba",
		Flag: "🇦🇼",
	},
	{
		Name: "Australia",
		Flag: "🇦🇺",
	},
	{
		Name: "Austria",
		Flag: "🇦🇹",
	},
	{
		Name: "Azerbaijan",
		Flag: "🇦🇿",
	},
	{
		Name: "Bahamas",
		Flag: "🇧🇸",
	},
	{
		Name: "Bahrain",
		Flag: "🇧🇭",
	},
	{
		Name: "Bangladesh",
		Flag: "🇧🇩",
	},
	{
		Name: "Barbados",
		Flag: "🇧🇧",
	},
	{
		Name: "Belarus",
		Flag: "🇧🇾",
	},
	{
		Name: "Belgium",
		Flag: "🇧🇪",
	},
	{
		Name: "Belize",
		Flag: "🇧🇿",
	},
	{
		Name: "Benin",
		Flag: "🇧🇯",
	},
	{
		Name: "Bermuda",
		Flag: "🇧🇲",
	},
	{
		Name: "Bhutan",
		Flag: "🇧🇹",
	},
	{
		Name: "Bolivia",
		Flag: "🇧🇴",
	},
	{
		Name: "Bosnia and Herzegovina",
		Flag: "🇧🇦",
	},
	{
		Name: "Botswana",
		Flag: "🇧🇼",
	},
	{
		Name: "Bouvet Island",
		Flag: "🇧🇻",
	},
	{
		Name: "Brazil",
		Flag: "🇧🇷",
	},
	{
		Name: "British Indian Ocean Territory",
		Flag: "🇮🇴",
	},
	{
		Name: "Brunei",
		Flag: "🇧🇳",
	},
	{
		Name: "Bulgaria",
		Flag: "🇧🇬",
	},
	{
		Name: "Burkina Faso",
		Flag: "🇧🇫",
	},
	{
		Name: "Burundi",
		Flag: "🇧🇮",
	},
	{
		Name: "Cambodia",
		Flag: "🇰🇭",
	},
	{
		Name: "Cameroon",
		Flag: "🇨🇲",
	},
	{
		Name: "Canada",
		Flag: "🇨🇦",
	},
	{
		Name: "Cape Verde",
		Flag: "🇨🇻",
	},
	{
		Name: "Caribbean Netherlands",
		Flag: "🇧🇶",
	},
	{
		Name: "Cayman Islands",
		Flag: "🇰🇾",
	},
	{
		Name: "Central African Republic",
		Flag: "🇨🇫",
	},
	{
		Name: "Chad",
		Flag: "🇹🇩",
	},
	{
		Name: "Chile",
		Flag: "🇨🇱",
	},
	{
		Name: "China",
		Flag: "🇨🇳",
	},
	{
		Name: "Christmas Island",
		Flag: "🇨🇽",
	},
	{
		Name: "Cocos (Keeling) Islands",
		Flag: "🇨🇨",
	},
	{
		Name: "Colombia",
		Flag: "🇨🇴",
	},
	{
		Name: "Comoros",
		Flag: "🇰🇲",
	},
	{
		Name: "Republic of the Congo",
		Flag: "🇨🇬",
	},
	{
		Name: "DR Congo",
		Flag: "🇨🇩",
	},
	{
		Name: "Cook Islands",
		Flag: "🇨🇰",
	},
	{
		Name: "Costa Rica",
		Flag: "🇨🇷",
	},
	{
		Name: "Côte d'Ivoire (Ivory Coast)",
		Flag: "🇨🇮",
	},
	{
		Name: "Croatia",
		Flag: "🇭🇷",
	},
	{
		Name: "Cuba",
		Flag: "🇨🇺",
	},
	{
		Name: "Curaçao",
		Flag: "🇨🇼",
	},
	{
		Name: "Cyprus",
		Flag: "🇨🇾",
	},
	{
		Name: "Czechia",
		Flag: "🇨🇿",
	},
	{
		Name: "Denmark",
		Flag: "🇩🇰",
	},
	{
		Name: "Djibouti",
		Flag: "🇩🇯",
	},
	{
		Name: "Dominica",
		Flag: "🇩🇲",
	},
	{
		Name: "Dominican Republic",
		Flag: "🇩🇴",
	},
	{
		Name: "Ecuador",
		Flag: "🇪🇨",
	},
	{
		Name: "Egypt",
		Flag: "🇪🇬",
	},
	{
		Name: "El Salvador",
		Flag: "🇸🇻",
	},
	{
		Name: "England",
		Flag: "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	},
	{
		Name: "Equatorial Guinea",
		Flag: "🇬🇶",
	},
	{
		Name: "Eritrea",
		Flag: "🇪🇷",
	},
	{
		Name: "Estonia",
		Flag: "🇪🇪",
	},
	{
		Name: "Eswatini (Swaziland)",
		Flag: "🇸🇿",
	},
	{
		Name: "Ethiopia",
		Flag: "🇪🇹",
	},
	{
		Name: "Falkland Islands",
		Flag: "🇫🇰",
	},
	{
		Name: "Faroe Islands",
		Flag: "🇫🇴",
	},
	{
		Name: "Fiji",
		Flag: "🇫🇯",
	},
	{
		Name: "Finland",
		Flag: "🇫🇮",
	},
	{
		Name: "France",
		Flag: "🇫🇷",
	},
	{
		Name: "French Guiana",
		Flag: "🇬🇫",
	},
	{
		Name: "French Polynesia",
		Flag: "🇵🇫",
	},
	{
		Name: "French Southern and Antarctic Lands",
		Flag: "🇹🇫",
	},
	{
		Name: "Gabon",
		Flag: "🇬🇦",
	},
	{
		Name: "Gambia",
		Flag: "🇬🇲",
	},
	{
		Name: "Georgia",
		Flag: "🇬🇪",
	},
	{
		Name: "Germany",
		Flag: "🇩🇪",
	},
	{
		Name: "Ghana",
		Flag: "🇬🇭",
	},
	{
		Name: "Gibraltar",
		Flag: "🇬🇮",
	},
	{
		Name: "Greece",
		Flag: "🇬🇷",
	},
	{
		Name: "Greenland",
		Flag: "🇬🇱",
	},
	{
		Name: "Grenada",
		Flag: "🇬🇩",
	},
	{
		Name: "Guadeloupe",
		Flag: "🇬🇵",
	},
	{
		Name: "Guam",
		Flag: "🇬🇺",
	},
	{
		Name: "Guatemala",
		Flag: "🇬🇹",
	},
	{
		Name: "Guernsey",
		Flag: "🇬🇬",
	},
	{
		Name: "Guinea",
		Flag: "🇬🇳",
	},
	{
		Name: "Guinea-Bissau",
		Flag: "🇬🇼",
	},
	{
		Name: "Guyana",
		Flag: "🇬🇾",
	},
	{
		Name: "Haiti",
		Flag: "🇭🇹",
	},
	{
		Name: "Heard Island and McDonald Islands",
		Flag: "🇭🇲",
	},
	{
		Name: "Honduras",
		Flag: "🇭🇳",
	},
	{
		Name: "Hong Kong",
		Flag: "🇭🇰",
	},
	{
		Name: "Hungary",
		Flag: "🇭🇺",
	},
	{
		Name: "Iceland",
		Flag: "🇮🇸",
	},
	{
		Name: "India",
		Flag: "🇮🇳",
	},
	{
		Name: "Indonesia",
		Flag: "🇮🇩",
	},
	{
		Name: "Iran",
		Flag: "🇮🇷",
	},
	{
		Name: "Iraq",
		Flag: "🇮🇶",
	},
	{
		Name: "Ireland",
		Flag: "🇮🇪",
	},
	{
		Name: "Isle of Man",
		Flag: "🇮🇲",
	},
	{
		Name: "Italy",
		Flag: "🇮🇹",
	},
	{
		Name: "Jamaica",
		Flag: "🇯🇲",
	},
	{
		Name: "Japan",
		Flag: "🇯🇵",
	},
	{
		Name: "Jersey",
		Flag: "🇯🇪",
	},
	{
		Name: "Jordan",
		Flag: "🇯🇴",
	},
	{
		Name: "Kazakhstan",
		Flag: "🇰🇿",
	},
	{
		Name: "Kenya",
		Flag: "🇰🇪",
	},
	{
		Name: "Kiribati",
		Flag: "🇰🇮",
	},
	{
		Name: "North Korea",
		Flag: "🇰🇵",
	},
	{
		Name: "South Korea",
		Flag: "🇰🇷",
	},
	{
		Name: "Kosovo",
		Flag: "🇽🇰",
	},
	{
		Name: "Kuwait",
		Flag: "🇰🇼",
	},
	{
		Name: "Kyrgyzstan",
		Flag: "🇰🇬",
	},
	{
		Name: "Laos",
		Flag: "🇱🇦",
	},
	{
		Name: "Latvia",
		Flag: "🇱🇻",
	},
	{
		Name: "Lebanon",
		Flag: "🇱🇧",
	},
	{
		Name: "Lesotho",
		Flag: "🇱🇸",
	},
	{
		Name: "Liberia",
		Flag: "🇱🇷",
	},
	{
		Name: "Libya",
		Flag: "🇱🇾",
	},
	{
		Name: "Liechtenstein",
		Flag: "🇱🇮",
	},
	{
		Name: "Lithuania",
		Flag: "🇱🇹",
	},
	{
		Name: "Luxembourg",
		Flag: "🇱🇺",
	},
	{
		Name: "Macau",
		Flag: "🇲🇴",
	},
	{
		Name: "Madagascar",
		Flag: "🇲🇬",
	},
	{
		Name: "Malawi",
		Flag: "🇲🇼",
	},
	{
		Name: "Malaysia",
		Flag: "🇲🇾",
	},
	{
		Name: "Maldives",
		Flag: "🇲🇻",
	},
	{
		Name: "Mali",
		Flag: "🇲🇱",
	},
	{
		Name: "Malta",
		Flag: "🇲🇹",
	},
	{
		Name: "Marshall Islands",
		Flag: "🇲🇭",
	},
	{
		Name: "Martinique",
		Flag: "🇲🇶",
	},
	{
		Name: "Mauritania",
		Flag: "🇲🇷",
	},
	{
		Name: "Mauritius",
		Flag: "🇲🇺",
	},
	{
		Name: "Mayotte",
		Flag: "🇾🇹",
	},
	{
		Name: "Mexico",
		Flag: "🇲🇽",
	},
	{
		Name: "Micronesia",
		Flag: "🇫🇲",
	},
	{
		Name: "Moldova",
		Flag: "🇲🇩",
	},
	{
		Name: "Monaco",
		Flag: "🇲🇨",
	},
	{
		Name: "Mongolia",
		Flag: "🇲🇳",
	},
	{
		Name: "Montenegro",
		Flag: "🇲🇪",
	},
	{
		Name: "Montserrat",
		Flag: "🇲🇸",
	},
	{
		Name: "Morocco",
		Flag: "🇲🇦",
	},
	{
		Name: "Mozambique",
		Flag: "🇲🇿",
	},
	{
		Name: "Myanmar",
		Flag: "🇲🇲",
	},
	{
		Name: "Namibia",
		Flag: "🇳🇦",
	},
	{
		Name: "Nauru",
		Flag: "🇳🇷",
	},
	{
		Name: "Nepal",
		Flag: "🇳🇵",
	},
	{
		Name: "Netherlands",
		Flag: "🇳🇱",
	},
	{
		Name: "New Caledonia",
		Flag: "🇳🇨",
	},
	{
		Name: "New Zealand",
		Flag: "🇳🇿",
	},
	{
		Name: "Nicaragua",
		Flag: "🇳🇮",
	},
	{
		Name: "Niger",
		Flag: "🇳🇪",
	},
	{
		Name: "Nigeria",
		Flag: "🇳🇬",
	},
	{
		Name: "Niue",
		Flag: "🇳🇺",
	},
	{
		Name: "Norfolk Island",
		Flag: "🇳🇫",
	},
	{
		Name: "North Macedonia",
		Flag: "🇲🇰",
	},
	{
		Name: "Northern Mariana Islands",
		Flag: "🇲🇵",
	},
	{
		Name: "Norway",
		Flag: "🇳🇴",
	},
	{
		Name: "Oman",
		Flag: "🇴🇲",
	},
	{
		Name: "Pakistan",
		Flag: "🇵🇰",
	},
	{
		Name: "Palau",
		Flag: "🇵🇼",
	},
	{
		Name: "Palestine",
		Flag: "🇵🇸",
	},
	{
		Name: "Panama",
		Flag: "🇵🇦",
	},
	{
		Name: "Papua New Guinea",
		Flag: "🇵🇬",
	},
	{
		Name: "Paraguay",
		Flag: "🇵🇾",
	},
	{
		Name: "Peru",
		Flag: "🇵🇪",
	},
	{
		Name: "Philippines",
		Flag: "🇵🇭",
	},
	{
		Name: "Pitcairn Islands",
		Flag: "🇵🇳",
	},
	{
		Name: "Poland",
		Flag: "🇵🇱",
	},
	{
		Name: "Portugal",
		Flag: "🇵🇹",
	},
	{
		Name: "Puerto Rico",
		Flag: "🇵🇷",
	},
	{
		Name: "Qatar",
		Flag: "🇶🇦",
	},
	{
		Name: "Réunion",
		Flag: "🇷🇪",
	},
	{
		Name: "Romania",
		Flag: "🇷🇴",
	},
	{
		Name: "Russia",
		Flag: "🇷🇺",
	},
	{
		Name: "Rwanda",
		Flag: "🇷🇼",
	},
	{
		Name: "Saint Barthélemy",
		Flag: "🇧🇱",
	},
	{
		Name: "Saint Helena, Ascension and Tristan da Cunha",
		Flag: "🇸🇭",
	},
	{
		Name: "Saint Kitts and Nevis",
		Flag: "🇰🇳",
	},
	{
		Name: "Saint Lucia",
		Flag: "🇱🇨",
	},
	{
		Name: "Saint Martin",
		Flag: "🇲🇫",
	},
	{
		Name: "Saint Pierre and Miquelon",
		Flag: "🇵🇲",
	},
	{
		Name: "Saint Vincent and the Grenadines",
		Flag: "🇻🇨",
	},
	{
		Name: "Samoa",
		Flag: "🇼🇸",
	},
	{
		Name: "San Marino",
		Flag: "🇸🇲",
	},
	{
		Name: "São Tomé and Príncipe",
		Flag: "🇸🇹",
	},
	{
		Name: "Saudi Arabia",
		Flag: "🇸🇦",
	},
	{
		Name: "Scotland",
		Flag: "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
	},
	{
		Name: "Senegal",
		Flag: "🇸🇳",
	},
	{
		Name: "Serbia",
		Flag: "🇷🇸",
	},
	{
		Name: "Seychelles",
		Flag: "🇸🇨",
	},
	{
		Name: "Sierra Leone",
		Flag: "🇸🇱",
	},
	{
		Name: "Singapore",
		Flag: "🇸🇬",
	},
	{
		Name: "Sint Maarten",
		Flag: "🇸🇽",
	},
	{
		Name: "Slovakia",
		Flag: "🇸🇰",
	},
	{
		Name: "Slovenia",
		Flag: "🇸🇮",
	},
	{
		Name: "Solomon Islands",
		Flag: "🇸🇧",
	},
	{
		Name: "Somalia",
		Flag: "🇸🇴",
	},
	{
		Name: "South Africa",
		Flag: "🇿🇦",
	},
	{
		Name: "South Georgia",
		Flag: "🇬🇸",
	},
	{
		Name: "South Sudan",
		Flag: "🇸🇸",
	},
	{
		Name: "Spain",
		Flag: "🇪🇸",
	},
	{
		Name: "Sri Lanka",
		Flag: "🇱🇰",
	},
	{
		Name: "Sudan",
		Flag: "🇸🇩",
	},
	{
		Name: "Suriname",
		Flag: "🇸🇷",
	},
	{
		Name: "Svalbard and Jan Mayen",
		Flag: "🇸🇯",
	},
	{
		Name: "Sweden",
		Flag: "🇸🇪",
	},
	{
		Name: "Switzerland",
		Flag: "🇨🇭",
	},
	{
		Name: "Syria",
		Flag: "🇸🇾",
	},
	{
		Name: "Taiwan",
		Flag: "🇹🇼",
	},
	{
		Name: "Tajikistan",
		Flag: "🇹🇯",
	},
	{
		Name: "Tanzania",
		Flag: "🇹🇿",
	},
	{
		Name: "Thailand",
		Flag: "🇹🇭",
	},
	{
		Name: "Timor-Leste",
		Flag: "🇹🇱",
	},
	{
		Name: "Togo",
		Flag: "🇹🇬",
	},
	{
		Name: "Tokelau",
		Flag: "🇹🇰",
	},
	{
		Name: "Tonga",
		Flag: "🇹🇴",
	},
	{
		Name: "Trinidad and Tobago",
		Flag: "🇹🇹",
	},
	{
		Name: "Tunisia",
		Flag: "🇹🇳",
	},
	{
		Name: "Turkey",
		Flag: "🇹🇷",
	},
	{
		Name: "Turkmenistan",
		Flag: "🇹🇲",
	},
	{
		Name: "Turks and Caicos Islands",
		Flag: "🇹🇨",
	},
	{
		Name: "Tuvalu",
		Flag: "🇹🇻",
	},
	{
		Name: "Uganda",
		Flag: "🇺🇬",
	},
	{
		Name: "Ukraine",
		Flag: "🇺🇦",
	},
	{
		Name: "United Arab Emirates",
		Flag: "🇦🇪",
	},
	{
		Name: "United Kingdom",
		Flag: "🇬🇧",
	},
	{
		Name: "United States",
		Flag: "🇺🇸",
	},
	{
		Name: "United States Minor Outlying Islands",
		Flag: "🇺🇲",
	},
	{
		Name: "Uruguay",
		Flag: "🇺🇾",
	},
	{
		Name: "Uzbekistan",
		Flag: "🇺🇿",
	},
	{
		Name: "Vanuatu",
		Flag: "🇻🇺",
	},
	{
		Name: "Vatican City (Holy See)",
		Flag: "🇻🇦",
	},
	{
		Name: "Venezuela",
		Flag: "🇻🇪",
	},
	{
		Name: "Vietnam",
		Flag: "🇻🇳",
	},
	{
		Name: "British Virgin Islands",
		Flag: "🇻🇬",
	},
	{
		Name: "United States Virgin Islands",
		Flag: "🇻🇮",
	},
	{
		Name: "Wales",
		Flag: "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
	},
	{
		Name: "Wallis and Futuna",
		Flag: "🇼🇫",
	},
	{
		Name: "Western Sahara",
		Flag: "🇪🇭",
	},
	{
		Name: "Yemen",
		Flag: "🇾🇪",
	},
	{
		Name: "Zambia",
		Flag: "🇿🇲",
	},
	{
		Name: "Zimbabwe",
		Flag: "🇿🇼",
	},
}
