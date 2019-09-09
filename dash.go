// +build ignore

package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Mock Data
	data := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}

    // Header Row
	header := widgets.NewParagraph()
	header.Text = "iRobot Create 2 Dashboard"

	// Controls
	controls := widgets.NewList()
	controls.Title = "Command & Control"
	controls.Rows = []string{
		"[w] Move forwards",
		"[a] Rotate counter-clockwise",
		"[s] Move backwards",
		"[d] Rotate clockwise",
	}
	controls.TextStyle = ui.NewStyle(ui.ColorYellow)
	controls.WrapText = false
	controls.TitleStyle.Fg = ui.ColorCyan

	// Wheel Telemetry
	wheel_header := widgets.NewParagraph()
	wheel_header.Title = "Wheel Telemetry"
	wheel_header.TitleStyle.Fg = ui.ColorCyan

	// Battery 
	battery := widgets.NewParagraph()
	battery.Title = "Battery"
	battery.TitleStyle.Fg = ui.ColorCyan

	// Sensors
	sensor_lc := widgets.NewGauge()
	sensor_lc.Title = "Left Cliff Sensor"
	sensor_lc.Percent = 0
	sensor_lc.BarColor = ui.ColorRed
	sensor_lc.BorderStyle.Fg = ui.ColorWhite
	sensor_lc.TitleStyle.Fg = ui.ColorCyan

	sensor_rc := widgets.NewGauge()
	sensor_rc.Title = "Right Cliff Sensor"
	sensor_rc.Percent = 0
	sensor_rc.BarColor = ui.ColorRed
	sensor_rc.BorderStyle.Fg = ui.ColorWhite
	sensor_rc.TitleStyle.Fg = ui.ColorCyan

	sensor_lfc := widgets.NewGauge()
	sensor_lfc.Title = "Left Front Cliff Sensor"
	sensor_lfc.Percent = 0
	sensor_lfc.BarColor = ui.ColorRed
	sensor_lfc.BorderStyle.Fg = ui.ColorWhite
	sensor_lfc.TitleStyle.Fg = ui.ColorCyan

	sensor_rfc := widgets.NewGauge()
	sensor_rfc.Title = "Right Front Cliff Sensor"
	sensor_rfc.Percent = 0
	sensor_rfc.BarColor = ui.ColorRed
	sensor_rfc.BorderStyle.Fg = ui.ColorWhite
	sensor_rfc.TitleStyle.Fg = ui.ColorCyan

	sensor_bumpers := widgets.NewParagraph()
	sensor_bumpers.Title = "Bumpers"

	sensor_ultrasonic := widgets.NewParagraph()
	sensor_ultrasonic.Title = "Ultrasonic"

	// Bumpers
	bumper_l := widgets.NewGauge()
	bumper_l.Title = "Left Bumper"
	bumper_l.TitleStyle.Fg = ui.ColorCyan

	bumper_fl := widgets.NewGauge()
	bumper_fl.Title = "Front Left Bumper"
	bumper_fl.TitleStyle.Fg = ui.ColorCyan

	bumper_cl := widgets.NewGauge()
	bumper_cl.Title = "Ctr Left Bumper"
	bumper_cl.TitleStyle.Fg = ui.ColorCyan

	bumper_r := widgets.NewGauge()
	bumper_r.Title = "Right Bumper"
	bumper_r.TitleStyle.Fg = ui.ColorCyan

	bumper_fr := widgets.NewGauge()
	bumper_fr.Title = "Front Right Bumper"
	bumper_fr.TitleStyle.Fg = ui.ColorCyan

	bumper_cr := widgets.NewGauge()
	bumper_cr.Title = "Ctr Right Bumper"
	bumper_cr.TitleStyle.Fg = ui.ColorCyan

	// Ultrasonic Sensor
	ultra := widgets.NewSparkline()
	ultra.Data = data[3:]
	ultra.LineColor = ui.ColorGreen

	// Ultrasonic Group
	ultra_group := widgets.NewSparklineGroup(ultra)
	ultra_group.Title = "Ultrasonic Sensor"
	ultra_group.TitleStyle.Fg = ui.ColorCyan

	// Render the Widgets in a Grid
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/3, controls),
			ui.NewCol(1.0/3, wheel_header),
			ui.NewCol(1.0/3, battery),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/3, ultra_group),
			ui.NewCol(2.0/3,
				ui.NewRow(1.0/2,
					ui.NewCol(1.0/4, sensor_lc),
					ui.NewCol(1.0/4, sensor_lfc),
					ui.NewCol(1.0/4, sensor_rfc),
					ui.NewCol(1.0/4, sensor_rc),
				),
				ui.NewRow(1.0/2,
					ui.NewCol(1.0/6, bumper_l),
					ui.NewCol(1.0/6, bumper_fl),
					ui.NewCol(1.0/6, bumper_cl),
					ui.NewCol(1.0/6, bumper_cr),
					ui.NewCol(1.0/6, bumper_fr),
					ui.NewCol(1.0/6, bumper_r),
				),
			),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/2, controls),
			ui.NewCol(1.0/2, wheel_header),
		),
	)

	ui.Render(grid)


	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "q" {
				break
			}
		}
	}


}