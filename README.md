# easy-smtp-receiver

The simplest smtp server ever.
I just want to quickly receive small amount of mail.
You know, sometimes you want to receive mail with
custom address. Only for the first time and just to
feel more professional, you may use nice email address
like hr@mycompany.com. But it takes ages and billions
of dollars to setup. Instead, you just pull docker container
and start it. Let it run and over time, emails will just be
all listed in docker logs.

## Usage

1. Set up MX record for mycompany.com to point to mail.mycompany.com.
1. Set up A record for mail.mycompany.com to point to your server ip address.
1. On your server install docker with command:
```
curl -fsSL https://get.docker.com | sh
```

4. Run my server software
```
docker pull ghcr.io/nottgy/easy-smtp-receiver:latest
docker run --name ESR -p 25:25 -d ghcr.io/nottgy/easy-smtp-receiver
```

Check mail by running this command
```
docker logs ESR
```

## Updating/Restarting

First stop container
```
docker stop ESR
```

Then remove container
```
docker rm ESR
```

then go back to the installation process from the step 4
