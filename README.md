# Bank System Project in Golang

Welcome to the Bank System Project in Golang! This project aims to replicate the functionalities of a complete bank system, including managing bank accounts, balances, and transactions, all powered by the efficiency and concurrency features of Golang.

## Features

- **Complete Bank System**: This project offers a comprehensive set of APIs to manage various aspects of a bank system, including creating accounts, checking balances, transferring funds, and more.

- **Concurrency**: Leveraging the power of Go's concurrency features, this project ensures efficient handling of multiple requests simultaneously, providing a seamless banking experience.

- **Error Handling with Panic and Recover**: Robust error handling mechanisms using panic and recover ensure that the system remains stable even in unexpected scenarios, providing reliability to users.

- **PostgreSQL Database**: The project utilizes PostgreSQL to persistently store bank account details, balances, and transaction records, ensuring data integrity and durability.

- **Interfaces and Structs**: Extensive usage of interfaces and structs facilitates seamless data migration between the codebase and the database, enhancing maintainability and scalability.

- **API Testing with Postman**: The APIs designed in this project have been thoroughly tested using Postman, ensuring reliability and correctness in functionality.

## Routes

### Account Management

- `GET /account`: Get all accounts.
- `GET /account/:account_id`: Get details of a specific account.
- `POST /account`: Create a new account.
- `PATCH /account/:account_id`: Update details of a specific account.
- `DELETE /account/:account_id`: Delete a specific account.

## Getting Started

To get started with the Bank System Project:

1. **Clone the Repository**: Clone this repository to your local machine.

    ```bash
    git clone https://github.com/your-username/bank-system-golang.git
    ```

2. **Install Dependencies**: Make sure you have Go installed on your machine. Then, navigate to the project directory and install the required dependencies.

    ```bash
    cd bank-system-golang
    go mod tidy
    ```

3. **Set Up PostgreSQL Database**: Ensure you have PostgreSQL installed and running on your system. Create a new database for the project and configure the database connection details in the project's configuration file (`config.yaml`).

4. **Build and Run**: Build and run the project using the following command:

    ```bash
    go build
    ./bank-system-golang
    ```

5. **Testing with Postman**: Import the provided Postman collection (`bank-system-postman-collection.json`) into your Postman application. Use the imported collection to test the various APIs of the bank system project.

## Configuration

Modify the `config.yaml` file to configure the database connection details, server port, and other necessary parameters according to your environment.

```yaml
database:
  host: localhost
  port: 5432
  user: your_username
  password: your_password
  dbname: bank_system
  sslmode: disable

server:
  port: 8080
