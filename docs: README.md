# image-management-service

Image Management Service is a Go-based RESTful API for managing images. With this service, you can upload, download, and manage images in a simple and efficient way.

## Features

- Reads a list of image URLs from a file named links.txt and downloads them to a folder named images.
- Upload images to the server
- Store metadata about each image, including its original URL, local name, file extension, file size, and download date
- Get a list of stored images
- Download stored images
- Concurrently upload new images

## Requirements

- Go 1.16+
- Docker
- Docker Compose

## Installation

1. Clone the repository:

```
git clone https://github.com/yeganeh666/image-management-service.git
```

2. Build the Docker containers:

```
make docker
```

3. Access the web service at `http://localhost:8080`.

## API Documentation

API documentation is available in Swagger format. You can view the Swagger docs by visiting http://localhost:8080/swagger/index.html in your web browser.

## Contributing

Contributions are welcome! Please create a new branch and submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
