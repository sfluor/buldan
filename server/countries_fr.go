package main

// From ChatGPT TODO double check
var frenchCountries = []rawCountry{
	{
		name: "Afghanistan",
		flag: "🇦🇫",
	},
	{
		name: "Afrique du Sud",
		flag: "🇿🇦",
	},
	{
		name: "Albanie",
		flag: "🇦🇱",
	},
	{
		name: "Algérie",
		flag: "🇩🇿",
	},
	{
		name: "Allemagne",
		flag: "🇩🇪",
	},
	{
		name: "Andorre",
		flag: "🇦🇩",
	},
	{
		name: "Angola",
		flag: "🇦🇴",
	},
	{
		name: "Antigua-et-Barbuda",
		flag: "🇦🇬",
	},
	{
		name: "Arabie saoudite",
		flag: "🇸🇦",
	},
	{
		name: "Argentine",
		flag: "🇦🇷",
	},
	{
		name: "Arménie",
		flag: "🇦🇲",
	},
	{
		name: "Australie",
		flag: "🇦🇺",
	},
	{
		name: "Autriche",
		flag: "🇦🇹",
	},
	{
		name: "Azerbaïdjan",
		flag: "🇦🇿",
	},
	{
		name: "Bahamas",
		flag: "🇧🇸",
	},
	{
		name: "Bahreïn",
		flag: "🇧🇭",
	},
	{
		name: "Bangladesh",
		flag: "🇧🇩",
	},
	{
		name: "Barbade",
		flag: "🇧🇧",
	},
	{
		name: "Belgique",
		flag: "🇧🇪",
	},
	{
		name: "Belize",
		flag: "🇧🇿",
	},
	{
		name: "Bénin",
		flag: "🇧🇯",
	},
	{
		name: "Bhoutan",
		flag: "🇧🇹",
	},
	{
		name: "Biélorussie",
		flag: "🇧🇾",
	},
	{
		name: "Birmanie",
		flag: "🇲🇲",
	},
	{
		name: "Bolivie",
		flag: "🇧🇴",
	},
	{
		name: "Bosnie-Herzégovine",
		flag: "🇧🇦",
	},
	{
		name: "Botswana",
		flag: "🇧🇼",
	},
	{
		name: "Brésil",
		flag: "🇧🇷",
	},
	{
		name: "Brunei",
		flag: "🇧🇳",
	},
	{
		name: "Bulgarie",
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
		name: "Cambodge",
		flag: "🇰🇭",
	},
	{
		name: "Cameroun",
		flag: "🇨🇲",
	},
	{
		name: "Canada",
		flag: "🇨🇦",
	},
	{
		name: "Cap-Vert",
		flag: "🇨🇻",
	},
	{
		name: "République centrafricaine",
		flag: "🇨🇫",
	},
	{
		name: "Chili",
		flag: "🇨🇱",
	},
	{
		name: "Chine",
		flag: "🇨🇳",
	},
	{
		name: "Chypre",
		flag: "🇨🇾",
	},
	{
		name: "Colombie",
		flag: "🇨🇴",
	},
	{
		name: "Comores",
		flag: "🇰🇲",
	},
	{
		name: "République du Congo",
		flag: "🇨🇬",
	},
	{
		name: "République démocratique du Congo",
		flag: "🇨🇩",
	},
	{
		name: "Corée du Nord",
		flag: "🇰🇵",
	},
	{
		name: "Corée du Sud",
		flag: "🇰🇷",
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
		name: "Croatie",
		flag: "🇭🇷",
	},
	{
		name: "Cuba",
		flag: "🇨🇺",
	},
	{
		name: "Danemark",
		flag: "🇩🇰",
	},
	{
		name: "Djibouti",
		flag: "🇩🇯",
	},
	{
		name: "Dominique",
		flag: "🇩🇲",
	},
	{
		name: "République dominicaine",
		flag: "🇩🇴",
	},
	{
		name: "Égypte",
		flag: "🇪🇬",
	},
	{
		name: "Émirats arabes unis",
		flag: "🇦🇪",
	},
	{
		name: "Équateur",
		flag: "🇪🇨",
	},
	{
		name: "Érythrée",
		flag: "🇪🇷",
	},
	{
		name: "Espagne",
		flag: "🇪🇸",
	},
	{
		name: "Estonie",
		flag: "🇪🇪",
	},
	{
		name: "Eswatini",
		flag: "🇸🇿",
	},
	{
		name: "États-Unis",
		flag: "🇺🇸",
	},
	{
		name: "Éthiopie",
		flag: "🇪🇹",
	},
	{
		name: "Fidji",
		flag: "🇫🇯",
	},
	{
		name: "Finlande",
		flag: "🇫🇮",
	},
	{
		name: "France",
		flag: "🇫🇷",
	},
	{
		name: "Gabon",
		flag: "🇬🇦",
	},
	{
		name: "Gambie",
		flag: "🇬🇲",
	},
	{
		name: "Géorgie",
		flag: "🇬🇪",
	},
	{
		name: "Ghana",
		flag: "🇬🇭",
	},
	{
		name: "Grèce",
		flag: "🇬🇷",
	},
	{
		name: "Grenade",
		flag: "🇬🇩",
	},
	{
		name: "Guatemala",
		flag: "🇬🇹",
	},
	{
		name: "Guinée",
		flag: "🇬🇳",
	},
	{
		name: "Guinée-Bissau",
		flag: "🇬🇼",
	},
	{
		name: "Guinée équatoriale",
		flag: "🇬🇶",
	},
	{
		name: "Guyana",
		flag: "🇬🇾",
	},
	{
		name: "Haïti",
		flag: "🇭🇹",
	},
	{
		name: "Honduras",
		flag: "🇭🇳",
	},
	{
		name: "Hongrie",
		flag: "🇭🇺",
	},
	{
		name: "Îles Cook",
		flag: "🇨🇰",
	},
	{
		name: "Îles Marshall",
		flag: "🇲🇭",
	},
	{
		name: "Îles Salomon",
		flag: "🇸🇧",
	},
	{
		name: "Inde",
		flag: "🇮🇳",
	},
	{
		name: "Indonésie",
		flag: "🇮🇩",
	},
	{
		name: "Irak",
		flag: "🇮🇶",
	},
	{
		name: "Iran",
		flag: "🇮🇷",
	},
	{
		name: "Irlande",
		flag: "🇮🇪",
	},
	{
		name: "Islande",
		flag: "🇮🇸",
	},
	{
		name: "Italie",
		flag: "🇮🇹",
	},
	{
		name: "Jamaïque",
		flag: "🇯🇲",
	},
	{
		name: "Japon",
		flag: "🇯🇵",
	},
	{
		name: "Jordanie",
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
		name: "Kirghizistan",
		flag: "🇰🇬",
	},
	{
		name: "Kiribati",
		flag: "🇰🇮",
	},
	{
		name: "Kosovo",
		flag: "🇽🇰",
	},
	{
		name: "Koweït",
		flag: "🇰🇼",
	},
	{
		name: "Laos",
		flag: "🇱🇦",
	},
	{
		name: "Lesotho",
		flag: "🇱🇸",
	},
	{
		name: "Lettonie",
		flag: "🇱🇻",
	},
	{
		name: "Liban",
		flag: "🇱🇧",
	},
	{
		name: "Liberia",
		flag: "🇱🇷",
	},
	{
		name: "Libye",
		flag: "🇱🇾",
	},
	{
		name: "Liechtenstein",
		flag: "🇱🇮",
	},
	{
		name: "Lituanie",
		flag: "🇱🇹",
	},
	{
		name: "Luxembourg",
		flag: "🇱🇺",
	},
	{
		name: "Macédoine du Nord",
		flag: "🇲🇰",
	},
	{
		name: "Madagascar",
		flag: "🇲🇬",
	},
	{
		name: "Malaisie",
		flag: "🇲🇾",
	},
	{
		name: "Malawi",
		flag: "🇲🇼",
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
		name: "Malte",
		flag: "🇲🇹",
	},
	{
		name: "Maroc",
		flag: "🇲🇦",
	},
	{
		name: "Maurice",
		flag: "🇲🇺",
	},
	{
		name: "Mauritanie",
		flag: "🇲🇷",
	},
	{
		name: "Mexique",
		flag: "🇲🇽",
	},
	{
		name: "Micronésie",
		flag: "🇫🇲",
	},
	{
		name: "Moldavie",
		flag: "🇲🇩",
	},
	{
		name: "Monaco",
		flag: "🇲🇨",
	},
	{
		name: "Mongolie",
		flag: "🇲🇳",
	},
	{
		name: "Monténégro",
		flag: "🇲🇪",
	},
	{
		name: "Mozambique",
		flag: "🇲🇿",
	},
	{
		name: "Namibie",
		flag: "🇳🇦",
	},
	{
		name: "Nauru",
		flag: "🇳🇷",
	},
	{
		name: "Népal",
		flag: "🇳🇵",
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
		name: "Norvège",
		flag: "🇳🇴",
	},
	{
		name: "Nouvelle-Zélande",
		flag: "🇳🇿",
	},
	{
		name: "Oman",
		flag: "🇴🇲",
	},
	{
		name: "Ouganda",
		flag: "🇺🇬",
	},
	{
		name: "Ouzbékistan",
		flag: "🇺🇿",
	},
	{
		name: "Pakistan",
		flag: "🇵🇰",
	},
	{
		name: "Palaos",
		flag: "🇵🇼",
	},
	{
		name: "Panama",
		flag: "🇵🇦",
	},
	{
		name: "Papouasie-Nouvelle-Guinée",
		flag: "🇵🇬",
	},
	{
		name: "Paraguay",
		flag: "🇵🇾",
	},
	{
		name: "Pays-Bas",
		flag: "🇳🇱",
	},
	{
		name: "Pérou",
		flag: "🇵🇪",
	},
	{
		name: "Philippines",
		flag: "🇵🇭",
	},
	{
		name: "Pologne",
		flag: "🇵🇱",
	},
	{
		name: "Portugal",
		flag: "🇵🇹",
	},
	{
		name: "Qatar",
		flag: "🇶🇦",
	},
	{
		name: "Roumanie",
		flag: "🇷🇴",
	},
	{
		name: "Royaume-Uni",
		flag: "🇬🇧",
	},
	{
		name: "Russie",
		flag: "🇷🇺",
	},
	{
		name: "Rwanda",
		flag: "🇷🇼",
	},
	{
		name: "Saint-Kitts-et-Nevis",
		flag: "🇰🇳",
	},
	{
		name: "Saint-Marin",
		flag: "🇸🇲",
	},
	{
		name: "Saint-Vincent-et-les-Grenadines",
		flag: "🇻🇨",
	},
	{
		name: "Sainte-Lucie",
		flag: "🇱🇨",
	},
	{
		name: "Salvador",
		flag: "🇸🇻",
	},
	{
		name: "Samoa",
		flag: "🇼🇸",
	},
	{
		name: "São Tomé-et-Principe",
		flag: "🇸🇹",
	},
	{
		name: "Sénégal",
		flag: "🇸🇳",
	},
	{
		name: "Serbie",
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
		name: "Singapour",
		flag: "🇸🇬",
	},
	{
		name: "Slovaquie",
		flag: "🇸🇰",
	},
	{
		name: "Slovénie",
		flag: "🇸🇮",
	},
	{
		name: "Somalie",
		flag: "🇸🇴",
	},
	{
		name: "Soudan",
		flag: "🇸🇩",
	},
	{
		name: "Soudan du Sud",
		flag: "🇸🇸",
	},
	{
		name: "Sri Lanka",
		flag: "🇱🇰",
	},
	{
		name: "Suède",
		flag: "🇸🇪",
	},
	{
		name: "Suisse",
		flag: "🇨🇭",
	},
	{
		name: "Suriname",
		flag: "🇸🇷",
	},
	{
		name: "Syrie",
		flag: "🇸🇾",
	},
	{
		name: "Tadjikistan",
		flag: "🇹🇯",
	},
	{
		name: "Tanzanie",
		flag: "🇹🇿",
	},
	{
		name: "Tchad",
		flag: "🇹🇩",
	},
	{
		name: "République tchèque",
		flag: "🇨🇿",
	},
	{
		name: "Thaïlande",
		flag: "🇹🇭",
	},
	{
		name: "Timor oriental",
		flag: "🇹🇱",
	},
	{
		name: "Togo",
		flag: "🇹🇬",
	},
	{
		name: "Tonga",
		flag: "🇹🇴",
	},
	{
		name: "Trinité-et-Tobago",
		flag: "🇹🇹",
	},
	{
		name: "Tunisie",
		flag: "🇹🇳",
	},
	{
		name: "Turkménistan",
		flag: "🇹🇲",
	},
	{
		name: "Turquie",
		flag: "🇹🇷",
	},
	{
		name: "Tuvalu",
		flag: "🇹🇻",
	},
	{
		name: "Ukraine",
		flag: "🇺🇦",
	},
	{
		name: "Uruguay",
		flag: "🇺🇾",
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
		name: "Viêt Nam",
		flag: "🇻🇳",
	},
	{
		name: "Yémen",
		flag: "🇾🇪",
	},
	{
		name: "Zambie",
		flag: "🇿🇲",
	},
	{
		name: "Zimbabwe",
		flag: "🇿🇼",
	},
}
