package bot

import (
	"github.com/Jacobbrewer1/botter/config"
	"github.com/bwmarrin/discordgo"
	"log"
)

func verifier(s *discordgo.Session, m *discordgo.MessageCreate, roleRequired role, singleRole ...bool) bool {
	// Level 0 is everyone so they auto have permissions
	if roleRequired.level == 0 || config.IgnoreVerification {
		return true
	}

	processFunc := func(m *discordgo.MessageCreate, message string) {
		log.Printf("%v does not have permissions for %v", m.Author.Username, m.Content)
		if _, err := sendReplyMessage(s, m, message); err != nil {
			log.Println(err)
			return
		}
		return
	}

	if singleRole != nil {
		if singleRole[0] {
			if !verifySingleRole(m.Member.Roles, roleRequired.id) {
				go processFunc(m, failedVerificationResponse)
				return false
			}
			log.Printf("%v has permissions for %v", m.Author.Username, m.Content)
			return true
		}
	}

	if !verifyMultiRole(m.Member.Roles, roleRequired.level) {
		go processFunc(m, failedVerificationResponse)
		return false
	}
	log.Printf("%v has permissions for %v", m.Author.Username, m.Content)
	return true
}

func verifySingleRole(memberRoles []string, roleRequired string) bool {
	for _, memberRole := range memberRoles {
		if memberRole == roleRequired {
			return true
		}
	}
	return false
}

func verifyMultiRole(memberRoles []string, level int) bool {
	for _, role := range memberRoles {
		for _, serverRole := range roles {
			if serverRole.level < level {
				continue
			}
			if role == serverRole.id && serverRole.level <= level {
				return true
			}
		}
	}
	return false
}
