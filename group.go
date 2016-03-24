package godisco

import (
	"encoding/json"
	"fmt"
)

// GroupResponse expected struct for groups
type GroupResponse struct {
	Basic struct {
		ID          int    `json:"id"`
		Automatic   bool   `json:"automatic"`
		Name        string `json:"name"`
		UserCount   int    `json:"user_count"`
		AliasLevel  int    `json:"alias_level"`
		Visible     bool   `json:"visible"`
		Domain      string `json:"automatic_membership_email_domains"`
		Retroactive bool   `json:"automatic_membership_retroactive"`
		Primary     bool   `json:"primary_group"`
		Title       string `json:"title"`
		TrustLevel  int    `json:"grant_trust_level"`
		Messages    bool   `json:"has_messages"`
		Mentionable bool   `json:"mentionable"`
	} `json:"basic_group"`
}

// GroupMembersResponse defines list of members in a group
type GroupMembersResponse struct {
	Members []Member `json:"members"`
	Owners  []Member `json:"owners"`
	Meta    struct {
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
}

// Member information about a member
type Member struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar_template"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Posted   string `json:"last_posted_at"`
	Seen     string `json:"last_seen_at"`
}

// GetGroup show details of a given group
func GetGroup(req Requester, groupName string) (groupInfo *GroupResponse, err error) {
	endpoint := fmt.Sprintf("/groups/%s.json", groupName)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupInfo)
	return groupInfo, err
}

// GetGroupMembers list members of a given group
func GetGroupMembers(req Requester, groupName string) (groupMemberInfo *GroupMembersResponse, err error) {
	endpoint := fmt.Sprintf("/groups/%s/members.json?limit=10000", groupName)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupMemberInfo)
	return groupMemberInfo, err
}

//@TODO
// Implement add to group - Name: "DockerForMacWinBeta" - id: 45
// APi Call: POST - '%s/admin/groups/%s?api_key=%s&api_username=%s'
// alias_level = 0

type groupUpdate struct {
	GroupID string   `json:"group_id"`
	Users   []string `json:"users"`
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
		GroupID: groupID,
		Users:   members,
	}
	data, err := json.Marshal(update)
	endpoint := "/admin/groups/bulk"
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupInfo)
	return groupInfo, err
}
