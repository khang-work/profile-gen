package sample

import "github.com/khang-work/profile-gen/prof"

// GenerateSample generates a sample Profile with generic information.
func GenerateSample() prof.Profile {
	return prof.Profile{
		PersonalInfo: prof.PersonalInfo{
			Name: "John Doe",
			Address: prof.Address{
				Street: "123 Main St",
				City:   "Somewhere",
				State:  "CA",
				Zip:    "90001",
			},
			PhoneNumber: "+1 (555) 123-4567",
			Email:       "johndoe@example.com",
			LinkedIn:    "https://linkedin.com/in/johndoe",
			GitHub:      "https://github.com/johndoe",
		},
		ProfessionalSummary: "Experienced software developer with a passion for building innovative solutions. Skilled in multiple programming languages and frameworks with a strong background in software architecture and problem-solving.",
		TechnicalSkills: prof.TechnicalSkills{
			Languages:            []string{"Go", "Python", "JavaScript"},
			Frameworks:           []string{"React", "Django", "Spring Boot"},
			ToolsAndTechnologies: []string{"Docker", "Kubernetes", "CI/CD"},
			Databases:            []string{"PostgreSQL", "MySQL", "MongoDB"},
			DevelopmentPractices: []string{"Agile", "Test-Driven Development", "Pair Programming"},
			Other:                []string{"Cloud Computing", "Machine Learning"},
		},
		Education: []prof.Education{
			{
				Degree:         "Bachelor of Science in Computer Science",
				University:     "University of Somewhere",
				GraduationDate: "May 2024",
				Coursework:     []string{"Data Structures", "Algorithms", "Database Systems", "Operating Systems"},
				Awards:         []string{"Dean's List", "Best Capstone Project"},
				Minors:         []string{"Mathematics"},
				GPA:            "4.0",
			},
		},
		ProfessionalExperience: []prof.ProfessionalExperience{
			{
				JobTitle:  "Software Engineer",
				Company:   "TechCorp Inc.",
				Location:  "San Francisco, CA",
				StartDate: "June 2020",
				EndDate:   "Present",
				Responsibilities: []string{
					"Developed and maintained backend services using Go and Python",
					"Collaborated with front-end developers to integrate React-based user interfaces",
					"Improved system performance by optimizing database queries",
					"Implemented microservices architecture in the cloud environment",
				},
			},
		},
		Certifications: []prof.Certification{
			{
				Name: "AWS Certified Solutions Architect",
				Year: "2023",
			},
			{
				Name: "Certified Kubernetes Administrator",
				Year: "2022",
			},
		},
		AdditionalInfo: prof.AdditionalInfo{
			Languages: []string{"English", "Spanish"},
			Volunteer: "Volunteer Software Developer at Local Nonprofit",
			Interests: []string{"AI Research", "Open Source Contributions", "Cycling"},
		},
	}
}
