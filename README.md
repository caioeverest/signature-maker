### signature-maker
This project creates a simple email signature

### How do I use this?
- First of all you need to clone the repository with the "go get"
`go get github.com/caioever/signature-maker`

- After cloning this repository you go at the folder and execute the following commands
`export HTTP_PORT=80` First you need to set some port
`go run main.go` Then just run the main.go and Voilá!

- Now It's simple, you just need to send a POST request with a json body that contains all parameters in the handler.
```json
{
	"imgLink": "http://",
	"orgName": "xpto",
	"name": "Jonh Smith",
	"role": "Software Engineer",
	"area": "Development",
	"address": "Some place",
	"phone1": "111 111 1111",
	"phone2": null
}
```

#Curl Exemple
```curl
curl -X POST \
  http://{your host}:{some port}/signature \
  -H 'Content-Type: application/json' \
  -d '{
	"imgLink": "http://",
	"orgName": "xpt",
	"name": "Jonh Smith",
	"role": "Software Engineer",
	"area": "Development",
	"address": "Some place",
	"phone1": "111 111 1111",
	"phone2": null
}'
```
