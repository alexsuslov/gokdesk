# Gokdesk

gokdesk is a Golang wrapper for accessing OKDESK using the REST API.

## Use
```
gokdesk 
  -config string
        path to  env file (default ".env")
  -get string
        get issue by id
```

## Get issue by id
```
gokdesk -get 142753
```
![screen](https://raw.githubusercontent.com/alexsuslov/gokdesk/main/doc/screen.png)
