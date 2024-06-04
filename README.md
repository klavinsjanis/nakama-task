# Nakama Task

## Project Setup

1. **Clone Repository:** Begin by cloning the project repository.
2. **Download Dependencies:**
   - Obtain Go 1.21.6.
   - Install Docker.
3. **Start the Project:**
   - Execute `docker-compose up --build` in your shell to initiate the project.
4. **Stop the Project:**
   - Use `docker-compose down` to halt the project.
   - Optionally, include `-v` to remove associated Docker volumes.

## Description

The project utilizes Docker to create Nakama and PostgreSQL containers. Golang plugins enhance Nakama server with custom logic.

A custom RPC function is implemented alongside a new table in the Nakama database. The MD5 hashing algorithm secures file content hashes.

Once the containers are operational and stable, integration tests can be manually conducted to validate custom functionality.

## Potential Improvements

1. **Database Migrations:** Implement proper database migration procedures.
2. **Error Handling Enhancement:** Enhance error handling mechanisms for smoother operation.
3. **Request Validation Refinement:** Improve request validation processes to fortify security.
4. **Dependency Inversion:** Implement dependency inversion for better modularity and testability.