package mongo

type Battery struct {
	BatteryLevel float64 `bson:"battery_level"`
	LowPowerMode bool    `bson:"low_power_mode"`
	BatteryState string  `bson:"battery_state"`
}
