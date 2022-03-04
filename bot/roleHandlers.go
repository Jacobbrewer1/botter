package bot

import (
	"errors"
	"fmt"
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Example : .grant @Squad @doozy
// The grant bool will br true for a grant and false for a void
func roleHandler(s *discordgo.Session, m *discordgo.MessageCreate, vfy role, grantBool bool, vfySingleRole ...bool) (bool, error) {
	if !verifier(s, m, vfy, vfySingleRole...) {
		return true, nil
	}

	confirmRole := func(s *discordgo.Session, guildID, testRoleId string, guildRoles []*discordgo.Role) error {
		for _, r := range guildRoles {
			if r.ID == testRoleId {
				return nil
			}
		}
		return errors.New("id of role does not exist in Guild")
	}

	if strings.Contains(message.query, "<@") {
		elms := strings.Split(message.query, " ")
		if message.isEmpty() || (len(elms) != 2) {
			return false, errors.New("incorrect format of query")
		}
		roleId, err := helper.FormatIdFromMention(elms[0])
		if err != nil {
			return false, err
		}
		userId, err := helper.FormatIdFromMention(elms[1])
		if err != nil {
			return false, err
		}
		user, err := s.GuildMember(m.GuildID, userId)
		// If the user ID is not correct do not continue
		if err != nil {
			return false, err
		}
		guildRoles, err := s.GuildRoles(m.GuildID)
		if err != nil {
			return false, err
		}
		// Testing to make sure that the role exists
		if err := confirmRole(s, m.GuildID, roleId, guildRoles); err != nil {
			return false, err
		}
		var rolePassId string
		var roleName string
		for _, r := range guildRoles {
			if r.ID == roleId {
				rolePassId = r.ID
				roleName = r.Name
				break
			}
		}
		if grantBool {
			if err := addRole(s, m.GuildID, rolePassId, user.User.ID); err != nil {
				return false, err
			}
			if _, err := sendReplyMessage(s, m, fmt.Sprintf(roleSuccessResponse, user.User.ID, roleName)); err != nil {
				return true, err
			}
		} else {
			if err := voidRole(s, m.GuildID, rolePassId, user.User.ID); err != nil {
				return false, err
			}
			if _, err := sendReplyMessage(s, m, fmt.Sprintf(roleVoidSuccessResponse, user.User.ID, roleName)); err != nil {
				return true, err
			}
		}
	} else {
		return false, errors.New("query is in the wrong format")
	}
	return true, nil
}

func addRole(s *discordgo.Session, guildId, roleID, userId string) error {
	return s.GuildMemberRoleAdd(guildId, userId, roleID)
}

func voidRole(s *discordgo.Session, guildId, roleId, userId string) error {
	return s.GuildMemberRoleRemove(guildId, userId, roleId)
}
