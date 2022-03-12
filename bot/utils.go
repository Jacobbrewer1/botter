package bot

var (
	roleReactionIdCustom = ""
)

const (
	// bot name
	botName = "botter"

	// Verification failed response
	failedVerificationResponse = "You don't have permission to do that"

	// Verification
	verificationJoinDm = "Hey there %v! :wave:\n" +
		"In order to get full access to %v, you need to react to the verification message in the rules channel!\n" +
		"I will let you know when you have successfully done this. If you have any problems feel free to dm a moderator :smiley:"
	verificationsPassedDm = "Congratulations on passing the verification to %v. Now you have full access to the member channels of the server! :smiley:\n" +
		"Please be sure to follow the rules. I don't want to have to mute you :cry:"
	verificationRemovedByUser = "Un-reacting to the verification message will cause you to loose access to the channels of the %v :cry:\n" +
		"In order to get back the role, just re-react to the message :smiley:"

	// Responses
	unknownResponse              = "Unknown command entered"
	roleSuccessResponse          = "<@%v> now has the role `%v`"
	roleVoidSuccessResponse      = "<@%v> no longer has the role `%v`"
	apiResponseCodeText          = "Response code for search `%v` was - %v"
	poweredByGiphyResponse       = "https://giphy.com/embed/dyp02B1LtyGWF6rgsv"
	gifTrendingText              = "Trending"
	messageIsServerInviteDm      = "You are not allowed to send server invites within the server %v!\nI have already deleted the message for you :stuck_out_tongue:"
	failedTimeoutDm              = "You should be timeout now, however the gods are on your side and I can't.\nSo how about you do us a favour an i will let this one slide if you don't post banned words in the server %v anymore aye? :pray:\n[I have already deleted the messages for you :smile:]"
	messageContainedBannedWordDm = "You are not allowed to send a message containing banned words withing the server %v!\nYou now have got a timeout until %v"
	inviteFormatResponse         = "For an invite link use `.invite` otherwise use ```.invite @user```"
	rcrResponse                  = "Custom roles reset for all users"

	// Embed Colours
	redEmbed   = 15158332
	blueEmbed  = 3447003
	greenEmbed = 3066993

	// Emojis
	runReactionEmoji    = "\U000025B6" //"▶"
	faceWithTongueEmoji = "\U0001F61B"
	joyEmoji            = "\U0001F602"
	controllerEmoji     = "\U0001F3AE"
	sleepEmoji          = "\U0001F634"
	redEmoji            = "\U0001F534"
	blueEmoji           = "\U0001F535"

	// Server emojis
	pikaHi = "<a:pikaHi:828930215522729985>"

	// Server channel id's
	regulationsChannel = "811659973082873906"
	memeChannelId      = "811180172446269472"

	// Github values
	jacobGithubLogon     = "Jacobbrewer1"
	issueNotAssignedText = "Issue not assigned"

	// Formula 1 responses
	genericF1Response  = "It's %v at the %v"
	raceDayResponse    = "It's race day at the %v"
	practice           = "Practice"
	qualifying         = "Qualifying"
	guildSportsChannel = "811662158303068221"
)
