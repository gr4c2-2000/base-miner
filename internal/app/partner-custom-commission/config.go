package partnercustomcommission

type PartnerConfig struct {
	Partner Partner
}
type Partner struct {
	ProgramName     string
	CommissionType  int
	CommissionValue int
	AffiliantId     int
	DeliveryEmail   []string
	RunOn           int
}
