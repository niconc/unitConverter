package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// We're going to use plain vars here with the analogous functions
	// and not pointer vars.
	fU     string  // from unit: the name of the unit to convert from.
	tU     string  // to unit: the name of the unit to convert to.
	numF   float64 // the value to be converted.
	numRes float64 // the final converted value.
	uConv  = map[string]float64{"meters_feet": 3.280839895, "meters_kilometers": 0.001, "meters_miles": 0.0006213689, "feet_meters": 0.3048, "feet_kilometers": 0.0003048, "feet_miles": 0.0001893932, "kilometers_meters": 1000.0, "kilometers_feet": 3280.839895, "kilometers_miles": 0.6213688756, "miles_meters": 1609.35, "miles_feet": 5280.019685, "miles_kilometers": 1.60935}
)

func main() {
	fmt.Println("Unit Converter")
	fmt.Printf("Command-line conversion utility for length units.\nIt supports the units: meters, feet, kilometers and miles.\n\n")

	// flags: the format we're using is:
	// [SAMPLE]->  $ go run unitConverter.go -from=meters -to=feet 25  <-[SAMPLE]
	flag.StringVar(&fU, "from", "meters", "The unit we are converting from.")
	flag.StringVar(&tU, "to", "meters", "The unit we are converting to.")

	// Parse() the command-line flags from os.Args[1:] and assign to the vars defined.
	// Parse() is parsing only the flags (like -from=... OR -to=...)
	// and NOT values alone (like 25) FROM SAMPLE above.
	flag.Parse()

	// arguments from os package perspective -- All Arguments:
	fmt.Println("os.Args:", os.Args)
	// arguments from flags package perspective (after having parsed the flags)
	// only the non-flag arguments (args without name, like the number to convert from):
	fmt.Println("flag.Args:", flag.Args())

	// check if -from -to flags are empty and return:
	if fU == "" {
		fmt.Println("Please provide a unit to convert from.")
		return
	}

	if tU == "" {
		fmt.Println("Please provide a unit to convert to.")
		return
	}

	// check if any non-flag (without name) arguments are passed: (there's only one in
	// this format we're using: the number to be converted at the end)
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a value to convert.")
		return
	}

	// Scan to find the armuments text for the number value to be converted,
	// using Sscanf() and convert it from string to float64.
	_, err := fmt.Sscanf(os.Args[3], "%f", &numF)
	if err != nil {
		fmt.Println("Unit value fail to scanned:", err.Error())
	}

	// Conversion
	numRes, ok := ConvToUnit(fU, tU) // call the function for conversion
	if !ok {
		fmt.Println("Conversion Failed. Unit might not exist or units are the same.")
		return
	}
	// the final result: "1223.00 kilometers are 4012467.191585 feet (true)"
	fmt.Printf("%.2f %s are %f %s (%t)\n", numF, fU, numRes, tU, ok)
}

// ConvToUnit function is converting between units.
// Units used: [meters, feet, kilometers, miles]
func ConvToUnit(from, to string) (float64, bool) {
	unitVal, ok := uConv[from+"_"+to]
	if !ok {
		return 0.0, false
	}
	return (numF * unitVal), true
}

/*
This is the first function with switch{}. I didn't use it, instead I'm using maps
func ConvToUnit_OLD(from, to string) (float64, bool) {
	switch {
	// meters to ...
	case from == "meters" && to == "feet": // meters to feet: 1 X 3.280839895.
		return numF * 3.280839895, true
	case from == "meters" && to == "kilometers": // meters to kilometers: 1 X 0.001.
		return numF * 0.001, true
	case from == "meters" && to == "miles": // meters to miles: 1 X 0.0006213689.
		return numF * 0.0006213689, true
	// feet to...
	case from == "feet" && to == "meters": // feet to meters: 1 X 0.3048.
		return numF * 0.3048, true
	case from == "feet" && to == "kilometers": // feet to kilometers: 1 X 0.3048.
		return numF * 0.0003048, true
	case from == "feet" && to == "miles": // feet to miles: 1 X 0.0003048.
		return numF * 0.0001893932, true
	// kilometers to ...
	case from == "kilometers" && to == "meters": // kilometers to meters: 1 X 1000.
		return numF * 1000, true
	case from == "kilometers" && to == "feet": // kilometers to feet: 1 X 3280.839895.
		return numF * 3280.839895, true
	case from == "kilometers" && to == "miles": // kilometers to miles: 1 X 0.6213688756.
		return numF * 0.6213688756, true
	// miles to ...
	case from == "miles" && "to" == "meters": // miles to meters: 1 X 1609.35.
		return numF * 1609.35, true
	case from == "miles" && "to" == "feet": // miles to meters: 1 X 5280.019685.
		return numF * 5280.019685, true
	case from == "miles" && "to" == "kilometers": // miles to meters: 1 X 1.60935
		return numF * 1.60935, true
	default:
		return 0.0, false
	}
}
*/
