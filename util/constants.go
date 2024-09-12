package util

/// global application variable to print line number in case 
/// '-n' or '-b' flag is passed 
var GLOBAL_COUNTER = 1

/// for this application only GreenColorCode and BlueColorCode will be used
/// ansi color codes
var BlackColorCode int8 =	30
var RedColorCode int8 =	31
var GreenColorCode int8 =	32
var YellowColorCode int8 =	33
var BlueColorCode int8 =	34
var MagentaColorCode int8 =	35
var CyanColorCode int8 =	36
var WhiteColorCode int8 =	37

/// some application constants
var DashReadStdin string = "-"
var NumberFlag string = "-n"
var NumberNonBlankFlag string = "-b"
var EmptyString string = ""
var NewLineCharacter string = "\n"

