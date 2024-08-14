# Dream Mail Coding Challenge

## Prompt:
Create a generic email service service that is built on top of at least two of the below email providers. It should accepts
the basic information to send an email and attempt to send it using a primary service. If the primary service fails, it should 
failover to whichever service is the fallback without affecting your customers.


### Interpretation
This prompt is quite vague, as such I will attempt to define what my generic email service will do.

- Actions
    1. Happy Path: 
        - Have Valid Sender, Receiver, Body
        - Use, first Sender API

    2. Tests for Potential Failures
        - Valid Sender/Receiver Email
        - Check for API connection
        - Check for fallback failure

### Dependencies
To keep this as "Golang" as possible I will be using very minimal front end setup.
The following will be my tools outside of the standard Golang Library

    1. Echo      --> To create the webserver and be able to handle errors from Handlers
    2. Htmx      --> To minimize Front End and focus purely on the prompt
    3. Templ     --> A Go framework to create front end components

I have also added a couple of Go dependencies that help with security and formatting

    1. Gosec     --> A security dependency that checks for vulnerabillities 
    2. Gofmt     --> A Formatter for Go


### Data Flow
The user must be able to input text just like any email, meaning we will use the a form
to process user input from the front end.

```
  User Data -------------> Validation -----------> Request  
    ^                                                |
    |                                                |
    |-----------------Response-to-user---------------| 
```

### Testing
Currently, I am testing for failures in the sender email. the function `ValidateSend()` 
is responsible to making sure the email is correctly specified before sending it off to
MailGun or SparkPost.

### To Be Included



### Running the program

1. To initialize the Email Service run: `make startup`
2. After startup use: `make run` 

The Server will be hosted on localhost:9001
