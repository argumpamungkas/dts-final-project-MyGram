package models

// RequestUserLogin represents the model for an User Login
type RequestUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RequestSocialMedia represents the model for an Social Media
type RequestSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

// RequestPhoto represents the model for an Photo
type RequestPhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

// RequestComment represents the model for an Comment
type RequestComment struct {
	Comment string `json:"comment_message"`
}
