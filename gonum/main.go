package main

// Created By: Gavin, 2020-06-12
// Modified By: Gavin, 2020-06-15

// 【链接】https://github.com/gonum/plot/wiki/Example-plots
// 【成功】例子1：Bar Charts
// 【成功】例子2：Quartile Plots
// 注意1：例子需要自己下载一些依赖库，放到 /usr/local/go/src 文件夹中
// 注意2：github.com, golang.org, gonum.org中的库都是我新增的（从Github上下载压缩包，解压之后直接放入对应的位置即可）

// 例子：Quartile Plots
import (
	"gonum.org/v1/plot/vg"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {
	// Get some data to display in our plot.
	rand.Seed(int64(0))
	n := 10
	uniform := make(plotter.Values, n)
	normal := make(plotter.Values, n)
	expon := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		uniform[i] = rand.Float64()
		normal[i] = rand.NormFloat64()
		expon[i] = rand.ExpFloat64()
	}

	// Create the plot and set its title and axis label.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Quartile plots"
	p.Y.Label.Text = "Values"

	// Make boxes for our data and add them to the plot.
	q0, err := plotter.NewQuartPlot(0, uniform)
	if err != nil {
		panic(err)
	}
	q1, err := plotter.NewQuartPlot(1, normal)
	if err != nil {
		panic(err)
	}
	q2, err := plotter.NewQuartPlot(2, expon)
	if err != nil {
		panic(err)
	}
	p.Add(q0, q1, q2)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	if err := p.Save(3*vg.Inch, 4*vg.Inch, "quartile.png"); err != nil {
		panic(err)
	}
}
