package key

type FKey uint8

func FKeyFor(fieldname string) FKey {

	// TODO:  very temporary.  Eventually need to build a static map of all field names upon program initialization.
	switch fieldname {
	// NOTE:  FKey = 0 is reserved!  Do not use.
	case "Row":
		return 1
	case "Col":
		return 2
	case "Embodiment":
		return 3
	case "Label":
		return 4
	case "Issued":
		return 5
	case "Status":
		return 6
	default:
		return 255
	}
}

func FieldnameFor(fkey FKey) string {

	switch fkey {
	// NOTE:  FKey = 0 is reserved!  Do not use.
	case 1:
		return "Row"
	case 2:
		return "Col"
	case 3:
		return "Embodiment"
	case 4:
		return "Label"
	case 5:
		return "Issued"
	case 6:
		return "Status"
	default:
		return ""
	}
}
