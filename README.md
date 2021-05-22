# biloba
Notify Slack of AWS charges using AWS Cost Explorer.
  
## Docker 
  
```bash
> docker --version
Docker version 19.03.12, build 48a66213fe

> docker-compose --version
docker-compose version 1.27.2, build 18f557f9
```
  
## Build

```bash
> docker-compose up --build -d
```

### .env
  
```
AWS_ACCESS_KEY_ID={your_access_key_id}
AWS_SECRET_ACCESS_KEY={your_secret_access_key}
AWS_COST_EXPLORER_IMG_PATH={image_path}
SLACK_CHANNEL_ID={slack_channel_id}
SLACK_BOT_TOKEN={slack_bot_token}
```
  

## Run

```bash
> docker-compose exec app go run main.go
```
