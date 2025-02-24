package vcf

import "fmt"

const (
	PhoneTypeCell  PhoneType = "CELL"  // Mobile/Cell phone
	PhoneTypeHome  PhoneType = "HOME"  // Home phone
	PhoneTypeWork  PhoneType = "WORK"  // Work phone
	PhoneTypeFax   PhoneType = "FAX"   // Fax number (general)
	PhoneTypePager PhoneType = "PAGER" // Pager number
	PhoneTypeVoice PhoneType = "VOICE" // Voice telephone (default)
	PhoneTypeText  PhoneType = "TEXT"  // Text messaging (SMS)
	PhoneTypeVideo PhoneType = "VIDEO" // Video call-enabled phone
	PhoneTypeBBS   PhoneType = "BBS"   // Bulletin Board System (rarely used)
	PhoneTypeModem PhoneType = "MODEM" // Modem dial-up number
	PhoneTypeMsg   PhoneType = "MSG"   // Messaging service number
	PhoneTypeCar   PhoneType = "CAR"   // Car phone
	PhoneTypeISDN  PhoneType = "ISDN"  // ISDN phone
	PhoneTypePCS   PhoneType = "PCS"   // Personal Communication Services
)

type PhoneType string

func (p PhoneType) Validate() (err error) {
	switch p {
	case PhoneTypeCell:
	case PhoneTypeHome:
	case PhoneTypeWork:
	case PhoneTypeFax:
	case PhoneTypePager:
	case PhoneTypeVoice:
	case PhoneTypeText:
	case PhoneTypeVideo:
	case PhoneTypeBBS:
	case PhoneTypeModem:
	case PhoneTypeMsg:
	case PhoneTypeCar:
	case PhoneTypeISDN:
	case PhoneTypePCS:
	default:
		return fmt.Errorf("phone type of %v is not supported", p)
	}

	return nil
}
