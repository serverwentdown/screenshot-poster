package main

import (
	"os"
	"path/filepath"
	"io"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Source string
	Target string
	Delay time.Duration
	Webhook ConfigWebhook
	S3 ConfigS3
}

type ConfigWebhook struct {
	URL string
	Username string
	AvatarURL string `yaml:"avatar_url"`
}

type ConfigS3 struct {
	Secure bool
	Endpoint string
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket string
	Prefix string
}
	

func LoadConfig() (Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return Config{}, err
	}
	buf, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return Config{}, err
	}

	if config.Target == "" {
		config.Target = filepath.Join(config.Source, "resized")
	}

	return config, nil
}