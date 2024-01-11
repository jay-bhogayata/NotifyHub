[![ci workflow](https://github.com/jay-bhogayata/NotifyHub/actions/workflows/ci.yaml/badge.svg)](https://github.com/jay-bhogayata/NotifyHub/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jay-bhogayata/notifyHub)](https://goreportcard.com/report/github.com/jay-bhogayata/notifyHub)

# NotifyHub

NotifyHub is a robust notification service built with Go. It provides a simple way to send SMS and email notifications to users, leveraging the power of AWS services.

## Features

- Send SMS notifications.
- Send Email notifications.
- Easy to set up and use.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.x)
- AWS account with access to SES and SNS services

### Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/jay-bhogayat/notifyhub.git
```

2. Navigate to the project directory:

```bash
cd notifyhub
```

3. Build the application:

```bash
go build -o notifyhub
```

### Configuration

1. set environment variables

```bash
export PORT=8080
export SENDER_EMAIL=your_email_aws_ses_verified_sender_email
```

2. Configure your AWS credentials. Create a file named `credentials` at `~/.aws/` with the following content:

```bash
[default]
aws_access_key_id = YOUR_ACCESS_KEY
aws_secret_access_key = YOUR_SECRET_KEY
```

Replace `YOUR_ACCESS_KEY` and `YOUR_SECRET_KEY` with your actual AWS access key and secret key.

3. Configure your AWS region. Create a file named `config` at `~/.aws/` with the following content:

```bash 
[default]
region=us-east-1
```

### Running the Application

After completing the above steps, start the application by running:

```bash
./notifyhub
```

Now, the NotifyHub service is up and running!


### run with docker

```bash
docker build -t notifyhub .
docker run --rm -p 8080:8080 --env-file .env -v /home/your_user_name/.aws:/root/.aws/ notifyhub
```
