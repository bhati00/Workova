package constant

type JobType int

const (
	JobTypeFullTime JobType = iota + 1
	JobTypePartTime
	JobTypeContract
	JobTypeInternship
	JobTypeTemporary
)

type WorkMode int

const (
	WorkModeRemote WorkMode = iota + 1
	WorkModeOnsite
	WorkModeHybrid
)

type ExperienceLevel int

const (
	ExperienceLevelEntry ExperienceLevel = iota + 1
	ExperienceLevelMid
	ExperienceLevelSenior
	ExperienceLevelLead
	ExperienceLevelExecutive
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyINR Currency = "INR"
)
