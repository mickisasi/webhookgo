package webhookgo

import "time"

type Embed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Color       int    `json:"color,omitempty"`

	Footer    footer    `json:"footer,omitempty"`
	Image     image     `json:"image,omitempty"`
	Thumbnail thumbnail `json:"thumbnail,omitempty"`
	Author    author    `json:"author,omitempty"`
	Fields    []field   `json:"fields,omitempty"`
}

type footer struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type image struct {
	Url string `json:"url,omitempty"`
}

type thumbnail struct {
	Url string `json:"url,omitempty"`
}

type author struct {
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

func NewEmbed() *Embed {
	return &Embed{}
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetAuthor(name, iconUrl, url string) *Embed {
	e.Author.Name = name
	e.Author.Url = url
	e.Author.IconUrl = iconUrl
	return e
}

func (e *Embed) SetURL(url string) *Embed {
	e.Url = url
	return e
}

func (e *Embed) AddField(name, value string, inline bool) *Embed {
	e.Fields = append(e.Fields, field{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetThumbnail(url string) *Embed {
	e.Thumbnail.Url = url
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) SetImage(url string) *Embed {
	e.Image.Url = url
	return e
}

func (e *Embed) SetFooter(text, iconUrl string) *Embed {
	e.Footer.Text = text
	e.Footer.IconUrl = iconUrl
	return e
}

func (e *Embed) SetTimestamp(t time.Time) *Embed {
	currentTime := t.UTC().Format(time.RFC3339)
	e.Timestamp = currentTime
	return e
}
