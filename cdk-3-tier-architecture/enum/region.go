package enum

type AZEnum int

const (
	AZ_AP_NORTHEAST_1A AZEnum = iota
	AZ_AP_NORTHEAST_1B
	AZ_AP_NORTHEAST_1C
)

func (az AZEnum) String() string {
	switch az {
	case AZ_AP_NORTHEAST_1A:
		return "ap-northeast-1a"
	case AZ_AP_NORTHEAST_1B:
		return "ap-northeast-1b"
	case AZ_AP_NORTHEAST_1C:
		return "ap-northeast-1c"
	default:
		panic("Not supported AZEnum Type")
	}
}
