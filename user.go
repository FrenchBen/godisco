package godisco

import (
	"encoding/json"
	"fmt"
)

// UserReq interface to facilitate testing
type UserReq interface {
	GetUser(user string) (userInfo *UserResponse, err error)
}

// UserResponse define json response
type UserResponse struct {
	UserBadges []struct {
		ID        int    `json:"id"`
		Granted   string `json:"granted_at"`
		BadgeID   int    `json:"badge_id"`
		UserID    int    `json:"user_id"`
		GrantedBy int    `json:"granted_by_id"`
	} `json:"user_badges,omitempty"`
	Badges []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		GrantCount   int    `json:"grant_count"`
		AllowTitle   bool   `json:"allow_title"`
		MultiGrant   bool   `json:"multiple_grant"`
		Icon         string `json:"icon"`
		Image        string `json:"image"`
		Listable     bool   `json:"listable"`
		Enabled      bool   `json:"enabled"`
		BadgeGroupID int    `json:"badge_grouping_id"`
		System       bool   `json:"system"`
		Slug         string `json:"slug"`
		BadgeTypeID  int    `json:"badge_type_id"`
	} `json:"badges,omitempty"`
	BadgeTypes []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Order int    `json:"sort_order"`
	} `json:"badge_types,omitempty"`
	Users []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Avatar    string `json:"avatar_template"`
		Name      string `json:"name"`
		Moderator bool   `json:"moderator"`
		Admin     bool   `json:"admin"`
	} `json:"users,omitempty"`
	User      userInfo `json:"user"`
	Errors    []string `json:"errors,omitempty"`
	ErrorType string   `json:"error_type,omitempty"`
}

type userInfo struct {
	ID           int      `json:"id"`
	Username     string   `json:"username"`
	Avatar       string   `json:"avatar_template"`
	Name         string   `json:"name"`
	Posted       string   `json:"last_posted_at"`
	Seen         string   `json:"last_seen_at"`
	Created      string   `json:"created_at"`
	Website      string   `json:"website_name"`
	Edit         bool     `json:"can_edit"`
	EditUser     bool     `json:"can_edit_username"`
	EditEmail    bool     `json:"can_edit_email"`
	EditName     bool     `json:"can_edit_name"`
	PM           bool     `json:"can_send_private_messages"`
	PMUser       bool     `json:"can_send_private_message_to_user"`
	TrustLevel   int      `json:"trust_level"`
	Moderator    bool     `json:"moderator"`          // true,
	Admin        bool     `json:"admin"`              // true,
	Title        string   `json:"title"`              // null,
	AvatarID     int      `json:"uploaded_avatar_id"` // 15,
	BadgeCount   int      `json:"badge_count"`        // 0,
	CustomFields struct{} `json:"custom_fields"`      // {},
	Pending      int      `json:"pending_count"`      // 0,
	View         int      `json:"profile_view_count"` // 2,
	InvitedBy    string   `json:"invited_by"`         // null,
	Groups       []struct {
		ID          int    `json:"id"`                                 // 10,
		Auto        bool   `json:"automatic"`                          // true,
		Name        string `json:"name"`                               // "trust_level_0",
		Count       int    `json:"user_count"`                         // 2,
		Alias       int    `json:"alias_level"`                        // 0,
		Visible     bool   `json:"visible"`                            // true,
		Domains     string `json:"automatic_membership_email_domains"` // null,
		Retro       bool   `json:"automatic_membership_retroactive"`   // false,
		Primary     bool   `json:"primary_group"`                      // false,
		Title       string `json:"title"`                              // null,
		Trust       string `json:"grant_trust_level"`                  // null,
		Messages    bool   `json:"has_messages"`                       // false,
		Mentionable bool   `json:"mentionable"`                        // false
	} `json:"groups"` //
	Featured []string `json:"featured_user_badge_ids"` // [],
	Card     string   `json:"card_badge"`              // null

}

// GetUser sends a request for information about the user
func GetUser(req Requester, user string) (userInfo *UserResponse, err error) {
	endpoint := fmt.Sprintf("/users/%s.json", user)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &userInfo)
	return userInfo, err
}
