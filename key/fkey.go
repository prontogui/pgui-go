package key

type FKey uint8

func FKeyFor(fieldname string) FKey {

	// TODO:  very temporary.  Eventually need to build a static map of all field names upon program initialization.
	switch fieldname {
	case "Row":
		return 0
	case "Col":
		return 1
	case "Embodiment":
		return 2
	case "Label":
		return 3
	case "Issued":
		return 4
	case "Status":
		return 5
	default:
		return 255
	}
}

func FieldnameFor(fkey FKey) string {

	switch fkey {
	case 0:
		return "Row"
	case 1:
		return "Col"
	case 2:
		return "Embodiment"
	case 3:
		return "Label"
	case 4:
		return "Issued"
	case 5:
		return "Status"
	default:
		return ""
	}
}
