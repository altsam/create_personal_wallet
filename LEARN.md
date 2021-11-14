# Create your personal Solana wallet

Welcome to creating a Solana wallet quest wherein you’ll be creating your very own personal wallet application that can create a new wallet for you, import any old wallet, request airdrop in your wallet and finally, get the balance of your wallet! 

The best thing about Solana and other blockchains is that they allow us to manage money very seamlessly. To transact money on the blockchain, you need to use a software that facilitates these transactions called a wallet. Crypto wallets are physical devices or virtual programs that allow us to easily store and retrieve our crypto assets i.e. any cryptocurrencies that we might have on the blockchain.

Without further ado, let’s get started on creating our wallet.

# Setting up development environment

To create our CLI application, we’ll be using the Go language. Before we proceed, make sure you’ve a working Go runtime on your device. Let’s get started with setting up our project!


We’ll be using the Cobra go package to create our CLI scaffolding i.e. setting up our project skeleton for us that’ll be helpful to create our CLI application. To do so, run the following command
```
go install github.com/spf13/cobra/cobra@latest
```

Once this is done, we can create our code scaffolding using the following command
```
cobra init --pkg-name personal-wallet
```

On UNIX systems (Linux or MacOS), if you face an error stating “command not found”, append the following snippet in you “.bashrc” or “.zshrc”
```
export GOPATH=$HOME/go 
export GOBIN=$GOPATH/bin 
export PATH=${PATH}:$GOBIN
```
After running the `cobra init` command, you’ll see something like this

![1]()

Your project directory would look something like this

![2]()

# Installing the Solana SDK

In the `personal-wallet` directory, run the following command
```
go mod init personal-wallet
```
This creates a `go.mod` file, wherein we’ll be tracking all the dependencies of the different packages that will be used.

To interact with the Solana blockchain, we’ll be using the `solana-go-sdk`. Install the solana-go-sdk package in your project. In the terminal, type the following:
```
go get -u github.com/portto/solana-go-sdk@v1.8.1
```
Running this command must’ve modified your `go.mod` file. The updated `go.mod` file should have the following lines
```
module personal-wallet 
go 1.17 
require github.com/portto/solana-go-sdk v1.10.2 // indirect
```

Now, try running the following command
```
go run main.go
```
If you get an error message stating that some required packages are missing, then you can install them by running `go get <the-package-name>`

# Create your CLI commands

Let’s first give a description to our CLI. Inside the `cmd` directory, open the `root.go` file. Within it, find the variable `rootCmd`. Let’s change the description of our command.
```
Use:   "personal-wallet", 
Short: "Solana personal wallet", 
Long: `A CLI wallet application created in Go that interacts with the Solana blockchain.`
```
Now, if you run `main.go`, you’ll see the following output

![3]()

Let’s create commands for our CLI. In this quest, we’ll be implementing `Create New Wallet`, `Import Old Wallet`, `Request Airdrop` and `Transfer` functionality.
Let’s create commands for each.
```
cobra add create-wallet
cobra add import-wallet
cobra add request-airdrop
cobra add transfer
```
If you now check your `cmd` directory, you’ll see that for each command, a `.go` file has been created. Let’s update the description of each the same way we did for our main command.
Inside `createWallet.go`, update the `Use`, `Short` and `Long` parameters within the `createWalletCmd` variable with the following values.
```
Use:   "createWallet", 
Short: "Creates a new wallet", 
Long: “Creates a new wallet and provides wallet address and private key.”,
```

Similarly, for `importWallet.go`
```
Use:   "importWallet", 
Short: "Imports and existing wallet", 
Long: "Imports and existing wallet from a given private key.",
```

For `requestAirdrop.go`
```
Use:   "requestAirdrop", 
Short: "Request airdrop in Solana", 
Long: "Request airdrop to your public address.",
```

For `transfer.go`
```
Use:   "createWallet", 
Short: "Creates a new wallet", 
Long: "Transfer SOL from your wallet to other Solana wallets.",
```

Now, when you run `go run main.go`, it’ll list all the commands we just created

![4]()

# Create your wallet structure

Before we begin the implementation of the functionality, we need to first define the structure of our wallet. Within the `solana-sdk` exists a type for account that defines the object representing a Solana wallet using its private key in different forms(base58, byte slice and hex value). 

Create a new file inside the `cmd` directory and call it `utils.go`. Within this we’ll be defining our wallet structure and all the utility functions that we’ll be using for our CLI. Import the `solana-sdk` packages within this file.

```
package cmd

import ( 
"context"	
"io/ioutil"
"log"

"github.com/portto/solana-go-sdk/client"
"github.com/portto/solana-go-sdk/client/rpc"
"github.com/portto/solana-go-sdk/common"
"github.com/portto/solana-go-sdk/program/sysprog"
"github.com/portto/solana-go-sdk/types" 
)
```

Now, create a new type inside this file 
```
type Wallet struct { 
}
```
Let’s add the following lines to it
```
account   types.Account 
c 	    *client.Client 
```
`account` is a parameter that’ll be holding the Solana wallet object and the `c` parameter holds the RPC client object that’ll be used to connect to the Solana network.

Putting everything together, the final code would look like this
```
package cmd

import (
   "github.com/portto/solana-go-sdk/client"
   "github.com/portto/solana-go-sdk/types"
)

type Wallet struct {
   account types.Account
   c       *client.Client
}
```
If you see an error from VS Code mentioning that “the required packages are missing”, then simply run the following command in the `personal-wallet` root directory.
```
go mod tidy
``` 

# Implement create new wallet functionality

The `solana-go-sdk` package provides a `NewAccount()` function that returns a newly generated Solana wallet. Let’s create a function that creates a new account for us.

We’ll be adding out functionalities in the `utils.go`. Create an empty function called `CreateNewWallet` and add it to the end of `utils.go`. This function takes in the RPCEndpoint as a string and returns a Wallet object.

```
func CreateNewWallet(RPCEndpoint string) Wallet {
}
```

`solana-go-sdk` provides a `NewAccount()` method in `types` for creating a new account. Also, `NewClient` method within `client` provides the `Client` object that holds the RPCEndpoint object within the wallet struct. We can use these methods in this function. 
We can create the account by running the following command
```
newAccount := types.NewAccount()
```

Let’s store the private key of this newly created account in a new file called data
```
data := []byte(newAccount.PrivateKey) // convert the private key to byte array for storage
err := ioutil.WriteFile("data", data, 0644) 
if err != nil { 
log.Fatal(err) 
}
```

Final function looks like
```
func CreateNewWallet(RPCEndpoint string) Wallet {
// create a new wallet using types.NewAccount()
   	newAccount := types.NewAccount()
   	data := []byte(newAccount.PrivateKey)

  	err := ioutil.WriteFile("data", data, 0644)
   	if err != nil {
       		log.Fatal(err)
   	}

  	return Wallet{
       		newAccount,
       	client.NewClient(RPCEndpoint),
   	}
}
```
Now we just need to call this function in `createWallet.go`. We’ll be updating `run` parameter within `createWalletCmd`. 
```
var createWalletCmd = &cobra.Command{
   Use:   "createWallet",
   Short: "Creates a new wallet",
   Long:  "Creates a new wallet and provides wallet address and private key.",
   Run: func(cmd *cobra.Command, args []string) {
       fmt.Println("Creating new wallet.")
       wallet := CreateNewWallet(rpc.DevnetRPCEndpoint)
       fmt.Println("Public Key: " + wallet.account.PublicKey.ToBase58())
       fmt.Println("Private Key Saved in 'data' file")
   },
}
```
The `run` parameter contains the function that’ll be executed when the command is executed. We’re simply calling over here the function we created in `utils.go`. 

# Implementing import old wallet functionality

What if we wanted to import one of our existing wallets? `solana-go-sdk` provides us with a handy function to do so, called `AccountFromBytes`.
Within the `utils.go` file add a new empty function called `ImportOldWallet`. This function will take the private key in byte array

What if we wanted to import one of our existing wallets? `solana-go-sdk` provides us with a handy function to do so, called `AccountFromBytes`.
Within the `utils.go` file add a new empty function called `ImportOldWallet`. This function will take the private key in a byte array within the `key_data` file and then import the existing wallet to our `key_data` file which is holding the key.
Again we’ll be using this function to create our `wallet` object which will be used in future functions.
We’ll be adding out functionalities in the `utils.go`. Create an empty function called `CreateNewWallet` and add it to the end of `utils.go`. This function takes in the privateKey as a byte array, RPCEndpoint as a string and returns a Wallet object.

```
func ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {
}
```

Let’s fill in this function.
```
wallet, err := types.AccountFromBytes(privateKey)
if err != nil {
return Wallet{}, err
}

return Wallet {
wallet,
client.NewClient(RPCEndpoint),
}, nil
```
The `solana-go-sdk` provides a handy function to import an account using the `AccountFromBytes` function. Finally, the `ImportOldWallet` function would look like the following.
```
func ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) { 
// import a wallet with bytes slice private key 
wallet, err := types.AccountFromBytes(privateKey) 
if err != nil { 
return Wallet{}, err 
} 
return Wallet{ 
wallet, 
client.NewClient(RPCEndpoint), 
}, nil 
}
```
Let’s call this function in our `importWallet` command. We’ll be updating the `run` parameter within `importWalletCmd`. 
```
var importWalletCmd = &cobra.Command{
Use:   "importWallet",
Short: "Imports and existing wallet",
Long:  "Imports and existing wallet from a given private key in the 'data' file and returns a wallet object.",
Run: func(cmd *cobra.Command, args []string) {
fmt.Println("Importing wallet from the 'key_data' file.") 
wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
      		fmt.Println("Public Key: " + wallet.account.PublicKey.ToBase58())
      		balance, _ := GetBalance()
       		fmt.Println("Wallet balance: " + strconv.Itoa(int(balance/1e9)) + “SOL”)
},
}
```
The `run` parameter contains the function that’ll be executed when the command is executed. We’re simply calling over here the function we created in `utils.go`. We’ll soon see how to implement the `GetBalance` function that we’re using over here to get the balance of our wallet.

# Implementing get balance functionality
Every wallet must be able to tell its balance. Let’s implement the balance functionality for our wallet. We’ll be adding a `GetBalance` function in `utils.go`, this function can then be used to fetch the balance of the wallet existing in the `key_data` file.
Create an empty function called `GetBalance` which returns amount as an integer(uint64).
```
func GetBalance() (uint64, error) {
}
```
First we need to import our wallet, since we’re developing the application, we’ll be passing the endpoint as devnet.
```
wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
balance, err := wallet.c.GetBalance( 
context.TODO(),                      // request context
wallet.account.PublicKey.ToBase58(), // wallet to fetch balance for 
) 
if err != nil { 
return 0, nil 
}
```
The `GetBalance` function returns us the balance in lamports. We can then convert this to SOL by dividing with 1e9.
Finally, our function looks like
```
func GetBalance() (uint64, error) {
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    balance, err := wallet.c.GetBalance( 
        context.TODO(),                      // request context
        wallet.account.PublicKey.ToBase58(), // wallet to fetch balance for 
    ) 
    if err != nil { 
        return 0, nil 
    }
    return balance, nil
}
```

# Implementing the airdrop functionality
For every transaction on the blockchain, we pay fees in SOL. To test out transacting on the blockchain, Solana allows us to ‘airdrop’ ourselves some play SOL to our wallet. Let’s create a function that allows us to airdrop SOL into our wallets.
Create an empty function called `RequestAirdrop` and add it to the end of `utils.go`. This function takes in the amount of SOL to be airdropped and returns the transaction confirmation hash as string. 
```
func ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {
}
```
Let’s first import our wallet from the `key_data` file and also convert our SOL amount to lamports as the `RequestAirdrop` function takes the amount in lamports (1SOL = 1e9 lamports). We can do so using the following command
```
wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
amount = amount * 1e9 // turning SOL into lamports 
```
Now, we can use the `RequestAirdrop` function to airdrop SOL in our wallets, in Devnet.
```
txhash, err := wallet.c.RequestAirdrop( 
    context.TODO(),                      // request context 
    wallet.account.PublicKey.ToBase58(), // wallet address requesting airdrop 
    amount,                              // amount of SOL in lamport 
) 
```

Finally, our function will look like
```
func RequestAirdrop(amount uint64) (string, error) { 
    // request for SOL using RequestAirdrop() 
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    amount = amount * 1e9 // turning SOL into lamports 
    txhash, err := wallet.c.RequestAirdrop( 
        context.TODO(),                      // request context wallet.account.PublicKey.ToBase58(), // wallet address requesting airdrop 
        amount,                              // amount of SOL in lamport 
    ) 
    if err != nil { 
        return "", err 
    } 
    return txhash, nil 
}
```


Let’s call this function in our `requestAirdrop` command. We’ll be updating the `run` parameter within `requestAirdropCmd`. 
```
var requestAirdropCmd = &cobra.Command{
    Use:   "requestAirdrop",
    Short: "Request airdrop in Solana",
    Long:  "Request airdrop to your public address passed to the command.",
    Run: func(cmd *cobra.Command, args []string) {
            wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
            fmt.Println("Requesting airdrop to: " + wallet.account.PublicKey.ToBase58()) 
            amount, _ := strconv.ParseUint(args[0], 10, 64) 
            txhash, _ := RequestAirdrop(amount) 
            fmt.Println("Airdropped " + strconv.Itoa(int(amount)) + " SOL.\nTransaction hash: " + txhash)
        },
}
```
The `run` parameter contains the function that’ll be executed when the command is executed. Over here, we’re calling the `RequestAirdrop` function that’ll airdrop SOL into the wallet existing in our `key_data` file. You might notice that we’re using `args[0]` over here. The airdrop amount will be passed as a CLI option and that’s why we’re using command line arguments.

# Implementing the transfer function
What’s the point of a wallet if you cannot transfer funds? Let’s create a transfer function that takes in the public key of the receiver and transfers SOL to them.
Create an empty function called `Transfer`. This function takes in the receiver address as string and the amount in uint64. It returns the transactions hash as a string.
```
func Transfer(receiver string, amount uint64) (string, error) { 
}
```
To transfer funds, we first need to create a transaction message and the transaction signer details alongwith the public key of the recipient. We then send this message to the blockchain, and receive a transaction hash for our transaction confirmation.

First, we fetch the most recent blockhash. Again, we’ll be transacting over the devnet. We can use the `GetRecentBlockhash` to do so.
```
wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
response, err := wallet.c.GetRecentBlockhash(context.TODO()) 
if err != nil { 
    return "", err 
} 
```
Now, let’s create a new transfer message with the latest block hash. We can do so by using the `NewMessage` function from `types` which we had imported earlier.
```
message := types.NewMessage( 
    wallet.account.PublicKey, // public key of the transaction signer
    []types.Instruction{
        sysprog.Transfer( 
            wallet.account.PublicKey,  //public key of the transaction sender 
            common.PublicKeyFromString(receiver), // wallet address of receiver 
            amount,                               // transaction amount in lamport 
        ), 
    }, 
    response.Blockhash, // recent block hash 
) 
```

Now that we have create the transaction message, we’ll now be creating a transaction object using the `NewTransaction` function alongwith the transaction signer details.
```
tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account}) 
if err != nil { 
    return "", err 
} 
```
Now we can send this transaction to the blockchain using the `SendTransaction2` function.
```
txhash, err := wallet.c.SendTransaction2(context.TODO(), tx) 
if err != nil { 
    return "", err 
} 
```

Finally, your function would look like
```
func Transfer(receiver string, amount uint64) (string, error) { 
    // fetch the most recent blockhash 
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    response, err := wallet.c.GetRecentBlockhash(context.TODO()) 
    if err != nil { 
        return "", err 
    } 

    // make a transfer message with the latest block hash 
    message := types.NewMessage( 
        wallet.account.PublicKey, // public key of the transaction signer
        []types.Instruction{
            sysprog.Transfer( 
                wallet.account.PublicKey, // public key of the transaction sender 
                common.PublicKeyFromString(receiver), // wallet address of the transaction receiver 
                amount,                               // transaction amount in lamport 
            ), 
        }, 
        response.Blockhash, // recent block hash 
    ) 

    // create a transaction with the message and TX signer 
    tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account}) 
    if err != nil { 
        return "", err 
    } 

    // send the transaction to the blockchain 
    txhash, err := wallet.c.SendTransaction2(context.TODO(), tx) 
    if err != nil { 
        return "", err 
    } 
    return txhash, nil 
}
```

Let’s call this function in our `transfer` command. We’ll be updating the `run` parameter within `transferCmd`. 
```
var transferCmd = &cobra.Command{
    Use:   "transfer",
    Short: "Transfer SOL",
    Long:  "Transfer SOL from your wallet to other Solana wallets.",
    Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Recepient address: " + args[0]) 
            fmt.Println("Amount to be sent: " + args[1]) 
            amount, _ := strconv.ParseUint(args[1], 10, 64) 
            txhash, _ := Transfer(args[0], amount) 
            fmt.Println("Transaction complete.\nTransaction hash: " + txhash)
        },
}
```
The `run` parameter contains the function that’ll be executed when the command is executed. The first parameter to our CLI would be the recipient’s transaction ID and the second parameter would be the amount to be airdropped in SOL. In the above function we’re simply calling Transfer function that’ll provide us with the transaction hash.

# Putting it all together

Now that we’ve implemented our four commands, let’s test it out. You can test out your commands with the following snippet
```
go run main.go <command name>
```
Let’s test our `createWallet` command. Run the following
```
go run main.go createWallet
```
You should see something like this

![5]()

Let’s test the `importWallet` command. You can use this command to get your existing wallet details such as your public key. To import a new wallet, simply replace the `key_data` file with your new wallet’s private key stored as a byte array. Run the following command.
```
go run main.go importWallet
```
You should see the following

![6]()

Since we’ve a newly created wallet, our balance is 0. Let’s airdrop ourselves some SOL to play with. Keep in mind, we can airdrop a maximum 5 SOL in one transaction. If we try to airdrop more, our balance will remain unaffected.
```
go run main.go requestAirdrop 3
```

![7]()

We can check this transaction hash on the Solana explorer to confirm the transaction.
![8]()

Run the `importWallet` command again to confirm the change in balance.
![9]()

Now, let’s try to transfer some of this SOL to another address. We’ll be transferring 1 SOL to `7tWk3ZKZ6ohSkb9Yxrj87uvSYCLwH3QhYjGiG9yiUEKF`.
To do so, you run the following command
```
go run main.go transfer <recepient public address> <amount in SOL>
```

![10]()

You can confirm the transaction by viewing it on the explorer.

![11]()

We can see that 2 SOL has been received by `7tWk3ZKZ6ohSkb9Yxrj87uvSYCLwH3QhYjGiG9yiUEKF` sent from `BE7b78GDLRGVmdorGa89SaoEFDpdaJ39qGMVHGfh6LFt` which is our wallet.

# Conclusion
Congratulations to all the quest masters on completing this quest. You now have a functioning personal Solana blockchain CLI wallet. You can make your CLI wallet more robust by using CLI flags provided by the `cobra` go package and make the code more robust, the possibilities are endless! Cheers to all learners on their Solana learning journey!

