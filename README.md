# Dream Mail Coding Challenge

## Prompt:
Create a generic email service service that is built on top of at least two of the below email providers. It should accepts
the basic information to send an email and attempt to send it using a primary service. If the primary service fails, it should 
failover to whichever service is the fallback without affecting your customers.


### Interpretation
This prompt is quite open, given the timeframe I will attempt to create a sender service
with two email services, (`MailGun, SparkPost`).

- Actions
    1. Happy Path: 
        - Have Valid Sender, Receiver, Subject, Body
        - Use, first Sender

    2. Not So Happy Path
        - Have Valid Sender, Receiver, Subject, Body
        - First Sender Fails
        - Second Sender Works 

    3. The Bad Place:
        - Bad Sender Email
        - Bad Reciever Email (Needs to be implemented)

    2. Tests for Potential Failures
        - Valid Sender/Receiver Email
        - Check for API connection
        - Check for fallback failure

### Dependencies
To keep this as "Golang" as possible I will be using very minimal front end setup.
The following will be my tools outside of the standard Golang Library

    1. Echo      --> To create the webserver and be able to handle errors from Handlers
    2. Htmx      --> To minimize Front End and focus purely on the prompt
    3. Templ     --> To create the loyout form in html that go can then run 

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
Three areas of testing occur, first, testing for any race conditions that could 
cause incorrect data to be sent to a email. Second, making sure that the emails 
are correctly specified. Third, making sure that the reciever email follows a 
proper tld protocol based on a self generated list.

### Desired Updates
    1. UI
        - General Layout and Design (CSS...)
        - Modifying Layout to have potential for more options
    2. Testing
        - Test for Endpoints
    3. Logic
        - Check for valid Recipient Email
        - Correctly specified Body/Subject
    4. Middleware
        - Logging (amount of emails sent)
    5. Security
        - limit amout of emails sent out at one time
        - Authentication
    

### Running the program

The two email services involved are [MailGun](https://www.mailgun.com/) and [SparkPost](https://developers.sparkpost.com/). 
To run the code on your own computer make sure to provide the follwing in and .env file.

```
export MAIL_GUN_API_KEY=<INSERT APIKEY>
export SPARK_POST_API=<INSERT APIKEY>
```


1. To initialize the Email Service run: `make startup`
2. After startup use: `make run` 

The Server will be hosted on localhost:9001
