package bot

import (
	"errors"
	"fmt"
	"github.com/Jacobbrewer1/botter/api"
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

func createGithubIssue(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if message.isEmpty() {
		return errors.New("no query parameters provided")
	}
	elms := strings.Split(helper.RemoveMultiSpaces(message.query), "-")
	if len(elms) > 2 {
		return errors.New("length of m.content split is greater than 2")
	}
	var ni = api.NewIssue{
		Title: elms[0],
	}
	if len(elms) == 2 {
		ni.Body = elms[1] + fmt.Sprintf("\n\n**Issue created by %v**", m.Author.Username)
	} else {
		ni.Body = fmt.Sprintf("Issue created by %v", m.Author.Username)
	}
	i, u, err := api.CreateIssue(ni.Title, ni.Body)
	if err != nil {
		return err
	}

	if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(issue.response, *i.Number, *i.Title,
		helper.RemoveBoldness(strings.Replace(*i.Body, "\n", "", 1)),
		*u.Name, *i.HTMLURL)); err != nil {
		return err
	}
	return nil
}

func getGithubIssues(s *discordgo.Session, m *discordgo.MessageCreate) error {
	issues, err := api.GetBotterIssues()
	if err != nil {
		return err
	}

	var issueString string
	for issueNumber, i := range issues {
		temp := ""
		if i.IsAssigned() {
			var users string
			if len(i.Assignees) > 1 {
				for _, u := range i.Assignees {
					x, err := api.GetUser(*u.Login)
					if err != nil {
						return err
					}
					if users == "" {
						users = *x.Name
					} else {
						users = users + ", " + *x.Name
					}
				}
			} else {
				u, err := api.GetUser(*i.Assignee.Login)
				if err != nil {
					return err
				}
				users = *u.Name
			}
			temp = fmt.Sprintf(listIssues.response, issueNumber+1, *i.Number, *i.Title, users, *i.HTMLURL)
		} else {
			temp = fmt.Sprintf(listIssues.response, issueNumber+1, *i.Number, *i.Title, issueNotAssignedText, *i.HTMLURL)
		}
		if (len(temp)+len(issueString)) > 1999 && issueString != "" {
			if _, err := sendMessage(s, m.ChannelID, issueString); err != nil {
				return err
			}
			issueString = ""
		}
		issueString = issueString + temp + "\n\n"
	}
	if _, err := sendMessage(s, m.ChannelID, issueString); err != nil {
		return err
	}
	return nil
}

func autoIssue(s *discordgo.Session, channelId string, e error) {
	log.Println("creating auto issue for botter error")
	i, u, err := api.CreateIssue(autoIssueTitle,
		fmt.Sprintf(autoIssueBody, time.Now().Format(time.RFC1123), e),
		autoCreatedLabel)
	if err != nil {
		log.Println(err)
		return
	}

	if _, err := sendMessage(s, channelId, fmt.Sprintf(issue.response, *i.Number, *i.Title,
		helper.RemoveBoldness(strings.Replace(*i.Body, "\n", "", 1)),
		*u.Name, *i.HTMLURL)); err != nil {
		log.Println(err)
		return
	}
}
