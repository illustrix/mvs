package auto

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"os/exec"
	"path/filepath"
)

type Auto struct {
	wd string
}

func NewAuto() *Auto {
	return &Auto{}
}

func (a *Auto) Init() error {
	wd, err := getWorkingDir()
	if err != nil {
		return err
	}
	a.wd = wd
	return nil
}

func (a *Auto) GrabScreen() (image.Image, error) {
	if a.wd == "" {
		if err := a.Init(); err != nil {
			return nil, err
		}
	}

	buf := make([]byte, 100)
	stderr := bytes.NewBuffer(buf)
	buf2 := make([]byte, 100)
	stdout := bytes.NewBuffer(buf2)
	cmd := exec.Command("screencapture", "-x", "tmp/screenshot.png")
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Dir = a.wd
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	imgFile, err := os.Open(filepath.Join(a.wd, "tmp/screenshot.png"))
	if err != nil {
		return nil, fmt.Errorf("failed to open screenshot file: %w", err)
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (a *Auto) AnalyticsColor(img image.Image) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	colorCount := make(map[string]int)
	for y := range height {
		for x := range width {
			r, g, b, _ := img.At(x, y).RGBA()
			colorKey := fmt.Sprintf("%d_%d_%d", r>>8, g>>8, b>>8)
			colorCount[colorKey]++
		}
	}

	fmt.Printf("Color analytics: %v\n", colorCount)
}
