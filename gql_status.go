package gqldefs

type GQLStatus struct {
	Status      string
	Description string
}

func NewGQLStatus(status StatusSubclass) GQLStatus {
	var t string
	switch status.StatusClass().Category() {
	case "S", "N", "X":
		if status.Subcode() == "000" {
			t = status.StatusClass().Description()
		} else {
			t = status.StatusClass().Description() + " - " + status.Description()
		}
	default:
		t = status.Description()
	}
	switch status.StatusClass().Category() {
	case "X":
		return GQLStatus{Status: status.Code() + status.Subcode(), Description: "error: " + t}
	case "W":
		return GQLStatus{Status: status.Code() + status.Subcode(), Description: "warn: " + t}
	case "S", "N":
		return GQLStatus{Status: status.Code() + status.Subcode(), Description: "note: " + t}
	case "I":
		return GQLStatus{Status: status.Code() + status.Subcode(), Description: "info: " + t}
	default:
		return GQLStatus{}
	}
}
