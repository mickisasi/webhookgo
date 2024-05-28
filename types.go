package webhookgo

import "time"

type BadRequestResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type WebhookResponse struct {
	Type      int    `json:"type"`
	ChannelID string `json:"channel_id"`
	Content   string `json:"content"`
	Embeds    []struct {
		Type   string `json:"type"`
		Title  string `json:"title"`
		Color  int    `json:"color"`
		Fields []struct {
			Name   string `json:"name"`
			Value  string `json:"value"`
			Inline bool   `json:"inline"`
		} `json:"fields"`
		Thumbnail struct {
			URL      string `json:"url"`
			ProxyURL string `json:"proxy_url"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
		} `json:"thumbnail"`
	} `json:"embeds"`
	Timestamp time.Time `json:"timestamp"`
	Flags     int       `json:"flags"`
	ID        string    `json:"id"`
	Author    struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Discriminator string `json:"discriminator"`
		PublicFlags   int    `json:"public_flags"`
		Flags         int    `json:"flags"`
		Bot           bool   `json:"bot"`
	} `json:"author"`
	Pinned          bool   `json:"pinned"`
	MentionEveryone bool   `json:"mention_everyone"`
	Tts             bool   `json:"tts"`
	WebhookID       string `json:"webhook_id"`
}
