package reports

var ModerationReasons = []string{
	"duplicate listing",
	"incorrect course",
	"incorrect college",
	"missing edition",
	"missing author",
	"missing price",
	"unsafe meeting request",
	"spam content",
	"counterfeit material",
	"copyright concern",
	"offensive language",
	"misleading condition",
	"unreachable seller",
	"expired semester",
	"wrong cover image",
	"wrong ISBN",
	"unapproved advertisement",
	"external payment request",
	"personal data exposure",
	"admin review required",
}

var ConditionLabels = map[string]string{
	"new": "New", "like_new": "Like new", "good": "Good", "fair": "Fair", "worn": "Worn", "annotated": "Annotated", "missing_pages": "Missing pages", "unknown": "Condition not supplied",
}

var WorkflowLabels = map[string][]string{
	"registration": {"collect email", "collect name", "collect college", "validate email", "check duplicate", "persist profile", "issue token"},
	"listing":      {"collect title", "collect course", "collect college", "collect price", "validate fields", "persist pending record", "queue admin review", "publish decision"},
	"catalog":      {"read published records", "normalize query", "match course", "match college", "match title", "match price", "sort recent", "paginate results"},
	"trade":        {"load listing", "check availability", "check buyer identity", "collect meeting place", "persist request", "confirm with seller", "complete exchange"},
	"report":       {"load listing", "collect reason", "persist audit", "notify administrator", "review evidence", "resolve case"},
	"reopen":       {"open file", "create buckets", "write entity", "close handle", "open handle", "read entity", "confirm identity"},
}

func WorkflowSteps(name string) []string { return WorkflowLabels[name] }
func ReasonKnown(reason string) bool {
	for _, v := range ModerationReasons {
		if v == reason {
			return true
		}
	}
	return false
}
func ConditionKnown(condition string) bool { _, ok := ConditionLabels[condition]; return ok }
