# biloba
  
![biloba](https://socialify.git.ci/tokizuoh/biloba/image?description=1&font=Source%20Code%20Pro&language=1&logo=https%3A%2F%2Fuser-images.githubusercontent.com%2F37968814%2F119221357-ebdb4080-bb29-11eb-9474-ed254d7627b8.png&owner=1&theme=Light)
  
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
