## Simple SMTP client Go
in order to solve Computer assignment of Computer Networks in SBU (Spring 1400) We implemented a email sender (SMTP client) in Go

more info about project could be found [here](./project-description.pdf)

there are few ways to send an email, most people use libraries to do this for example go's built-in SMTP package or external packages like [go-mail](https://github.com/go-gomail/gomail) but we just used socket and read [RFC](https://tools.ietf.org/html/rfc5321) and our network book to found out how SMTP work in action. so this is it. 

reading this small project source code might help you to understand how an email sent to SMTP server.


## how to use 

this program is written in pure go, so for compilation you need to install `go>=1.16`

after that that would be enough to run:

```sh
# compile and run 
go build 
./smtp-client -u username@gmail.com -p mypassword -dest reciever-email@gmail.com -subj "your subject" -body "and finally email body" 

# just run
go run main.go --help
```

## 

## limitations

+ not able to send attachment or multimedia message or even non-ASCII text.
+ only Gmail sender mail is supported by now (receiver is not limited)
+ multiple receivers or CC/BCC isn't supported yet. 
+ you should enable "not secure apps" in your Google account security setting.
+ your password isn't sent over network in plain-text but use with your own risk.
+ 2-Step verification is not implemented.




## Implementation details

you can read about details of implementation in `report.md`



## useful links 

+ about SMTP authentication [+](https://www.ndchost.com/wiki/mail/test-smtp-auth-telnet)
+ SMTP send mail explained [+](https://www.ndchost.com/wiki/mail/test-smtp-auth-telnet)
+ gmail less secure apps [+](https://support.google.com/accounts/answer/6010255?p=less-secure-apps&hl=en&visit_id=637528055941711149-3770501630&rd=1)







