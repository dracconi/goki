# goki!

Compile and
`goki -h`.

## Generate config
Config is in JSON and looks like this:
```json
   {
    "Links": ["http://boards.4chan.org/code/thread/thread.no",...],
    "Output": "output directory"
   }
```

You can generate it by `goki -make-conf config_file.json`.

## Download files
`goki [opts] config_file.json`,
available options are listed under `-h`.

# To-Do

* more options!
* make it multi-threaded
* better UI