package prepare

type DataMode string

const (
	ByCountry DataMode = "by_country"
	ByState   DataMode = "by_state"
)

func (mode DataMode) SubFolderPath() string {
	switch mode {
	case ByCountry:
		return "country"
	case ByState:
		return "state"
	default:
		return ""
	}
}

type PrepareData struct {
	DataConfig
	Mode DataMode
}
