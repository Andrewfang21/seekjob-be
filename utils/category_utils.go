package utils

import "fmt"

/*
TO BE DELETED

var ADZUNA_CATEGORIES = map[string]string{
	"accounting-finance-jobs":       "Accounting & Finance Jobs",
	"engineering-jobs":              "Engineering Jobs",
	"it-jobs":                       "IT Jobs",
	"pr-advertising-marketing-jobs": "PR, Advertising & Marketing Jobs",
}

var REMOTIVE_JOBS_CATEGORIES = map[string]string{
	"software-dev":    "Software Development",
	"marketing-sales": "Marketing / Sales",
	"product":         "Product",
}
*/

var adzunaCategories = []string{
	"accounting-finance-jobs",
	"engineering-jobs",
	"it-jobs",
	"pr-advertising-marketing-jobs",
}

var reedCategories = []string{
	"Accounting",
	"Marketing",
	"Project Manager",
	"Software Developer",
	"Software Engineering",
}

var remotiveCategories = []string{
	"software-dev",
	"marketing-sales",
	"product",
}

var theMuseCategories = []string{
	"Business & Strategy",
	"Data Science",
	"Engineering",
	"Finance",
	"Marketing & PR",
	"Project & Product Management",
}

// GetCategories returns list of categories according to the source
func GetCategories(source string) ([]string, error) {
	list := []string{}
	switch source {
	case "ADZUNA":
		list = append(list, adzunaCategories...)
		break
	case "REED":
		list = append(list, reedCategories...)
		break
	case "REMOTIVE":
		list = append(list, remotiveCategories...)
		break
	case "THEMUSE":
		list = append(list, theMuseCategories...)
		break
	default:
		return nil, fmt.Errorf("Source %s does not exist", source)
	}
	return list, nil
}
