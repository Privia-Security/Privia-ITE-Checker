# Finduk

### >_ Introduction
It is a tool prepared for IIS Tilde Enumeration vulnerability.
-recursiveRequest is a recursive function that sends OPTIONS requests to URLs formed by appending letters to a given prefix.
-It checks if the response status code is 404 and if the length of the prefix plus the current letter is 6.
-If so, it prints the success message.
-It then appends the current letter to the successLetters slice and checks if the length of the new prefix is divisible by 6.
-If so, it appends a newline character to the successLetters slice.
-Finally, it recursively calls itself with the new prefix.

### >_ Preview
![finduk_preview](https://github.com/Privia-Security/finduk/assets/81651239/bf2193c6-749d-4122-865c-42a66cb5d697)

### >_ Installation
```
go install github.com/Privia-Security/finduk@latest
```

### >_ How To Run
```
./finduk http://<IP-Address>/ txt,png
```

### >_ Help
```
Usage: <full URL> <extensions>
Example:
        finduk http://192.168.112.136/ txt,png
```

### >_ Features
- find file and directory names
- find file extension

### References
```
https://www.acunetix.com/vulnerabilities/web/microsoft-iis-tilde-directory-enumeration/
```
```
https://securitytrails.com/blog/microsoft-iis-shortname-tilde-vulnerability
```
