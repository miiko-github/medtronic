// Code generated by "stringer -type Command"; DO NOT EDIT

package medtronic

import "fmt"

const _Command_name = "AckNakSetClockSetAbsoluteTempBasalSuspendWakeupSetPercentTempBasalClockPumpIdBatteryReservoirHistoryPageCarbUnitsGlucoseUnitsCarbRatiosInsulinSensitivitiesModelBasalRatesBasalPatternABasalPatternBTempBasalLastHistoryPageGlucoseTargetsSettingsStatus"

var _Command_map = map[Command]string{
	6:   _Command_name[0:3],
	21:  _Command_name[3:6],
	64:  _Command_name[6:14],
	76:  _Command_name[14:34],
	77:  _Command_name[34:41],
	93:  _Command_name[41:47],
	105: _Command_name[47:66],
	112: _Command_name[66:71],
	113: _Command_name[71:77],
	114: _Command_name[77:84],
	115: _Command_name[84:93],
	128: _Command_name[93:104],
	136: _Command_name[104:113],
	137: _Command_name[113:125],
	138: _Command_name[125:135],
	139: _Command_name[135:155],
	141: _Command_name[155:160],
	146: _Command_name[160:170],
	147: _Command_name[170:183],
	148: _Command_name[183:196],
	152: _Command_name[196:205],
	157: _Command_name[205:220],
	159: _Command_name[220:234],
	192: _Command_name[234:242],
	206: _Command_name[242:248],
}

func (i Command) String() string {
	if str, ok := _Command_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Command(%d)", i)
}
