
# screenshot-poster

Watches for screenshots in a folder and posts them to Discord

## Features

- [x] Watch folder for files
- [x] Compress images
  - [ ] Configurable compression options
- [x] Upload images to S3
- [x] Send image URL to Discord

## Usage

You'll need a S3-compatible bucket (Backblaze and MinIO work as well) with a write capable access key and secret key. Additionally, the bucket should allow the public to read files.

Download a copy from GitHub Releases, and write the following configuration file as `config.yaml` in the same directory:

```yaml
# Folder to watch for screenshots
source: C:\path\to\folder
# Time to wait after file creation until the file is fully written
delay: 1s

# Upload configuration
s3:
  secure: yes
  endpoint: s3.us-east-1.amazonaws.com
  access_key: S3 Access Key ID
  secret_key: S3 Secret Key
  bucket: bucket-name
  # Optional prefix for bucket objects
  prefix: minecraft-rocks

# Webhook configuration
webhook:
  # Create a Discord webhook and change this url
  url: "https://discord.com/webhookURL"
  # Optional avatar URL
  avatar_url: "https://static.wikia.nocookie.net/minecraft_gamepedia/images/0/0b/Wooden_Pickaxe_JE2_BE2.png"
  # Optional bot username
  username: "Minecraft Screenshots"

# Resize configuration
resize:
  enabled: no
  quality: 95
```

<!-- vim: set conceallevel=2 et ts=2 sw=2: -->
