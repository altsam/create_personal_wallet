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
"github.com/portto/solana-go-sdk/client" 
"github.com/portto/solana-go-sdk/common" "github.com/portto/solana-go-sdk/program/sysprog" "github.com/portto/solana-go-sdk/types" 
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

# Conclusion

Congratulations on completing the Solana wallet quest! You now have your very own Solana blockchain wallet to play with and share with your friends. Cheers on your Solana development journey!
