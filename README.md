adamwestum
==========

**adamwestum** is a web app that generates dummy text content that's all Adam West quotes.

# Decisions

The frontend is built in Polymer.

The backend is written in Go. It provides an API to generate a list of quotes of arbitrary length.

# Developing

In order to run locally, you will need to do these things.

## /etc/adamwestum/adamwestum.gcfg

There is a config template file that you can copy. This is for running in development, a deployment version is further down.

```
[Host]
port = :2222
path = /path/to/adamwestum
```

## go build

Any time you change Go code, you will need to `go build` to generate a new executable.

## go test

Run the unit tests with `go test ./...`

## ./adamwestum

Run the server by executing the binary.

## Restart

You will need to restart the `adamwestum` program any time you change Go code or any file in `template/`, but _not_ when you change files in `static/`.

# Deployment

Here are some details you'll need in order to deploy the app to a server and run it.

## /etc/adamwestum/adamwestum.gcfg

There is a config template file that you can copy.

```
[Host]
port = :80
path = /path/to/adamwestum
```

## iptables

```
iptables -I INPUT 1 -p tcp --dport 80 -j ACCEPT
```

## systemd config

Edit `/lib/systemd/system/adamwestum.service`:

```
[Unit]
Description=adamwestum web service
After=syslog.target network.target

[Service]
Type=simple
ExecStart=/path/to/adamwestum/adamwestum.linux

[Install]
WantedBy=multi-user.target
```

Then, make a symbolic link:

```
ln -s /lib/systemd/system/adamwestum.service /etc/systemd/system/adamwestum.service
```

And start/enable your service:

```
systemctl start adamwestum.service
systemctl enable adamwestum.service
```

## compile/upload/deploy script

I deploy with a script like this:

```
echo "Compiling"
cd adamwestum; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o adamwestum.linux; cd ..
echo "Uploading"
tar cvf adamwestum.tar adamwestum/adamwestum.linux adamwestum/template adamwestum/static adamwestum/quotes
ssh -l sirsean adamwestum.link mkdir -p adamwestum
ssh -l sirsean adamwestum.link rm adamwestum/adamwestum.linux
scp adamwestum.tar sirsean@adamwestum.link:adamwestum.tar
ssh -l sirsean adamwestum.link tar xvf adamwestum.tar
echo "Restarting"
ssh -l root adamwestum.link systemctl restart adamwestum.service
echo "Deployed"
```

Note that you will need to set up your system to cross-compile for your target environment. (I deploy to a 64-bit Linux server.)

