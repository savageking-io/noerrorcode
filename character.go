package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type StatType uint

const (
	StatTypeString  StatType = 0
	StatTypeInteger StatType = 1
	StatTypeFloat   StatType = 2
)

type Character struct {
	StatsTypes   map[string]StatType
	StringStats  map[string]string
	IntegerStats map[string]int
	FloatStats   map[string]float64
}

func (d *Character) Init(statData bson.M) error {
	d.StatsTypes = make(map[string]StatType)
	d.StringStats = make(map[string]string)
	d.IntegerStats = make(map[string]int)
	d.FloatStats = make(map[string]float64)

	return nil
}

func (d *Character) UpdateStat(stat string, value any) error {
	log.Traceln("Character::UpdateStat")

	t, ok := d.StatsTypes[stat]
	if !ok {
		return fmt.Errorf("unknown stat: %s", stat)
	}

	switch t {
	case StatTypeString:
		d.StringStats[stat] = value.(string)
		return nil
	case StatTypeInteger:
		d.IntegerStats[stat] = value.(int)
		return nil
	case StatTypeFloat:
		d.FloatStats[stat] = value.(float64)
		return nil
	}

	return fmt.Errorf("unsupported stat type")
}
