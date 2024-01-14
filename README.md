# Friendville Town

The beginnings of the friendville.us town website.

Right now the site consists of static assets and plain HTML/CSS/JS located in the `static/`
directory.

## Cloning this repo

To clone this repo to your machine, run the following:

```
git clone https://github.com/lemonase/friendville-us.git
```

## Running a local dev server

There is a simple Go http server that serves the files in `./static/`, you can run by doing

```
go run static-http-server.go
```

If you have python installed, you can spin up a server by running:

```
python3 -m http.server --directory ./static/
```

## What's next

I have been playing around with UI components from a library called
[shoelace](https://github.com/shoelace-style/shoelace).
Likely will be sticking with this for the remaining development of the site.

Next big thing will be making an interactive map with pins.
