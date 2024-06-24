package graph

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"golang.org/x/exp/rand"
)

var (
	itemCnt = 7
	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < itemCnt; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCnt; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func Line(name string) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithColorsOpts(opts.Colors{"red"}),
	)
	line.SetXAxis(weeks).
		AddSeries(name, generateLineItems(),
			charts.WithLineChartOpts(opts.LineChart{ShowSymbol: true}), charts.WithLabelOpts(opts.Label{
				Show: true,
			}))

	return line
}

func BarStack() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "日次売上",
			Subtitle: "9/1-9/6",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "日付",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "売上",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "¥ {value}"},
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			// バージョン上がってから全部のShowをtrueにしないといけないっぽい
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "jpg",
					Title: "Anything you want",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					Lang:  []string{"data view", "turn off", "refresh"},
				},
			}},
		),
	)

	bar.SetXAxis(weeks).
		AddSeries("価格", generateBarItems()).
		AddSeries("税金", generateBarItems()).
		AddSeries("軽減税金", generateBarItems()).
		SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
			Stack: "stackA",
		}))

	return bar
}
