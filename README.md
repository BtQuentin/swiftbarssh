# Swiftbarssh


This is a small tool that read the SSH config file and display each host as a list.

You can click on every line, and it will open a SSH connection in the terminal.

The SSH config file should respect the following format:

```editorconfig
....
 
Host "Displayed Name for host"
   HostName <hostname> 
   User <username> 
   IdentityFile <identity_file>
   
Host "Displayed Name for host"
   HostName <hostname> 
   User <username> 
   IdentityFile <identity_file>

....
```

## Installation

- Clone the repo
- Change the `displayedText` and `configFile` in the `main.go` file
- Run `$ cd <path_to_repo> && go build .`
- Edit `metadata.txt` if you want to change the details
- Run `$ xattr -w "com.ameba.SwiftBar" "$(cat metadata.txt | base64)" swiftbarssh`
- Move the output file to the SwiftBar folder