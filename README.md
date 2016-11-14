[![Stories in Ready](https://badge.waffle.io/rodesousa/luz-lantern.png?label=ready&title=Ready)](https://waffle.io/rodesousa/luz-lantern)
# luz-lantern

## Dependancies
```bash
go get -v github.com/spf13/cobra/cobra
go get -v github.com/Sirupsen/logrus
```

## Conf
Example
```
cmd:
    - user:
        name: aze

    - user:
        name: lala

    - user:
        name : ippon

    - user:
        name : root

    - ping :
        url : "google.com"

    - ping :
        url : "localhost"
        expected : false

    - curl:
        url: "http://www.google.fr"
        name: "google"
```

## Commande Line

Run
```
lantern run 
lantern run -c conf.yaml
```

Mode server (port 8080)
```
lantern server start &
lantern server start -c conf.yaml &
lantern server stop | status
```

Flags
```
-c, --config string    conf file (default "conf.yaml")
-d, --debug            show debug message
-h, --help             help for lantern
    --logfile string   log file output (default is current path)
-o, --off              disable out console log
```

## luz with vim

- Check install https://github.com/rodesousa/vim_conf.git
- In vim, :GoInstallBinaries
