# E-Voting Application with Blockchain

This Go application provides a basic implementation of an e-voting system with blockchain features. It includes user registration, login, voting, and blockchain visualization.

## Overview

The application consists of the following components:

- `Candidate`: Represents a candidate in the election with a name and vote count.
- `Block`: Represents a block in the blockchain, containing details like index, timestamp, user ID, vote, hash, previous hash, and nonce.
- `Blockchain`: A slice of blocks forming the blockchain.
- `calculateHash`: Function to calculate the hash of a block based on its properties.
- `generateBlock`: Function to generate a new block with Proof of Work.
- `isBlockValid`: Function to check the validity of a new block.
- `register`: Function for user registration.
- `login`: Function for user login.
- `vote`: Function to allow a logged-in user to vote for a candidate.
- `main`: The main entry point for the application, containing the main menu and user interaction.

## Dependencies

- `github.com/fatih/color`: Used for adding color to console output.

## Usage

1. **Register**: Users can register with a unique user ID.
2. **Login**: Users can log in with their registered user ID.
3. **Vote**: Logged-in users can vote for one of the predefined candidates.
4. **Logout**: Users can log out to end their session.
5. **Show User's Voting History**: Displays the voting history of the currently logged-in user.
6. **Show Candidates' Votes**: Displays the current vote counts for each candidate.
7. **Show Blockchain**: Visualizes the blockchain, showing block details.
8. **Exit**: Exits the application.

## Proof of Work (PoW)

The application incorporates a basic Proof of Work mechanism to secure the blockchain and prevent spam.

## Colorful Console Output

The console output is enhanced with color using the `github.com/fatih/color` package. Different colors are used for various messages to improve readability.

## Getting Started

1. Install dependencies: `go get -u github.com/fatih/color`.
2. Run the application: `go run main.go`.

## Note

This is a simplified example for educational purposes. In a real-world scenario, additional security measures and a more sophisticated consensus algorithm would be necessary.

Feel free to modify and expand this code to meet specific requirements.

