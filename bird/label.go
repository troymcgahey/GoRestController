package bird

// ResolveLabel maps a selection string to a display label per feature spec:
// exactly "Eagle" -> "American Bald Eagle"; any other value -> "Chickadee".
func ResolveLabel(selection string) string {
	if selection == "Eagle" {
		return "American Bald Eagle"
	}
	return "Chickadee"
}
