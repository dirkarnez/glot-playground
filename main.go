package main

import (
	"fmt"

	"github.com/Arafatk/glot"
)

func main() {
	dimensions := 2
	// The dimensions supported by the plot
	persist := true
	debug := true
	plot, err := glot.NewPlot(dimensions, persist, debug)
	if err != nil {
		fmt.Println(err.Error())
	}
	pointGroupName := "Simple Circles"
	style := "circle"
	points := [][]float64{{7, 3, 13, 5.6, 11.1}, {12, 13, 11, 1, 7}}
	// Adding a point group
	err = plot.AddPointGroup(pointGroupName, style, points)
	if err != nil {
		fmt.Println(err.Error())
	}
	pointGroupName = "Simple Lines"
	style = "lines"
	points = [][]float64{{7, 3, 3, 5.6, 5.6, 7, 7, 9, 13, 13, 9, 9}, {10, 10, 4, 4, 5.4, 5.4, 4, 4, 4, 10, 10, 4}}
	plot.AddPointGroup(pointGroupName, style, points)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle("Example Plot")
	// Optional: Setting the title of the plot
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	// Optional: Setting label for X and Y axis
	plot.SetXrange(-2, 18)
	plot.SetYrange(-2, 18)
	// Optional: Setting axis ranges
	err = plot.SavePlot("./123.png")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("done")
}
