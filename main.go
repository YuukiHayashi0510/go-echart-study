package main

import (
	"ges/graph"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func main() {
	page := components.NewPage()
	page.AddCharts(graph.BarStack())
	page.AddCharts(graph.Line("売上個数"))
	page.PageTitle = "サンプル"

	f, err := os.Create("./html/bar.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
