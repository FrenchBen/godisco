package godisco

import (
	"encoding/json"
	"fmt"
)

//Category expected struct for a single category
type Category struct {
	Id                           int    `json:"id"`
	Name                         string `json:"name"`
	Color                        string `json:"color"`
	TextColor                    string `json:"text_color"`
	Slug                         string `json:"slug"`
	TopicCount                   int    `json:"topic_count"`
	PostCount                    int    `json:"post_count"`
	Position                     int    `json:"position"`
	Description                  string `json:"description"`
	DescriptionText              string `json:"description_text"`
	TopicUrl                     string `json:"topic_url"`
	ReadRestricted               bool   `json:"read_restricted"`
	Permission                   int    `json:"permission"`
	NotificationLevel            string `json:"notification_level"`
	TopicTemplate                string `json:"topic_template"`
	HasChildren                  bool   `json:"has_children"`
	SortOrder                    string `json:"sort_order"`
	SortAscending                string `json:"sort_ascending"`
	ShowSubcategoryList          bool   `json:"show_subcategory_list"`
	NumFeaturedTopics            int    `json:"num_featured_topics"`
	DefaultView                  string `json:"default_view"`
	SubcategoryListStyle         string `json:"subcategory_list_style"`
	DefaultTopPeriod             string `json:"default_top_period"`
	MinimumRequiredTags          int    `json:"minimum_required_tags"`
	NavigateToFirstPostAfterRead bool   `json:"navigate_to_first_post_after_read"`
	TopicsDay                    int    `json:"topics_day"`
	TopicsWeek                   int    `json:"topics_week"`
	TopicsMonth                  int    `json:"topics_month"`
	TopicsYear                   int    `json:"topics_year"`
	TopicsAllTime                int    `json:"topics_all_time"`
	DescriptionExcerpt           string `json:"description_excerpt"`
	SubcategoryIds               []int  `json:"subcategory_ids"`
	UploadedLogo                 string `json:"uploaded_logo"`
	UploadedBackground           string `json:"uploaded_background"`
}

//CategoriesResponse expected struct for GetCategoryList response
type CategoriesResponse struct {
	CategoryList struct {
		CanCreateCategory bool        `json:"can_create_category"`
		CanCreateTopic    bool        `json:"can_create_topic"`
		Draft             bool        `json:"draft"`
		DraftKey          string      `json:"draft_key"`
		DraftSequence     int         `json:"draft_sequence"`
		Categories        []*Category `json:"categories"`
	} `json:"category_list"`
}

//CategoryResponse expected struct for GetCategory and PostCategory response
type CategoryResponse struct {
	Category struct {
		Id                           int      `json:"id"`
		Name                         string   `json:"name"`
		Color                        string   `json:"color"`
		TextColor                    string   `json:"text_color"`
		Slug                         string   `json:"slug"`
		TopicCount                   int      `json:"topic_count"`
		PostCount                    int      `json:"post_count"`
		Position                     int      `json:"position"`
		Description                  string   `json:"description"`
		DescriptionText              string   `json:"description_text"`
		TopicUrl                     string   `json:"topic_url"`
		ReadRestricted               bool     `json:"read_restricted"`
		Permission                   int      `json:"permission"`
		NotificationLevel            string   `json:"notification_level"`
		CanEdit                      bool     `json:"can_edit"`
		TopicTemplate                string   `json:"topic_template"`
		HasChildren                  string   `json:"has_children"`
		SortOrder                    string   `json:"sort_order"`
		SortAscending                string   `json:"sort_ascending"`
		ShowSubcategoryList          bool     `json:"show_subcategory_list"`
		NumFeaturedTopics            int      `json:"num_featured_topics"`
		DefaultView                  string   `json:"default_view"`
		SubcategoryListStyle         string   `json:"subcategory_list_style"`
		DefaultTopPeriod             string   `json:"default_top_period"`
		MinimumRequiredTags          int      `json:"minimum_required_tags"`
		NavigateToFirstPostAfterRead bool     `json:"navigate_to_first_post_after_read"`
		AvailableGroups              []string `json:"available_groups"`
		AutoCloseHours               string   `json:"auto_close_hours"`
		AutoCloseBasedOnLastPost     bool     `json:"auto_close_based_on_last_post"`
		EmailIn                      string   `json:"email_in"`
		EmailInAllowStrangers        bool     `json:"email_in_allow_strangers"`
		MailinglistMirror            bool     `json:"mailinglist_mirror"`
		SuppressFromLatest           bool     `json:"suppress_from_latest"`
		AllTopicsWiki                bool     `json:"all_topics_wiki"`
		CannotDeleteReason           string   `json:"cannot_delete_reason"`
		AllowBadges                  bool     `json:"allow_badges"`
		TopicFeaturedLinkAllowed     bool     `json:"topic_featured_link_allowed"`
		UploadedLogo                 string   `json:"uploaded_logo"`
		UploadedBackground           string   `json:"uploaded_background"`
		// "group_permissions": []*GroupPermission
		//   type GroupPermission struct {
		//     "permission_type": 1,
		//     "group_name": "everyone"
		//   }
		// ],
		// "custom_fields": {},
	} `json:"category"`
}

//CategoryRequest expected struct for PostCategory request
type CategoryRequest struct {
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Color            string `json:"color"`
	TextColor        string `json:"text_color"`
	Description      string `json:"description"`
	ParentCategoryId string `json:"parent_category_id"`
	// permissions
}

// GetCategoryList show list of categories
func GetCategoryList(req Requester) (categories *CategoriesResponse, err error) {
	endpoint := "/categories.json"
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &categories)
	return categories, err
}

func GetCategory(req Requester, id string) (category *CategoryResponse, err error) {
	endpoint := fmt.Sprintf("/c/%s/show.json", id)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &category)
	return category, err
}

func PostCategory(req Requester, category *CategoryRequest) (response *CategoryResponse, err error) {
	data, err := json.Marshal(category)
	endpoint := "/categories.json"
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}
