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

Run my server software
```
docker run --name ESR ghcr.io/nottgy/easy-smtp-receiver -p 50:50
```

Check mail by running this command
```
docker logs ESR
```
