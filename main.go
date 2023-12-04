package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/fatih/color"

)

const difficulty = 2

type Candidate struct {
	Name  string
	Votes int
}

type Block struct {
	Index     int
	Timestamp string
	UserID    string
	Vote      string
	Hash      string
	PrevHash  string
	Nonce     int
}

type Blockchain []Block

var blockchain Blockchain
var currentUserID string
var candidates = []Candidate{
	{Name: "Candidate A", Votes: 0},
	{Name: "Candidate B", Votes: 0},
	{Name: "Candidate C", Votes: 0},
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.UserID + block.Vote + block.PrevHash + string(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, vote string) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.UserID = currentUserID
	newBlock.Vote = vote
	newBlock.PrevHash = prevBlock.Hash

	for nonce := 0; ; nonce++ {
		newBlock.Nonce = nonce
		newBlock.Hash = calculateHash(newBlock)

		// Check if the hash meets the difficulty criteria
		if strings.HasPrefix(newBlock.Hash, strings.Repeat("0", difficulty)) {
			break
		}
	}

	return newBlock
}

func isBlockValid(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}

	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	// Check if the hash meets the difficulty criteria
	if !strings.HasPrefix(newBlock.Hash, strings.Repeat("0", difficulty)) {
		return false
	}

	return true
}

func register() {
	fmt.Println("Enter your user ID:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userID := scanner.Text()
	currentUserID = userID
	color.Green("Registration successful.")
}

func login() {
	fmt.Println("Enter your user ID:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userID := scanner.Text()
	currentUserID = userID
	color.Green("Login successful.")
}

func vote() {
	if currentUserID == "" {
		color.Yellow("Please login first.")
		return
	}

	fmt.Println("Candidates:")
	for i, candidate := range candidates {
		fmt.Printf("%d. %s\n", i+1, candidate.Name)
	}

	var candidateIndex int
	fmt.Println("Enter the number of the candidate you want to vote for:")
	fmt.Scanln(&candidateIndex)

	if candidateIndex < 1 || candidateIndex > len(candidates) {
		color.Red("Invalid candidate selection.")
		return
	}

	vote := candidates[candidateIndex-1].Name
	candidates[candidateIndex-1].Votes++

	newBlock := generateBlock(blockchain[len(blockchain)-1], vote)

	if isBlockValid(newBlock, blockchain[len(blockchain)-1]) {
		blockchain = append(blockchain, newBlock)
		fmt.Println("Vote recorded successfully.")
	} else {
		color.Red("Error: Invalid block. Vote not recorded.")
	}
}


func main() {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		UserID:    "",
		Vote:      "",
		PrevHash:  "",
		Hash:      "",
		Nonce:     0,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	blockchain = append(blockchain, genesisBlock)

	var choice int

	for {
		fmt.Println()
		color.Cyan("1. Register")
		color.Cyan("2. Login")
		color.Cyan("3. Vote")
		color.Cyan("4. Logout")
		color.Cyan("5. Show User's Voting History")
		color.Cyan("6. Show Candidates' Votes")
		color.Cyan("7. Show Blockchain")
		color.Cyan("8. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			register()
		case 2:
			login()
		case 3:
			vote()
		case 4:
			currentUserID = ""
			color.Green("Logout successful.")
		case 5:
			if currentUserID == "" {
				color.Yellow("Please login first.")
				continue
			}
			color.Cyan("User's voting history:")
			for i, block := range blockchain {
				if block.UserID == currentUserID && block.Vote != "" {
					fmt.Printf("Block %d: You voted for %s\n", i, block.Vote)
				}
			}
		case 6:
			color.Cyan("Candidates' Votes:")
			for _, candidate := range candidates {
				fmt.Printf("%s: %d votes\n", candidate.Name, candidate.Votes)
			}
		case 7:
			fmt.Println("Blockchain:")
			for _, block := range blockchain {
				fmt.Printf("Block %d - Hash: %s, PrevHash: %s, Nonce: %d\n", block.Index, block.Hash, block.PrevHash, block.Nonce)
			}
		case 8:
			color.Yellow("Exiting...")
			return
		default:
			color.Red("Invalid choice.")
		}
	}
}
