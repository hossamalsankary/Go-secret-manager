
# Go-Secret-Manager

**Go-Secret-Manager** is a lightweight backend application built with Go (Golang) that securely manages secrets via a simple REST API. It enforces restricted access using an IP-based security mechanism, ensuring only trusted IP addresses can access stored secrets.

## Features

- **Secure Storage**: Stores secrets in a dedicated JSON module.  
- **IP-Based Access Control**: Only allows requests from IPs specified in `VAULT_IPS`.  
- **REST API**: Provides a simple REST interface for retrieving secrets.  
- **Easy to Deploy**: Containerized with Docker for easy setup and portability.

## Prerequisites

- **Golang** (if running outside of Docker)
- **Docker** (optional, for container-based deployment)
- **Environment Variable**: `VAULT_IPS` must be set with a comma-separated list of IPs allowed to access the service.

## Getting Started

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/your-username/go-secret-manager.git
   cd go-secret-manager
   ```
2. **Configure Allowed IPs**  
   Ensure the environment variable `VAULT_IPS` includes the IP addresses permitted to connect. For example:  
   ```bash
   export VAULT_IPS="41.235.69.116,127.0.0.1"
   ```

## Running with Docker

1. **Build the Docker Image**  
   ```bash
   docker build -t secret-manager .
   ```

2. **Run the Docker Container**  
   ```bash
   docker run -d -p 8080:8080 -e VAULT_IPS="41.235.69.116,127.0.0.1" secret-manager
   ```
   - The container will start the application on port **8080**.
   - Only the IPs listed in `VAULT_IPS` will be able to access it.

3. **Test the Application**  
   ```bash
   curl localhost:8080
   ```
   - If your IP is in `VAULT_IPS`, you’ll receive a valid response.
   - If not, you’ll get a **403 Forbidden** error.

## Running Locally (Without Docker)

1. **Install Golang**  
   Make sure you have Go installed and set up correctly on your machine.

2. **Export the Environment Variable**  
   ```bash
   export VAULT_IPS="41.235.69.116,127.0.0.1"
   ```

3. **Build and Run**  
   ```bash
   go build -o secret-manager .
   ./secret-manager
   ```
   - The application will start listening on port **8080** by default.

4. **Test the Application**  
   ```bash
   curl localhost:8080
   ```
   - If your IP is in `VAULT_IPS`, you’ll receive a valid response.
   - If not, you’ll get a **403 Forbidden** error.

## Storing Secrets

- All sensitive data is stored in a JSON module within the code.  
- You can modify this JSON module to include your own secrets.  
- Ensure proper security measures are in place when committing or distributing your code or images.

## Contributing

Contributions are welcome! To contribute:

1. Fork this repository  
2. Create a feature branch  
3. Make your changes and add tests if needed  
4. Submit a pull request for review

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use and modify this code for your own projects, subject to the terms of this license.


