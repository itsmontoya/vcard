package vcf

import "fmt"

const (
	PhotoTypeJPEG PhotoType = "JPEG" // image/jpeg
	PhotoTypePNG  PhotoType = "PNG"  // image/png
	PhotoTypeGIF  PhotoType = "GIF"  // image/gif
	PhotoTypeBMP  PhotoType = "BMP"  // image/bmp (vCard 4.0)
	PhotoTypeSVG  PhotoType = "SVG"  // image/svg+xml (vCard 4.0)
)

type PhotoType string

func (p PhotoType) Validate() (err error) {
	switch p {
	case PhotoTypeJPEG:
	case PhotoTypePNG:
	case PhotoTypeGIF:
	case PhotoTypeBMP:
	case PhotoTypeSVG:
	default:
		return fmt.Errorf("photo type of %v is not supported", p)
	}

	return nil
}
