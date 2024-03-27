package key

type FKey uint8

const (
	INVALID_FIELDNAME = 255
	INVALID_FKEY      = ""
)

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
	case "Choices":
		return 7
	case "Data":
		return 8
	case "ListItems":
		return 9
	case "Rows":
		return 10
	default:
		return INVALID_FIELDNAME
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
	case 7:
		return "Choices"
	case 8:
		return "Data"
	case 9:
		return "ListItems"
	case 10:
		return "Rows"
	default:
		return INVALID_FKEY
	}
}
