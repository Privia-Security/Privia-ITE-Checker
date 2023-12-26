# Finduk

### >_ Introduction
It is a tool prepared for IIS Tilde Enumeration vulnerability.

### >_ Preview
![finduk_preview](https://github.com/Privia-Security/finduk/assets/81651239/bf2193c6-749d-4122-865c-42a66cb5d697)

### >_ Installation
```
go install github.com/Privia-Security/finduk@latest
```

### >_ How To Run
```
./finduk http://192.168.112.136/ txt,png
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
