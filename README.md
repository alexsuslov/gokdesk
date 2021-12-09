# Gokdesk

gokdesk is a Golang wrapper for accessing OKDESK using the REST API.

## Use
```
gokdesk
  -comment string
        add issue comment by id
  -config string
        path to  env file (default ".env")
  -get string
        get issue by id
  -status string
        set issue status by id
      
```

## Get issue by id
```
gokdesk -get 142753
```
![screen](https://raw.githubusercontent.com/alexsuslov/gokdesk/main/doc/screen.png)

## Set status

```
cat status_update_req.json | gokdesk -status 142753
```
## Add comment

```
cat comment_add_req.json | gokdesk -comment 142753
```
