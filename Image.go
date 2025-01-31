package toolbox

import (
	"image"
	"net/http"
	"os"

	"github.com/mattn/go-sixel"
)

func PrintImageURL(url string) {
	// Get the image
	resp, err := http.Get(url)
	EoE(err, "Error getting image: ")
	defer resp.Body.Close()

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	EoE(err, "Error decoding image: ")

	// Create a Sixel encoder
	enc := sixel.NewEncoder(os.Stdout)

	// Encode the image
	err = enc.Encode(img)
	EoE(err, "Error encoding image: ")
}
