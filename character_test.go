package main

import "testing"

func TestCharacter_UpdateStat(t *testing.T) {
	types := make(map[string]StatType)
	types["test_string"] = StatTypeString
	types["test_int"] = StatTypeInteger
	types["test_float"] = StatTypeFloat
	types["unsupported"] = 10

	s := make(map[string]string)
	i := make(map[string]int)
	f := make(map[string]float64)

	type fields struct {
		StatsTypes   map[string]StatType
		StringStats  map[string]string
		IntegerStats map[string]int
		FloatStats   map[string]float64
	}
	type args struct {
		stat  string
		value any
	}

	_fields := fields{
		StatsTypes:   types,
		StringStats:  s,
		IntegerStats: i,
		FloatStats:   f,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Not found stat", fields{}, args{}, true},
		{"Integer", _fields, args{"test_int", 10}, false},
		{"String", _fields, args{"test_string", "string"}, false},
		{"Float", _fields, args{"test_float", 0.1}, false},
		{"Unsupported", _fields, args{"unsupported", 0.1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Character{
				StatsTypes:   tt.fields.StatsTypes,
				StringStats:  tt.fields.StringStats,
				IntegerStats: tt.fields.IntegerStats,
				FloatStats:   tt.fields.FloatStats,
			}
			if err := d.UpdateStat(tt.args.stat, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Character.UpdateStat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
