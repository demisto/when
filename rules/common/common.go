package common

import "github.com/demisto/when/rules"

var All = []rules.Rule{
	SlashDMY(rules.Override),
}
