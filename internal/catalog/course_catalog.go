package catalog

import "strings"

var CourseDirectory = map[string]string{
	"CS101": "Introduction to Programming", "CS102": "Data Structures", "CS201": "Algorithms", "CS202": "Operating Systems", "CS203": "Databases", "CS204": "Computer Networks", "CS205": "Web Engineering", "CS301": "Distributed Systems", "CS302": "Compilers", "CS303": "Artificial Intelligence", "CS304": "Machine Learning", "CS305": "Information Security", "CS306": "Mobile Development", "CS307": "Cloud Platforms", "CS308": "Human Computer Interaction", "CS309": "Software Testing", "CS310": "Computer Graphics",
	"MATH101": "Calculus I", "MATH102": "Calculus II", "MATH201": "Linear Algebra", "MATH202": "Probability", "MATH203": "Statistics", "MATH204": "Discrete Mathematics", "MATH301": "Numerical Analysis", "MATH302": "Optimization", "MATH303": "Number Theory", "MATH304": "Applied Mathematics",
	"PHYS101": "General Physics", "PHYS102": "Mechanics", "PHYS201": "Electromagnetism", "PHYS202": "Thermodynamics", "PHYS301": "Quantum Physics", "PHYS302": "Optics", "PHYS303": "Materials Physics", "PHYS304": "Modern Physics",
	"CHEM101": "General Chemistry", "CHEM102": "Organic Chemistry", "CHEM201": "Inorganic Chemistry", "CHEM202": "Analytical Chemistry", "CHEM301": "Biochemistry", "CHEM302": "Physical Chemistry", "CHEM303": "Polymer Chemistry",
	"BIO101": "General Biology", "BIO102": "Cell Biology", "BIO201": "Genetics", "BIO202": "Ecology", "BIO203": "Microbiology", "BIO301": "Molecular Biology", "BIO302": "Evolution", "BIO303": "Neuroscience",
	"ENG101": "Academic Writing", "ENG102": "Technical Writing", "ENG201": "World Literature", "ENG202": "Creative Writing", "ENG203": "Linguistics", "ENG301": "Literary Theory", "ENG302": "Translation Studies",
	"HIS101": "World History", "HIS102": "Modern History", "HIS201": "Asian History", "HIS202": "European History", "HIS203": "Public History", "HIS301": "Historical Methods", "HIS302": "Archives and Memory",
	"ART101": "Art Appreciation", "ART102": "Drawing", "ART201": "Painting", "ART202": "Design Basics", "ART203": "Photography", "ART301": "Digital Media", "ART302": "Exhibition Practice",
	"BUS101": "Business Foundations", "BUS102": "Accounting", "BUS201": "Marketing", "BUS202": "Finance", "BUS203": "Operations", "BUS301": "Entrepreneurship", "BUS302": "Strategy", "BUS303": "Business Analytics",
	"LAW101": "Legal Foundations", "LAW102": "Constitutional Law", "LAW201": "Contract Law", "LAW202": "Civil Procedure", "LAW203": "Criminal Law", "LAW301": "Technology Law", "LAW302": "Legal Research",
	"EDU101": "Education Foundations", "EDU102": "Learning Sciences", "EDU201": "Classroom Design", "EDU202": "Assessment", "EDU301": "Education Policy", "EDU302": "Inclusive Teaching",
	"SOC101": "Sociology", "SOC102": "Social Psychology", "SOC201": "Urban Studies", "SOC202": "Research Methods", "SOC301": "Public Policy", "SOC302": "Community Practice",
	"MED101": "Health Sciences", "MED102": "Anatomy", "MED201": "Public Health", "MED202": "Nutrition", "MED301": "Epidemiology", "MED302": "Health Communication",
}

func CourseName(code string) string {
	if v, ok := CourseDirectory[strings.ToUpper(strings.TrimSpace(code))]; ok {
		return v
	}
	return "Unknown course"
}
func KnownCourse(code string) bool {
	_, ok := CourseDirectory[strings.ToUpper(strings.TrimSpace(code))]
	return ok
}
func CourseCodes() []string {
	out := make([]string, 0, len(CourseDirectory))
	for c := range CourseDirectory {
		out = append(out, c)
	}
	return out
}
