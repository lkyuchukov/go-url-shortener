# go-url-shortener
A URL shortener written in Go

## Usage
```shell
# start the app

# create a new short url 
curl -X POST 'localhost:9000/shorten' -d '{"longUrl": "https://github.com/trending"}' 

# returns {"shortUrl":"http://localhost:9000/QFWFS4nF"}

# navigate to http://localhost:9000/QFWFS4nF and get redirected
```
