# receipt_processor

docker build --platform=linux/amd64 -t receipt_challenge .

docker run -p 8080:8080 receipt_challenge

You can test the API via Swagger UI at http://localhost:8080/swagger/index.html