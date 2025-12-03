package helpers

type LanguageEdge struct {
	Size int
	Node struct {
		Name  string
		Color string
	}
}

type Repository struct {
	StargazerCount int
	ForkCount      int
	Languages      struct {
		Edges []LanguageEdge
	} `graphql:"languages(first: 10)"`
}

type UserStats struct {
	User struct {
		ContributionsCollection struct {
			TotalCommitContributions            int
			TotalIssueContributions             int
			TotalPullRequestContributions       int
			TotalPullRequestReviewContributions int
			ContributionCalendar                struct {
				TotalContributions int
			}
		}
		Repositories struct {
			TotalCount int
			Nodes      []Repository
		} `graphql:"repositories(first: 100, ownerAffiliations: OWNER, privacy: PUBLIC)"`
	} `graphql:"user(login: $username)"`
}

type LanguageStat struct {
	Name       string
	Size       int
	Percentage float64
	Color      string
}
