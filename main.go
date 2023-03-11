package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"projects/dapp/quiz"
	"projects/dapp/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var myenv map[string]string

const envLoc = ".env"

func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

func main() {
	loadEnv()

	ctx := context.Background()

	client, err := ethclient.Dial(myenv["GATEWAY"])
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	//==========================================

	// Check balance
	// accountAddress := common.HexToAddress(myenv["PUBLICADDRESS"])
	// balance, _ := client.BalanceAt(ctx, accountAddress, nil)
	// fmt.Printf("Balance: %d\n", balance)

	//==========================================
	// create a new session
	// We havn't spedified a value for the contract field yet, we will do that after obtaining a contract instance
	// when we deploy a new contract or whaen we load an existing
	session := NewSession(ctx, client)

	//load or deploy contract, and update session with contract instance
	if myenv["CONTRACTADDR"] == "" {
		session = NewContract(session, client, myenv["QUESTION"], myenv["ANSWER"])
	}

	// --- Deploy contract if doesn't exist
	// if we have an existing contract, load it; if we've deployed a new contract, attemp to load it.
	if myenv["CONTRACTADDR"] != "" {
		session = LoadContract(session, client)
	}
	//NOTE: To force the DApp to deploy a new contract, remove the CONTRACTADD entry in the .env file, or set it to an empty string

	//==========================================
	// Loop to implement simple CLI
	for {
		fmt.Printf(
			"\nPick an option:\n" + "" +
				"[1] Show question.\n" +
				"[2] Send answer.\n" +
				"[3] Check if you answered correctly.\n" +
				"[4] Exit.\n",
		)

		// Reads a single UTF-8 character (rune)
		// from STDIN and switches to case.
		switch readStringStdin() {
		case "1":
			readQuestion(session)
			break
		case "2":
			fmt.Println("Type in your answer")
			sendAnswer(session, readStringStdin())
			break
		case "3":
			checkCorrect(session)
			break
		case "4":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			break
		}
	}
}

// Sessions are wrappers that allow us to make contract calls without having to pass aorund auth, creds,
//
//	and call parameters constantly
//
// A session wraps:
//
//	 a contract instance
//		a bind.CallOpts struct that contains opts for making contract calls
//		a bind.TransactOpts struct that contains auth creds and params for created a valid Ethereum transaction
func NewSession(ctx context.Context, client *ethclient.Client) (session quiz.QuizSession) {
	loadEnv()

	// read data from keystore file
	keystore, err := os.Open(myenv["KEYSTORE"])
	if err != nil {
		log.Printf(
			"could not load keystore from location %s: %v\n",
			myenv["KEYSTORE"],
			err,
		)
	}
	defer keystore.Close()

	// get transact opts
	keystorepass := myenv["KEYSTOREPASS"]
	auth, err := bind.NewTransactor(keystore, keystorepass)
	if err != nil {
		log.Printf("Error getting auth: %v\n", err)
	}

	// set gass price after creating new session
	gasPrice, err := client.SuggestGasPrice(ctx)
	auth.GasPrice = gasPrice
	if err != nil {
		log.Printf("Error setting gass price: %v\n", err)
	}

	// Return session without contract instance
	return quiz.QuizSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

// ==========================================
// NewContract deploys a contract if no existing contract exists
func NewContract(session quiz.QuizSession, client *ethclient.Client, question string, answer string) quiz.QuizSession {
	log.Println("Creating new contract\n")
	loadEnv()

	// hash answer before sending it over network
	ansHash := utils.StringToKeccak256(answer)
	log.Printf("Default Ans: %s (%v)\n", answer, ansHash)

	// Hash answer before sending it over Ethereum network.
	contractAddress, tx, instance, err := quiz.DeployQuiz(&session.TransactOpts, client, question, ansHash)
	if err != nil {
		log.Fatalf("could not deploy contract: %v\n", err)
	}
	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance
	utils.UpdateEnvFile("CONTRACTADDR", contractAddress.Hex(), envLoc, myenv)
	return session
}

// LoadContract loads a contract if one exists
func LoadContract(session quiz.QuizSession, client *ethclient.Client) quiz.QuizSession {
	loadEnv()

	addr := common.HexToAddress(myenv["CONTRACTADDR"])
	instance, err := quiz.NewQuiz(addr, client)
	if err != nil {
		log.Fatalf("could not load contract: %v\n", err)
		log.Println(ErrTransactionWait)
	}
	session.Contract = instance
	return session
}

//// Contract interaction

// ErrTransactionWait should be returned/printed when we encounter an error that may be a result of the transaction not being confirmed yet.
const ErrTransactionWait = "if you've just started the application, wait a while for the network to confirm your transaction."

// readQuestion prints out question stored in contract.
func readQuestion(session quiz.QuizSession) {
	qn, err := session.Question()
	if err != nil {
		log.Printf("could not read question from contract: %v\n", err)
		log.Println(ErrTransactionWait)
		return
	}
	fmt.Printf("Question: %s\n", qn)
	return
}

// sendAnswer sends answer to contract as a keccak256 hash.
func sendAnswer(session quiz.QuizSession, ans string) {
	// Send answer
	txSendAnswer, err := session.SendAnswer(utils.StringToKeccak256(ans))
	if err != nil {
		log.Printf("could not send answer to contract: %v\n", err)
		return
	}
	fmt.Printf("Answer sent! Please wait for tx %s to be confirmed.\n", txSendAnswer.Hash().Hex())
	return
}

// checkCorrect makes a contract message call to check if
// the current account owner has answered the question correctly.
func checkCorrect(session quiz.QuizSession) {
	win, err := session.CheckBoard()
	if err != nil {
		log.Printf("could not check leaderboard: %v\n", err)
		log.Println(ErrTransactionWait)
		return
	}
	fmt.Printf("Were you correct?: %v\n", win)
	return
}

// readStringStdin reads a string from STDIN and strips any trailing \n characters from it.
func readStringStdin() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("invalid option: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n") // Important!
	return output
}
