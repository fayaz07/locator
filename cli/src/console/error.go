package console

const (
	issuesUrl = "https://github.com/fayaz07/locator/issues/new"
)

func RequestLogIssue(msg string) {
	Error("Error reading input, please raise an issue at " + issuesUrl)
}
