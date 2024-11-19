package txt

import (
	"fmt"
	"strings"

	"github.com/khang-work/profile-gen/prof"
)

// ProfileToPlainTextResume generates a plain text resume from the given profile.
func ProfileToPlainTextResume(p *prof.Profile, txtLineLength int) string {
	sb := &strings.Builder{}

	// Personal Information.
	fmt.Fprintln(sb, p.PersonalInfo.Name)
	fmt.Fprintf(sb, "%s, %s, %s %s\n", p.PersonalInfo.Address.Street, p.PersonalInfo.Address.City, p.PersonalInfo.Address.State, p.PersonalInfo.Address.Zip)
	fmt.Fprintln(sb, p.PersonalInfo.PhoneNumber)
	fmt.Fprintln(sb, p.PersonalInfo.Email)
	fmt.Fprintln(sb)
	fmt.Fprintf(sb, "LinkedIn: %s\n", p.PersonalInfo.LinkedIn)
	fmt.Fprintf(sb, "GitHub: %s\n", p.PersonalInfo.GitHub)
	fmt.Fprintln(sb)
	fmt.Fprintln(sb)

	// Professional Summary.
	if p.ProfessionalSummary != "" {
		fmt.Fprintf(sb, "PROFESSIONAL SUMMARY\n\n")
		fmt.Fprintln(sb, p.ProfessionalSummary)
		fmt.Fprintln(sb)
		fmt.Fprintln(sb)
	}

	// Technical Skills.
	if len(p.TechnicalSkills.Languages) > 0 ||
		len(p.TechnicalSkills.Frameworks) > 0 ||
		len(p.TechnicalSkills.ToolsAndTechnologies) > 0 ||
		len(p.TechnicalSkills.Databases) > 0 ||
		len(p.TechnicalSkills.DevelopmentPractices) > 0 ||
		len(p.TechnicalSkills.Other) > 0 {
		fmt.Fprintln(sb, "TECHNICAL SKILLS")
		fmt.Fprintln(sb)

		if len(p.TechnicalSkills.Languages) > 0 {
			fmt.Fprintf(sb, "Languages: %s\n", strings.Join(p.TechnicalSkills.Languages, " | "))
		}

		if len(p.TechnicalSkills.Frameworks) > 0 {
			fmt.Fprintf(sb, "Frameworks: %s\n", strings.Join(p.TechnicalSkills.Frameworks, " | "))
		}

		if len(p.TechnicalSkills.ToolsAndTechnologies) > 0 {
			fmt.Fprintf(sb, "Tools & Technologies: %s\n", strings.Join(p.TechnicalSkills.ToolsAndTechnologies, " | "))
		}

		if len(p.TechnicalSkills.Databases) > 0 {
			fmt.Fprintf(sb, "Databases: %s\n", strings.Join(p.TechnicalSkills.Databases, " | "))
		}

		if len(p.TechnicalSkills.DevelopmentPractices) > 0 {
			fmt.Fprintf(sb, "Development Practices: %s\n", strings.Join(p.TechnicalSkills.DevelopmentPractices, " | "))
		}

		if len(p.TechnicalSkills.Other) > 0 {
			fmt.Fprintf(sb, "Other Skills: %s\n", strings.Join(p.TechnicalSkills.Other, " | "))
		}

		fmt.Fprintln(sb)
		fmt.Fprintln(sb)
	}

	// Education.
	if len(p.Education) > 0 {
		fmt.Fprintln(sb, "EDUCATION")
		fmt.Fprintln(sb)

		for _, edu := range p.Education {
			fmt.Fprintf(sb, "%s in %s\n", edu.Degree, edu.University)
			fmt.Fprintf(sb, "Graduation Date: %s\n", edu.GraduationDate)
			fmt.Fprintf(sb, "GPA: %s\n", edu.GPA)

			if len(edu.Minors) > 0 {
				fmt.Fprintf(sb, "Minors: %s\n", strings.Join(edu.Minors, " | "))
			}

			if len(edu.Awards) > 0 {
				fmt.Fprintf(sb, "Awards: %s\n", strings.Join(edu.Awards, " | "))
			}

			fmt.Fprintf(sb, "Relevant Coursework: %s\n", strings.Join(edu.Coursework, " | "))
			fmt.Fprintln(sb)
		}

		fmt.Fprintln(sb)
	}

	// Professional Experience.
	if len(p.ProfessionalExperience) > 0 {
		fmt.Fprintln(sb, "PROFESSIONAL EXPERIENCE")
		fmt.Fprintln(sb)

		for _, exp := range p.ProfessionalExperience {
			fmt.Fprintf(sb, "%s at %s (%s)\n", exp.JobTitle, exp.Company, exp.Location)
			fmt.Fprintf(sb, "From: %s To: %s\n", exp.StartDate, exp.EndDate)
			fmt.Fprintln(sb, "Responsibilities:")

			for _, responsibility := range exp.Responsibilities {
				fmt.Fprintf(sb, " - %s\n", responsibility)
			}

			fmt.Fprintln(sb)
		}

		fmt.Fprintln(sb)
	}

	// Certifications.
	if len(p.Certifications) > 0 {
		fmt.Fprintln(sb, "CERTIFICATIONS")
		fmt.Fprintln(sb)

		for _, cert := range p.Certifications {
			fmt.Fprintf(sb, "%s (%s)\n", cert.Name, cert.Year)
		}

		fmt.Fprintln(sb)
		fmt.Fprintln(sb)
	}

	// Additional Information.
	if len(p.AdditionalInfo.Languages) > 0 ||
		p.AdditionalInfo.Volunteer != "" ||
		len(p.AdditionalInfo.Interests) > 0 {
		fmt.Fprintln(sb, "ADDITIONAL INFORMATION")
		fmt.Fprintln(sb)

		if len(p.AdditionalInfo.Languages) > 0 {
			fmt.Fprintf(sb, "Languages: %s\n", strings.Join(p.AdditionalInfo.Languages, ", "))
		}

		if p.AdditionalInfo.Volunteer != "" {
			fmt.Fprintf(sb, "Volunteer Experience: %s\n", p.AdditionalInfo.Volunteer)
		}

		if len(p.AdditionalInfo.Interests) > 0 {
			fmt.Fprintf(sb, "Interests: %s\n", strings.Join(p.AdditionalInfo.Interests, ", "))
		}

		fmt.Fprintln(sb)
	}

	return wordWrap(sb.String(), txtLineLength)
}
