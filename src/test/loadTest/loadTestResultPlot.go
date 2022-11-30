package test

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotterLoadTest(allResultsSlice []allResults) error {

	for _, allResults := range allResultsSlice {

		// create plot, idea is x axis is number of concurrent users, y axis is successful run percentage
		p := plot.New()

		p.Title.Text = allResults.Name
		p.X.Label.Text = "Number of concurrent users"
		p.Y.Label.Text = "Successful run percentage"
		data := []interface{}{}
		for _, result := range allResults.Results {
			// data := []interface{}{result.Name, assignPoints(result.Results)}
			data = append(data, result.Name, assignPoints(result.Results))

		}

		err := plotutil.AddLinePoints(
			p,
			data...,
		)

		if err != nil {
			return err
		}
		// Save the plot to a PNG file.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, allResults.Name+".png"); err != nil {
			panic(err)
		}
	}
	return nil
}

// randomPoints returns some random x, y points.
func assignPoints(results []result) plotter.XYs {
	n := len(results)
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = float64(results[i].ConcurrentUsers)
		pts[i].Y = float64(results[i].SuccessPerc)
	}
	return pts
}
