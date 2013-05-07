package mapping

import (
	"goposm/element"
)

type TagMap map[string]map[string]bool

var PointTags TagMap
var WayTags TagMap
var RelationTags TagMap

type TagFilter interface {
	Filter(element.Tags) bool
}

func (m TagMap) Filter(tags element.Tags) bool {
	for k, v := range tags {
		values, ok := m[k]
		if !ok {
			delete(tags, k)
		} else {
			if _, ok := values["__any__"]; ok {
				continue
			} else if _, ok := values[v]; ok {
				continue
			} else {
				delete(tags, k)
			}
		}
	}
	if _, ok := tags["name"]; ok && len(tags) == 1 {
		delete(tags, "name")
	}
	if len(tags) > 0 {
		return true
	} else {
		return false
	}
}

// default mapping created from imposm defaultmapping.py
// TODO make configurable
func init() {
	PointTags = TagMap{
		"aeroway": map[string]bool{
			"aerodome": true,
			"gate":     true,
			"helipad":  true,
			"terminal": true,
		},
		"amenity": map[string]bool{
			"fire_station": true,
			"fuel":         true,
			"hospital":     true,
			"library":      true,
			"police":       true,
			"school":       true,
			"townhall":     true,
			"university":   true,
		},
		"highway": map[string]bool{
			"bus_stop":          true,
			"motorway_junction": true,
			"turning_circle":    true,
		},
		"name": map[string]bool{
			"__any__": true,
		},
		"place": map[string]bool{
			"city":     true,
			"country":  true,
			"county":   true,
			"hamlet":   true,
			"locality": true,
			"region":   true,
			"state":    true,
			"suburb":   true,
			"town":     true,
			"village":  true,
		},
		"population": map[string]bool{
			"__any__": true,
		},
		"railway": map[string]bool{
			"crossing":        true,
			"halt":            true,
			"level_crossing":  true,
			"station":         true,
			"subway_entrance": true,
			"tram_stop":       true,
		},
		"ref": map[string]bool{
			"__any__": true,
		},
	}
	WayTags = map[string]map[string]bool{
		"admin_level": map[string]bool{
			"__any__": true,
		},
		"aeroway": map[string]bool{
			"aerodrome": true,
			"apron":     true,
			"helipad":   true,
			"runway":    true,
			"taxiway":   true,
			"terminal":  true,
		},
		"amenity": map[string]bool{
			"cinema":           true,
			"college":          true,
			"fuel":             true,
			"hospital":         true,
			"library":          true,
			"parking":          true,
			"place_of_worship": true,
			"school":           true,
			"theatre":          true,
			"university":       true,
		},
		"area": map[string]bool{
			"__any__": true,
		},
		"boundary": map[string]bool{
			"administrative": true,
		},
		"bridge": map[string]bool{
			"__any__": true,
		},
		"building": map[string]bool{
			"__any__": true,
		},
		"highway": map[string]bool{
			"bridleway":      true,
			"cycleway":       true,
			"footway":        true,
			"living_street":  true,
			"motorway":       true,
			"motorway_link":  true,
			"path":           true,
			"pedestrian":     true,
			"primary":        true,
			"primary_link":   true,
			"residential":    true,
			"road":           true,
			"secondary":      true,
			"secondary_link": true,
			"service":        true,
			"steps":          true,
			"tertiary":       true,
			"track":          true,
			"trunk":          true,
			"trunk_link":     true,
			"unclassified":   true,
		},
		"landuse": map[string]bool{
			"allotments":        true,
			"basin":             true,
			"cemetery":          true,
			"commercial":        true,
			"farm":              true,
			"farmland":          true,
			"farmyard":          true,
			"forest":            true,
			"grass":             true,
			"industrial":        true,
			"meadow":            true,
			"park":              true,
			"quarry":            true,
			"railway":           true,
			"recreation_ground": true,
			"reservoir":         true,
			"residential":       true,
			"retail":            true,
			"village_green":     true,
			"wood":              true,
		},
		"leisure": map[string]bool{
			"common":         true,
			"garden":         true,
			"golf_course":    true,
			"nature_reserve": true,
			"park":           true,
			"pitch":          true,
			"playground":     true,
			"sports_centre":  true,
			"stadium":        true,
		},
		"name": map[string]bool{
			"__any__": true,
		},
		"natural": map[string]bool{
			"land":  true,
			"scrub": true,
			"water": true,
			"wood":  true,
		},
		"oneway": map[string]bool{
			"__any__": true,
		},
		"railway": map[string]bool{
			"funicular":    true,
			"light_rail":   true,
			"monorail":     true,
			"narrow_gauge": true,
			"preserved":    true,
			"rail":         true,
			"station":      true,
			"subway":       true,
			"tram":         true,
		},
		"ref": map[string]bool{
			"__any__": true,
		},
		"tunnel": map[string]bool{
			"__any__": true,
		},
		"waterway": map[string]bool{
			"canal":     true,
			"drain":     true,
			"river":     true,
			"riverbank": true,
			"stream":    true,
		}}

	RelationTags = map[string]map[string]bool{

		"admin_level": map[string]bool{
			"__any__": true,
		},
		"aeroway": map[string]bool{
			"aerodrome": true,
			"apron":     true,
			"helipad":   true,
			"runway":    true,
			"taxiway":   true,
			"terminal":  true,
		},
		"amenity": map[string]bool{
			"cinema":           true,
			"college":          true,
			"fuel":             true,
			"hospital":         true,
			"library":          true,
			"parking":          true,
			"place_of_worship": true,
			"school":           true,
			"theatre":          true,
			"university":       true,
		},
		"area": map[string]bool{
			"__any__": true,
		},
		"boundary": map[string]bool{
			"administrative": true,
		},
		"bridge": map[string]bool{
			"__any__": true,
		},
		"building": map[string]bool{
			"__any__": true,
		},
		"highway": map[string]bool{
			"bridleway":      true,
			"cycleway":       true,
			"footway":        true,
			"living_street":  true,
			"motorway":       true,
			"motorway_link":  true,
			"path":           true,
			"pedestrian":     true,
			"primary":        true,
			"primary_link":   true,
			"residential":    true,
			"road":           true,
			"secondary":      true,
			"secondary_link": true,
			"service":        true,
			"steps":          true,
			"tertiary":       true,
			"track":          true,
			"trunk":          true,
			"trunk_link":     true,
			"unclassified":   true,
		},
		"landuse": map[string]bool{
			"allotments":        true,
			"basin":             true,
			"cemetery":          true,
			"commercial":        true,
			"farm":              true,
			"farmland":          true,
			"farmyard":          true,
			"forest":            true,
			"grass":             true,
			"industrial":        true,
			"meadow":            true,
			"park":              true,
			"quarry":            true,
			"railway":           true,
			"recreation_ground": true,
			"reservoir":         true,
			"residential":       true,
			"retail":            true,
			"village_green":     true,
			"wood":              true,
		},
		"leisure": map[string]bool{
			"common":         true,
			"garden":         true,
			"golf_course":    true,
			"nature_reserve": true,
			"park":           true,
			"pitch":          true,
			"playground":     true,
			"sports_centre":  true,
			"stadium":        true,
		},
	}
}