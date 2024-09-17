# Test Client to server proxy config


## Description

Uses nginx reverse proxy configuration to forward tcp messages from a 
nodejs client sender to a go server. 

POC for ProductWebsite.

## Dependencies
- [go](https://go.dev/doc/install)
- [node](https://nodejs.org/en/download/package-manager)
- [nginx](https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/)

## Instructions to run

1) Start nginx reverse proxy servver

    ``` sudo nginx -c <workingdirectory>/ngnix_config.conf```

    Note: if you dont put the full working directory, it errors out 

2) Start go reciver server
    
    ``` go run go_reciever.go```

3) Start nodejs sender

    ```node client_sender.js ```