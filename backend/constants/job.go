package constants

// WorkMode constants
const (
	WorkModeOnsite   = "onsite"
	WorkModeRemote   = "remote"
	WorkModeHybrid   = "hybrid"
	WorkModeFlexible = "flexible"
)

// WorkType constants
const (
	WorkTypeFullTime   = "full_time"
	WorkTypePartTime   = "part_time"
	WorkTypeContract   = "contract"
	WorkTypeFreelance  = "freelance"
	WorkTypeInternship = "internship"
	WorkTypeTemporary  = "temporary"
)

// InterviewMode constants
const (
	InterviewModeOnline  = "online"
	InterviewModeOffline = "offline"
	InterviewModeHybrid  = "hybrid"
	InterviewModePhone   = "phone"
	InterviewModeVideo   = "video"
)

// PaySchedule constants
const (
	PayScheduleMonthly  = "monthly"
	PayScheduleBiWeekly = "bi_weekly"
	PayScheduleWeekly   = "weekly"
	PayScheduleDaily    = "daily"
	PayScheduleHourly   = "hourly"
	PayScheduleAnnually = "annually"
)

// Currency constants
const (
	CurrencyINR = "INR"
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyGBP = "GBP"
	CurrencyCAD = "CAD"
	CurrencyAUD = "AUD"
	CurrencyJPY = "JPY"
	CurrencySGD = "SGD"
)

// JobSkillType constants
const (
	SkillTypeRequired   = "required"
	SkillTypeGoodToHave = "good_to_have"
	SkillTypeOptional   = "optional"
	SkillTypePreferred  = "preferred"
)

// Job Source constants
const (
	SourceLinkedIn    = "linkedin"
	SourceNaukri      = "naukri"
	SourceIndeed      = "indeed"
	SourceGlassdoor   = "glassdoor"
	SourceCompanyPage = "company_page"
	SourceReferral    = "referral"
	SourceInternal    = "internal"
	SourceJobBoard    = "job_board"
	SourceRecruiter   = "recruiter"
)

// Education Level constants
const (
	EducationHighSchool   = "high_school"
	EducationDiploma      = "diploma"
	EducationBachelor     = "bachelor"
	EducationMaster       = "master"
	EducationPhD          = "phd"
	EducationProfessional = "professional"
	EducationNone         = "none"
	EducationAny          = "any"
)

// Contract Duration constants
const (
	ContractDuration3Months   = "3_months"
	ContractDuration6Months   = "6_months"
	ContractDuration1Year     = "1_year"
	ContractDuration2Years    = "2_years"
	ContractDurationPermanent = "permanent"
	ContractDurationProject   = "project_based"
)

// Bond Period constants
const (
	BondPeriodNone    = "none"
	BondPeriod6Months = "6_months"
	BondPeriod1Year   = "1_year"
	BondPeriod2Years  = "2_years"
	BondPeriod3Years  = "3_years"
)

// Shift Timings constants
const (
	ShiftDay        = "day"
	ShiftNight      = "night"
	ShiftEvening    = "evening"
	ShiftRotational = "rotational"
	ShiftFlexible   = "flexible"
	ShiftUS         = "us_hours"
	ShiftUK         = "uk_hours"
	ShiftGeneral    = "general"
)

// Validation slices for easy validation
var (
	ValidWorkModes = []string{
		WorkModeOnsite, WorkModeRemote, WorkModeHybrid, WorkModeFlexible,
	}

	ValidWorkTypes = []string{
		WorkTypeFullTime, WorkTypePartTime, WorkTypeContract,
		WorkTypeFreelance, WorkTypeInternship, WorkTypeTemporary,
	}

	ValidInterviewModes = []string{
		InterviewModeOnline, InterviewModeOffline, InterviewModeHybrid,
		InterviewModePhone, InterviewModeVideo,
	}

	ValidPaySchedules = []string{
		PayScheduleMonthly, PayScheduleBiWeekly, PayScheduleWeekly,
		PayScheduleDaily, PayScheduleHourly, PayScheduleAnnually,
	}

	ValidCurrencies = []string{
		CurrencyINR, CurrencyUSD, CurrencyEUR, CurrencyGBP,
		CurrencyCAD, CurrencyAUD, CurrencyJPY, CurrencySGD,
	}

	ValidSkillTypes = []string{
		SkillTypeRequired, SkillTypeGoodToHave, SkillTypeOptional, SkillTypePreferred,
	}

	ValidSources = []string{
		SourceLinkedIn, SourceNaukri, SourceIndeed, SourceGlassdoor,
		SourceCompanyPage, SourceReferral, SourceInternal, SourceJobBoard, SourceRecruiter,
	}

	ValidEducationLevels = []string{
		EducationHighSchool, EducationDiploma, EducationBachelor,
		EducationMaster, EducationPhD, EducationProfessional, EducationNone, EducationAny,
	}

	ValidContractDurations = []string{
		ContractDuration3Months, ContractDuration6Months, ContractDuration1Year,
		ContractDuration2Years, ContractDurationPermanent, ContractDurationProject,
	}

	ValidBondPeriods = []string{
		BondPeriodNone, BondPeriod6Months, BondPeriod1Year, BondPeriod2Years, BondPeriod3Years,
	}

	ValidShiftTimings = []string{
		ShiftDay, ShiftNight, ShiftEvening, ShiftRotational,
		ShiftFlexible, ShiftUS, ShiftUK, ShiftGeneral,
	}
)

// Helper functions for validation
func IsValidWorkMode(mode string) bool {
	return contains(ValidWorkModes, mode)
}

func IsValidWorkType(workType string) bool {
	return contains(ValidWorkTypes, workType)
}

func IsValidInterviewMode(mode string) bool {
	return contains(ValidInterviewModes, mode)
}

func IsValidPaySchedule(schedule string) bool {
	return contains(ValidPaySchedules, schedule)
}

func IsValidCurrency(currency string) bool {
	return contains(ValidCurrencies, currency)
}

func IsValidSkillType(skillType string) bool {
	return contains(ValidSkillTypes, skillType)
}

func IsValidSource(source string) bool {
	return contains(ValidSources, source)
}

func IsValidEducationLevel(education string) bool {
	return contains(ValidEducationLevels, education)
}

func IsValidContractDuration(duration string) bool {
	return contains(ValidContractDurations, duration)
}

func IsValidBondPeriod(period string) bool {
	return contains(ValidBondPeriods, period)
}

func IsValidShiftTiming(timing string) bool {
	return contains(ValidShiftTimings, timing)
}

// Helper function to check if slice contains a value
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
