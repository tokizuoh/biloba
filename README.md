# biloba
Notify Slack of AWS charges using AWS Cost Explorer.
  
## Build

```bash
> docker-compose up --build -d
```

### .env
  
```
AWS_ACCESS_KEY_ID={your_access_key_id}
AWS_SECRET_ACCESS_KEY={your_secret_access_key}
```
  

## Run

```bash
> docker-compose exec app go run main.go
```
