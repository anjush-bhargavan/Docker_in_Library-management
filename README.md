# Library Management System with Dockerisation

This project is a Library Management System designed for local libraries, providing features for users, agents, and administrators. The backend is implemented in Go language using the Gin web framework, PostgreSQL for data storage, Redis for caching, Cron for scheduled tasks, Razorpay for payment processing, SMTP for email notifications, Golint for code linting, and GORM as the ORM.

## Docker

Created docker file for creating an image for the library management project.

## Technologies Used

- **Backend:**
  - Go language
  - Gin web framework
  - PostgreSQL for data storage
  - Redis for caching
  - Cron for scheduled tasks
  - Razorpay for payment processing
  - SMTP for email notifications
  - GORM as the ORM
  - Golint for code linting

## API Documentation

For detailed API documentation, refer to [API Documentation](https://documenter.getpostman.com/view/30219361/2s9YeEdCnd).


## Setup and Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/Docker_in_Library-management.git

2.Install Docker:
  
  Install Docker on your machine. Follow the official Docker installation guide for your operating system:

  Docker Installation Guide:
    https://docs.docker.com/engine/install/

3. Build Docker image
   
  Build the Docker image for the library management system:
    
    ```bash
    docker build -t library-management .

4. Running with Docker

   Run docker container only the library management project:

     ```bash
      docker run -p 8080:8080 -d library-management

5. Running the project with multiple conatiners including dependencies:

   Run docker compose command from the directory:

   ```bash
       docker-compose up
