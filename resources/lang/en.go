package lang

import (
	"github.com/totoval/framework/helpers/locale"

	"github.com/totoval/framework/resources/lang"
)

func init() {
	validationTranslation := lang.ValidationTranslation{
		Required: "{0} is a required field",
		Len: lang.EmbeddedRule{
			String:  "{0} must be {1} in length",
			Numeric: "{0} must be equal to {1}",
			Array:   "{0} must contain {1}",
		},
		Min: lang.EmbeddedRule{
			String:  "{0} must be at least {1} in length",
			Numeric: "{0} must be {1} or greater",
			Array:   "{0} must contain at least {1}",
		},
		Max: lang.EmbeddedRule{
			String:  "{0} must be a maximum of {1} in length",
			Numeric: "{0} must be {1} or less",
			Array:   "{0} must contain at maximum {1}",
		},
		Eq: "{0} is not equal to {1}",
		Ne: "{0} should not be equal to {1}",
		Lt: lang.EmbeddedRule{
			String:   "{0} must be less than {1} in length",
			Numeric:  "{0} must be less than {1}",
			Array:    "{0} must contain less than {1}",
			Datetime: "{0} must be less than the current Date & Time",
		},
		Lte: lang.EmbeddedRule{
			String:   "{0} must be at maximum {1} in length",
			Numeric:  "{0} must be {1} or less",
			Array:    "{0} must contain at maximum {1}",
			Datetime: "{0} must be less than or equal to the current Date & Time",
		},
		Gt: lang.EmbeddedRule{
			String:   "{0} must be greater than {1} in length",
			Numeric:  "{0} must be greater than {1}",
			Array:    "{0} must contain more than {1}",
			Datetime: "{0} must be greater than the current Date & Time",
		},
		Gte: lang.EmbeddedRule{
			String:   "{0} must be at least {1} in length",
			Numeric:  "{0} must be {1} or greater",
			Array:    "{0} must contain at least {1}",
			Datetime: "{0} must be greater than or equal to the current Date & Time",
		},
		Eqfield:       "{0} must be equal to {1}",
		Eqcsfield:     "{0} must be equal to {1}",
		Necsfield:     "{0} cannot be equal to {1}",
		Gtcsfield:     "{0} must be greater than {1}",
		Gtecsfield:    "{0} must be greater than or equal to {1}",
		Ltcsfield:     "{0} must be less than {1}",
		Ltecsfield:    "{0} must be less than or equal to {1}",
		Nefield:       "{0} cannot be equal to {1}",
		Gtfield:       "{0} must be greater than {1}",
		Gtefield:      "{0} must be greater than or equal to {1}",
		Ltfield:       "{0} must be less than {1}",
		Ltefield:      "{0} must be less than or equal to {1}",
		Alpha:         "{0} can only contain alphabetic characters",
		Alphanum:      "{0} can only contain alphanumeric characters",
		Numeric:       "{0} must be a valid numeric value",
		Number:        "{0} must be a valid number",
		Hexadecimal:   "{0} must be a valid hexadecimal",
		Hexcolor:      "{0} must be a valid HEX color",
		Rgb:           "{0} must be a valid RGB color",
		Rgba:          "{0} must be a valid RGBA color",
		Hsl:           "{0} must be a valid HSL color",
		Hsla:          "{0} must be a valid HSLA color",
		Email:         "{0} must be a valid email address",
		Url:           "{0} must be a valid URL",
		Uri:           "{0} must be a valid URI",
		Base64:        "{0} must be a valid Base64 string",
		Base64url:     "{0} must contains a valid Base64 URL safe value",
		UrnRfc2141:    "{0} must contains a valid URN string",
		Contains:      "{0} must contain the text '{1}'",
		Containsany:   "{0} must contain at least one of the following characters '{1}'",
		Excludes:      "{0} cannot contain the text '{1}'",
		Excludesall:   "{0} cannot contain any of the following characters '{1}'",
		Excludesrune:  "{0} cannot contain the following '{1}'",
		Isbn:          "{0} must be a valid ISBN number",
		Isbn10:        "{0} must be a valid ISBN-10 number",
		Isbn13:        "{0} must be a valid ISBN-13 number",
		Uuid:          "{0} must be a valid UUID",
		Uuid3:         "{0} must be a valid version 3 UUID",
		Uuid4:         "{0} must be a valid version 4 UUID",
		Uuid5:         "{0} must be a valid version 5 UUID",
		Ascii:         "{0} must contain only ascii characters",
		Printascii:    "{0} must contain only printable ascii characters",
		Multibyte:     "{0} must contain multibyte characters",
		Datauri:       "{0} must contain a valid Data URI",
		Latitude:      "{0} must contain valid latitude coordinates",
		Longitude:     "{0} must contain a valid longitude coordinates",
		Ssn:           "{0} must be a valid SSN number",
		Ipv4:          "{0} must be a valid IPv4 address",
		Ipv6:          "{0} must be a valid IPv6 address",
		Ip:            "{0} must be a valid IP address",
		Cidr:          "{0} must contain a valid CIDR notation",
		Cidrv4:        "{0} must contain a valid CIDR notation for an IPv4 address",
		Cidrv6:        "{0} must contain a valid CIDR notation for an IPv6 address",
		TcpAddr:       "{0} must be a valid TCP address",
		Tcp4Addr:      "{0} must be a valid IPv4 TCP address",
		Tcp6Addr:      "{0} must be a valid IPv6 TCP address",
		UdpAddr:       "{0} must be a valid UDP address",
		Udp4Addr:      "{0} must be a valid IPv4 UDP address",
		Udp6Addr:      "{0} must be a valid IPv6 UDP address",
		IpAddr:        "{0} must be a resolvable IP address",
		Ip4Addr:       "{0} must be a resolvable IPv4 address",
		Ip6Addr:       "{0} must be a resolvable IPv6 address",
		UnixAddr:      "{0} must be a resolvable UNIX address",
		Mac:           "{0} must contain a valid MAC address",
		Unique:        "{0} must contain unique values",
		Iscolor:       "{0} must be a valid color",
		Oneof:         "{0} must be one of [{1}]",
		BtcAddr:       "{0} must be a valid bitcoin address",
		BtcAddrBech32: "{0} must be a valid bitcoin Bech32 address",
		EthAddr:       "{0} must be a valid ethereum address",

		PluralRuleMap: map[string]lang.PluralRule{
			"character": {
				One:   "character",
				Other: "characters",
			},
			"item": {
				One:   "item",
				Other: "items",
			},
		},

		FieldTranslation: lang.ValidationFieldTranslation{
			"Email": "email",
		},
	}
	customTranslation := lang.CustomTranslation{
		"auth.register.failed_existed":              "user register failed, user has been registered before",
		"auth.register.failed_system_error":         "user register failed, system error",
		"auth.register.failed_token_generate_error": "user register failed, token generate failed",

		"auth.login.failed_not_exist":            "user login failed, user doesn't exist",
		"auth.login.failed_wrong_password":       "user login failed, user password incorrect",
		"auth.login.failed_token_generate_error": "user login failed, token generate failed",
	}

	locale.AddLocale("en", &customTranslation, &validationTranslation)
}
