package imageprocessing

import "testing"

func TestNormalFlow(t *testing.T) {
	_, err := Crop("large_jpg.jpg")
	if err != nil {
		t.Error()
	}
	_, err2 := Crop("large_png.png")
	if err2 != nil {
		t.Error()
	}
}

func TestImageToSmoll(t *testing.T) {
	_, err := Crop("to_smoll")
	if err == nil {
		t.Error()
	}
}

func TestWeirdRatio(t *testing.T) {
	_, err := Crop("weird_img.jpeg")
	if err != nil {
		t.Error()
	}
}
