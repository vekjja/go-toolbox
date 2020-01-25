package darksky

import "math"

// UnitMeasures are the location specific terms for weather data.
type UnitMeasures struct {
	Degrees       string
	Speed         string
	Length        string
	Precipitation string
}

// ValidUnits : valid unit types for Dark Sky API
var ValidUnits = map[string]string{
	"auto": "Determin Units Based on Location",
	"ca":   "same as si, except uses kilometers per hour",
	"uk":   "same as si, except uses kilometers, and miles per hour",
	"uk2":  "same as si, except uses miles, and miles per hour",
	"us":   "Imperial units",
	"si":   "International System of Units",
}

// UnitFormats describe each regions UnitMeasures.
var UnitFormats = map[string]UnitMeasures{
	"us": {
		Degrees:       "°F",
		Speed:         "mph",
		Length:        "miles",
		Precipitation: "in/hr",
	},
	"si": {
		Degrees:       "°C",
		Speed:         "m/s",
		Length:        "kilometers",
		Precipitation: "mm/h",
	},
	"ca": {
		Degrees:       "°C",
		Speed:         "km/h",
		Length:        "kilometers",
		Precipitation: "mm/h",
	},
	// deprecated, use "uk2" in stead
	"uk": {
		Degrees:       "°C",
		Speed:         "mph",
		Length:        "kilometers",
		Precipitation: "mm/h",
	},
	"uk2": {
		Degrees:       "°C",
		Speed:         "mph",
		Length:        "miles",
		Precipitation: "mm/h",
	},
}

// Icons emoji used to represent current weather
var Icons = map[string]string{
	"rain":                "🌧",
	"clear-day":           "☀️",
	"clear-night":         "🌙",
	"snow":                "🌨☃️",
	"sleet":               "❆🌧❅",
	"wind":                "💨",
	"fog":                 "🌫",
	"cloudy":              "☁",
	"partly-cloudy-day":   "🌤",
	"partly-cloudy-night": "☁🌙",
}

// MoonPhaseIcon : return Moon Phase Icon based on lunation number
func MoonPhaseIcon(phase float64) string {
	var icon string
	switch {
	case phase == 0:
		icon = "🌑"
	case phase > 0 && phase < 0.25:
		icon = "🌒"
	case phase == 0.25:
		icon = "🌓"
	case phase > 0.25 && phase < 0.5:
		icon = "🌔"
	case phase == 0.5:
		icon = "🌕"
	case phase >= 0.5 && phase < 0.75:
		icon = "🌖"
	case phase == 0.75:
		icon = "🌗"
	case phase > 0.75:
		icon = "🌘"
	}
	return icon
}

// Directions contain all the combinations of N,S,E,W
var Directions = []string{
	"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
}

// GetBearings : get directional bearings based on degrees
func GetBearings(degrees float64) string {
	index := int(math.Mod((degrees+11.25)/22.5, 16))
	return Directions[index]
}
