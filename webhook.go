package webhookgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	url    string
	client *http.Client

	Username  string `json:"username,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Content   string `json:"content,omitempty"`

	Embeds []Embed `json:"embeds,omitempty"`
}

func NewWebhook(url string) *Webhook {
	return &Webhook{
		url:    url,
		Embeds: make([]Embed, 10),
		client: &http.Client{},
	}
}

func (w *Webhook) SetUsername(username string) *Webhook {
	w.Username = username
	return w
}

func (w *Webhook) SetAvatar(avatarUrl string) *Webhook {
	w.AvatarUrl = avatarUrl
	return w
}

func (w *Webhook) Send(message string) (WebhookResponse, error) {
	webhook := *w
	webhook.Content = message

	return w.send(webhook)
}

func (w *Webhook) SendEmbed(embed *Embed) (WebhookResponse, error) {
	webhook := *w
	webhook.Embeds = []Embed{*embed}

	return w.send(webhook)
}

func (w *Webhook) SendEmbeds(embeds []*Embed) (WebhookResponse, error) {
	webhook := *w

	var embedsList []Embed
	for _, e := range embeds {
		embedsList = append(embedsList, *e)
	}

	webhook.Embeds = embedsList

	return w.send(webhook)
}

func (w *Webhook) send(webhook Webhook) (WebhookResponse, error) {
	url := fmt.Sprintf("%s?wait=true", w.url)

	payload, err := json.Marshal(webhook)
	if err != nil {
		return WebhookResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return WebhookResponse{}, err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		return WebhookResponse{}, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		var responseBody WebhookResponse
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			return WebhookResponse{}, err
		}
		return responseBody, nil
	case 400:
		var responseBody BadRequestResponse
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			return WebhookResponse{}, err
		}
		return WebhookResponse{}, fmt.Errorf(responseBody.Message)
	case 401:
		return WebhookResponse{}, fmt.Errorf("invalid webhook token")
	default:
		return WebhookResponse{}, fmt.Errorf("received invalid status code [%d]", resp.StatusCode)
	}
}
