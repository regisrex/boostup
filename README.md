## Email service using smtp.google.com and `net/smtp` package
___
⚠️ Works only when 2fa is enabled

####  To use it 
Set these two env vars
```
# filename: .env
ADMIN_EMAIL_ADDRESS=<sender-email-address>
ADMIN_EMAIL_PASSWORD=<application-password>
```

The password isn't your google account password but the password given after creating the app  
in your google account [https://myaccount.google.com/apppasswords](https://myaccount.google.com/apppasswords)
