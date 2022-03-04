package bot

var roles = []role{everyone, botDeveloper, blueRole, redRole, minecraft, member, squad, moderator, botter, boss}

// Server Roles
// These roles should only be used for verifier and nothing else
type role struct {
	id    string
	level int
}

var everyone = role{
	id:    "811180171830886410",
	level: 0,
}

var botDeveloper = role{
	id:    "945426342407704659",
	level: 1,
}

var blueRole = role{
	id:    "936644834939252756",
	level: 2,
}

var redRole = role{
	id:    "936644674641346620",
	level: 3,
}

var minecraft = role{
	id:    "937774273827848272",
	level: 4,
}

var member = role{
	id:    "815309191903313990",
	level: 5,
}

var squad = role{
	id:    "811566288794288170",
	level: 6,
}

var moderator = role{
	id:    "811258877793533962",
	level: 7,
}

var botter = role{
	id:    "811352349712842782",
	level: 8,
}

var boss = role{
	id:    "811259138352611410",
	level: 9,
}
