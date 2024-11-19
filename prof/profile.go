package prof

type Profile struct {
	PersonalInfo           PersonalInfo             `json:"personal_info"`
	ProfessionalSummary    string                   `json:"professional_summary"`
	TechnicalSkills        TechnicalSkills          `json:"technical_skills"`
	Education              []Education              `json:"education"`
	ProfessionalExperience []ProfessionalExperience `json:"professional_experience"`
	Certifications         []Certification          `json:"certifications"`
	AdditionalInfo         AdditionalInfo           `json:"additional_info"`
}

// PersonalInfo represents the personal details.
type PersonalInfo struct {
	Name        string  `json:"name"`
	Address     Address `json:"address"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	LinkedIn    string  `json:"linkedin"`
	GitHub      string  `json:"github"`
}

// TechnicalSkills represents the technical skills section.
type TechnicalSkills struct {
	Languages            []string `json:"languages"`
	Frameworks           []string `json:"frameworks"`
	ToolsAndTechnologies []string `json:"tools_and_technologies"`
	Databases            []string `json:"databases"`
	DevelopmentPractices []string `json:"development_practices"`
	Other                []string `json:"other"`
}

// Education represents educational background.
type Education struct {
	Degree         string   `json:"degree"`
	University     string   `json:"university"`
	GraduationDate string   `json:"graduation_date"`
	Coursework     []string `json:"coursework"`
	Awards         []string `json:"awards"`
	Minors         []string `json:"minors"`
	GPA            string   `json:"gpa"`
}

// ProfessionalExperience represents each job experience.
type ProfessionalExperience struct {
	JobTitle         string   `json:"job_title"`
	Company          string   `json:"company"`
	Location         string   `json:"location"`
	StartDate        string   `json:"start_date"`
	EndDate          string   `json:"end_date"`
	Responsibilities []string `json:"responsibilities"`
}

// Certification represents certifications obtained.
type Certification struct {
	Name string `json:"name"`
	Year string `json:"year"`
}

// AdditionalInfo represents extra information such as language skills, interests, etc.
type AdditionalInfo struct {
	Languages []string `json:"languages"`
	Volunteer string   `json:"volunteer"`
	Interests []string `json:"interests"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}
