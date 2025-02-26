package vcf

import (
	"encoding/json"
	"fmt"
)

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

func (p *PhotoType) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	pt := PhotoType(str)
	if err = pt.Validate(); err != nil {
		return
	}

	*p = pt
	return
}
