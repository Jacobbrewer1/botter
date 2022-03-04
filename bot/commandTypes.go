package bot

import "github.com/Jacobbrewer1/botter/helper"

var commands = []command{help, actions, verification, ping, invite, serverAdTemplateCommand, hello, hi, hey, laugh, bossCommand, mrsBossCommand, minecraftHomeCoordinates,
	minecraftBrewing, mcPigRide, roleReactCustom, poll, gifCommand, stickerCommand, resetCustomRoles, grant, void, issue, listIssues}

type command struct {
	name           string
	trigger        string
	response       string
	secondResponse string
}

var verification = command{
	name:           "Verification",
	trigger:        "verificationmessage",
	response:       "In order to get access to %v, you must react to this message with the blue circle emoji :blue_circle:\n Enjoy the Server and remember, I'm watching you :eyes:",
	secondResponse: "",
}

var help = command{
	name:           "Help",
	trigger:        "help",
	response:       "To use me, start your command with `.` :wave:\nTo see my commands, try commanding with actions",
	secondResponse: "",
}

var ping = command{
	name:           "Ping",
	trigger:        "ping",
	response:       "Pong",
	secondResponse: "",
}

var invite = command{
	name:           "Invite",
	trigger:        "invite",
	response:       "Here is your invite to %v sent by %v\n%v",
	secondResponse: "%v has been DM'd and invited",
}

var serverAdTemplateCommand = command{
	name:    "Server Ad Template",
	trigger: "serveradtemplate",
	response: "**Instinct**\n" +
		"━━━━━━━━\n" +
		"**Are You Looking For A Good Gaming/Community Server On Discord? Well, Look No Further! Instinct Has A Variety Of Different Channels To Keep It Nice & Fresh! We Are Always Happy To Meet New People Since Our Goal Is To Make Our Server Amazing!**\n" +
		"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n" +
		"What We Offer:\n" +
		"🎮 - Gaming Channels!\n" +
		"👋 - Nice Community!\n" +
		"✅ - SFW Server!\n" +
		"🔎 - & So Much More!\n" +
		"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n" +
		"We Are Also Looking For:\n" +
		"🔧 - Server Staff!\n" +
		"💼 - Server Moderators!\n" +
		"⚡ - Server Boosters!\n" +
		"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n" +
		"👍If You're Interested In That, Join The Server For Application!\n" +
		"🔗Server Link: %v",
	secondResponse: "",
}

var hello = command{
	name:           "Hello",
	trigger:        "hello",
	response:       ":wave: <@%v>",
	secondResponse: "",
}

var hi = command{
	name:           "Hi",
	trigger:        "hi",
	response:       ":wave: <@%v>",
	secondResponse: "",
}

var hey = command{
	name:           "Hey",
	trigger:        "hey",
	response:       ":wave: <@%v>",
	secondResponse: "",
}

var laugh = command{
	name:           "Laugh",
	trigger:        "laugh",
	response:       ":joy:",
	secondResponse: "",
}

var bossCommand = command{
	name:           "Boss",
	trigger:        "boss",
	response:       "<@674370171904720897>",
	secondResponse: "",
}

var mrsBossCommand = command{
	name:           "Mrs Boss",
	trigger:        "mrsboss",
	response:       "<@696485894458310767>",
	secondResponse: "",
}

var minecraftHomeCoordinates = command{
	name:           "Minecraft home coordinates",
	trigger:        "mchome",
	response:       "-1017, 113, 638",
	secondResponse: "",
}

var minecraftBrewing = command{
	name:           "Minecraft brewing",
	trigger:        "mcbrew",
	response:       "https://www.reddit.com/r/gaming/comments/99yplr/minecraft_potion_brewing_guide/",
	secondResponse: "",
}

var mcPigRide = command{
	name:           "Minecraft Pig Ride",
	trigger:        "mcpigride",
	response:       "https://tenor.com/view/minecraft-gif-9643254",
	secondResponse: "",
}

var roleReactCustom = command{
	name:           "Role React Custom",
	trigger:        "customreact",
	response:       "Please ensure that the command is in the format of ```.customreact \"Title of the embed - Description of the embed\"```",
	secondResponse: "",
}

var actions = command{
	name:           "Actions",
	trigger:        "actions",
	response:       "",
	secondResponse: "",
}

var poll = command{
	name:           "Poll",
	trigger:        "poll",
	response:       "Please ensure that the command is in the format of ```.poll \"What the poll is\"```",
	secondResponse: "",
}

var gifCommand = command{
	name:           "Gif",
	trigger:        "gif",
	response:       "Please ensure that the command is in the format of ```.gif \"giphySearch\"```",
	secondResponse: "No gif found for the search `%v`",
}

var stickerCommand = command{
	name:           "Sticker",
	trigger:        "sticker",
	response:       "Please ensure that the command is in the format of ```.sticker \"giphySearch\"```",
	secondResponse: "No sticker found for the search `%v`",
}

var resetCustomRoles = command{
	name:           "Reset custom roles",
	trigger:        "rcr",
	response:       "Custom roles have been reset",
	secondResponse: "",
}

var grant = command{
	name:           "Grant Role",
	trigger:        "grant",
	response:       "Please ensure that the command is in the format of `.grant @role @user`",
	secondResponse: "",
}

var void = command{
	name:           "Void Role",
	trigger:        "void",
	response:       "Please ensure that the command is in the format of `.void @role @user`",
	secondResponse: "",
}

var issue = command{
	name:    "Create Github issue/request",
	trigger: "issue",
	response: "Issue created\n" +
		"Id: %v\n" +
		"Title: %v\n" +
		"Description: %v\n" +
		"Assigned to: %v\n" +
		"Url: %v",
	secondResponse: "Please ensure that the command is in the format of `.issue Title of the request - Description of the issue/request`",
}

var listIssues = command{
	name:    "List all current issues for botter",
	trigger: "listissues",
	response: "Issue %v\n" +
		"Id: %v\n" +
		"Title: %v\n" +
		"Assigned to: %v\n" +
		"Url: %v",
	secondResponse: "",
}

func (cmd command) equalsTriggerString(trigger string) bool {
	return cmd.trigger == helper.RemoveMultiSpaces(trigger)
}

func getCommand(trigger string) command {
	var blankCmd = command{}
	for _, cmd := range commands {
		if cmd.equalsTriggerString(trigger) {
			return cmd
		}
	}
	return blankCmd
}
