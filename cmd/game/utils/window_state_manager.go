package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

type WindowStateManager struct {
	path string
}

type windowStateFile struct {
	Monitors []monitorState `json:"monitors"`
	Monitor  int            `json:"monitor"`
	X        int            `json:"x"`
	Y        int            `json:"y"`
}

type monitorState struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Scale  string `json:"scale"`
}

func NewWindowStateManager(appName string) *WindowStateManager {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return &WindowStateManager{}
	}
	return &WindowStateManager{
		path: filepath.Join(configDir, appName, "window-state.json"),
	}
}

// 前回と同じモニター構成の場合、ウィンドウ位置を復元する
func (s *WindowStateManager) RestoreIfNeeded() {
	if s == nil || s.path == "" {
		return
	}

	data, err := os.ReadFile(s.path)
	if err != nil {
		return
	}

	var state windowStateFile
	if err := json.Unmarshal(data, &state); err != nil {
		return
	}
	monitors := ebiten.AppendMonitors(nil)
	if !sameMonitors(state.Monitors, monitorStates(monitors)) {
		return
	}
	if 0 <= state.Monitor && state.Monitor < len(monitors) {
		ebiten.SetMonitor(monitors[state.Monitor])
	}

	ebiten.SetWindowPosition(state.X, state.Y)
}

func (s *WindowStateManager) Save() error {
	if s == nil || s.path == "" {
		return nil
	}

	monitors := ebiten.AppendMonitors(nil)
	x, y := ebiten.WindowPosition()
	state := windowStateFile{
		Monitors: monitorStates(monitors),
		Monitor:  monitorIndex(monitors, ebiten.Monitor()),
		X:        x,
		Y:        y,
	}
	data, err := json.MarshalIndent(state, "", "\t")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0644)
}

func monitorStates(monitors []*ebiten.MonitorType) []monitorState {
	states := make([]monitorState, 0, len(monitors))
	for _, monitor := range monitors {
		width, height := monitor.Size()
		states = append(states, monitorState{
			Name:   monitor.Name(),
			Width:  width,
			Height: height,
			Scale:  fmt.Sprintf("%.3f", monitor.DeviceScaleFactor()),
		})
	}
	return states
}

func monitorIndex(monitors []*ebiten.MonitorType, current *ebiten.MonitorType) int {
	for i, monitor := range monitors {
		if monitor == current {
			return i
		}
	}
	return 0
}

func sameMonitors(a, b []monitorState) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
