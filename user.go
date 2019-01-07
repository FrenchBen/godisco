package godisco

import (
	"encoding/json"
	"fmt"
)

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
	User      UserInfo `json:"user"`
	Errors    []string `json:"errors,omitempty"`
	ErrorType string   `json:"error_type,omitempty"`
}

//UserInfo user information
type UserInfo struct {
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
	Featured []int64 `json:"featured_user_badge_ids"` // [],
	Card     string  `json:"card_badge"`              // null
}

type user struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Approved bool   `json:"approved"`
}

type invite struct {
	Email         string `json:"email"`
	GroupNames    string `json:"group_names"`
	CustomMessage string `json:"custom_message"`
}

// CreateResp typical response after user update
type CreateResp struct {
	Success   bool            `json:"success"`
	Message   string          `json:"message,omitempty"`
	Active    bool            `json:"active,omitempty"`
	UserID    int             `json:"user_id,omitempty"`
	Errors    json.RawMessage `json:"errors,omitempty"`
	Values    json.RawMessage `json:"values,omitempty"`
	Developer bool            `json:"is_developer"`
}

// InviteResp typical response after invite created
type InviteResp struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
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

//GetUsers send a request for information about all users
func GetUsers(req Requester, flag string, order string, ascending bool, page int, show_emails bool) (userInfo []*UserResponse, err error) {
	if flag == "" {
		flag = "active"
	}
	endpoint := fmt.Sprintf("/admin/users/list/%s.json", flag)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &userInfo)
	return userInfo, err
}

// CreateUser creates a new user based on details provided
func CreateUser(req Requester, name string, username string, email string, password string, active bool, approved bool) (response *CreateResp, err error) {
	update := &user{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
		Active:   active,
		Approved: approved,
	}
	data, err := json.Marshal(update)
	endpoint := "/users"
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

//DeactivateUser user
func DeactivateUser(req Requester, userID int) error {
	endpoint := fmt.Sprintf("/admin/users/%d/deactivate.json", userID)
	var data []byte
	_, _, err := req.Put(endpoint, data)
	return err
}

//ActivateUser user
func ActivateUser(req Requester, userID int) error {
	endpoint := fmt.Sprintf("/admin/users/%d/activate.json", userID)
	var data []byte
	_, _, err := req.Put(endpoint, data)
	return err
}

// InviteUser invite a new user
func InviteUser(req Requester, email string, message string) (*InviteResp, error) {
	i := &invite{
		Email:         email,
		CustomMessage: message,
	}
	data, err := json.Marshal(i)
	endpoint := "/invites"
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return nil, err
	}
	type response struct {
		Success string `json:"success"`
	}
	var r response
	err = json.Unmarshal(body, &r)
	invR := &InviteResp{}
	if r.Success == "OK" {
		invR.Success = true
	}
	invR.Message = r.Success
	return invR, err
}

// SendPasswordResetEmail Send password reset email
func SendPasswordResetEmail(req Requester, user string) error {
	type parameters struct {
		Login string `json:"login"`
	}
	type response struct {
		Result    string `json:"result"`
		UserFound bool   `json:"user_found"`
	}
	p := parameters{
		Login: user,
	}
	data, err := json.Marshal(p)
	endpoint := "/session/forgot_password"
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return err
	}
	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	if !r.UserFound {
		return fmt.Errorf("User %s not found", user)
	}
	return err
}
