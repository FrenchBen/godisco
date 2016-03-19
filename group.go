package godisco

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GroupResponse expected struct for groups
type GroupResponse struct{}

// GetGroup show details of a given group
func GetGroup(req Requester, group string) (groupInfo *GroupResponse, err error) {
	endpoint := fmt.Sprintf("/groups/%s.json", group)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupInfo)
	return groupInfo, err
}

// GetGroupMembers list members of a given group
func GetGroupMembers(req Requester, groupName string) (groupInfo *GroupResponse, err error) {
	endpoint := fmt.Sprintf("/groups/%s/members.json", groupName)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupInfo)
	return groupInfo, err
}

//@TODO
// Implement add to group - Name: "DockerForMacWinBeta" - id: 45
// APi Call: POST - '%s/admin/groups/%s?api_key=%s&api_username=%s'
// alias_level = 0

type groupUpdate struct {
	Name       string `json:"group[name]"`
	AliasLevel int    `json:"group[alias_level]"`
	Visible    bool   `json:"group[visible]"`
	Users      string `json:"group[usernames]"`
}

// GroupInfo describes the group update received
type GroupInfo struct {
	Basic struct {
		ID           int    `json:"id"`
		Automatic    bool   `json:"automatic"`
		Name         string `json:"name"`
		UserCount    int    `json:"user_count"`
		AliasLevel   int    `json:"alias_level"`
		Visible      bool   `json:"visible"`
		Domains      string `json:"automatic_membership_email_domains"`
		Retroactive  bool   `json:"automatic_membership_retroactive"`
		Primary      bool   `json:"primary_group"`
		Title        string `json:"title"`
		Trust        string `json:"grant_trust_level"`
		Incoming     string `json:"incoming_email"`
		Notification int    `json:"notification_level"`
		Messages     bool   `json:"has_messages"`
		Mentionable  bool   `json:"mentionable"`
	} `json:"basic_group"`
}

func updateGroupMembers(req Requester, groupName string, groupID string, members []string) (groupInfo *GroupInfo, err error) {
	update := &groupUpdate{
		Name:       groupName,
		AliasLevel: 0,
		Visible:    false,
		Users:      strings.Join(members, ","),
	}
	data, err := json.Marshal(update)
	endpoint := fmt.Sprintf("/admin/groups/%s", groupID)
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupInfo)
	return groupInfo, err
}
