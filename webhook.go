package main

import (
	"log"
	"io"
	"errors"
	"bytes"
	"encoding/json"
	"net/http"
)

type DiscordPayload struct {
	Username *string `json:"username"`
	AvatarURL *string `json:"avatar_url"`
	Content *string `json:"content"`
	Embeds []DiscordPayloadEmbed `json:"embeds"`
}

type DiscordPayloadEmbed struct {
	Image *DiscordPayloadEmbedImage `json:"image"`
}

type DiscordPayloadEmbedImage struct {
	Url *string `json:"url"`
}

var ErrBadResponseStatus = errors.New("bad response status code")

func PostEmbed(config ConfigWebhook, url string) error {
	blob, err := json.Marshal(DiscordPayload{
		Username: &config.Username,
		AvatarURL: &config.AvatarURL,
		Embeds: []DiscordPayloadEmbed{
			DiscordPayloadEmbed{
				Image: &DiscordPayloadEmbedImage {
					Url: &url,
				},
			},
		},
	})
	if err != nil {
		return err
	}

	err = WebhookPost(blob, config.URL)
	if err != nil {
		return err
	}
	return nil
}

func PostContent(config ConfigWebhook, url string) error {
	blob, err := json.Marshal(DiscordPayload{
		Username: &config.Username,
		AvatarURL: &config.AvatarURL,
		Content: &url,
	})
	if err != nil {
		return err
	}

	err = WebhookPost(blob, config.URL)
	if err != nil {
		return err
	}
	return nil
}

func WebhookPost(blob []byte, url string) error {
	buf := bytes.NewBuffer(blob)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := io.ReadAll(resp.Body)
		log.Printf("Post: status %s: %s", resp.Status, respBody)
		return ErrBadResponseStatus
	}

	return nil
}