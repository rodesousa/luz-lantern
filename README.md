[![Stories in Ready](https://badge.waffle.io/rodesousa/luz-lantern.png?label=ready&title=Ready)](https://waffle.io/rodesousa/luz-lantern)
# luz-lantern

## Conf

### Install

Download source :
```bash
git clone https://github.com/rodesousa/luz-lantern.git
```
Configure $GOPATH env :
```bash 
cd luz-lantern
ln -rs github.com $GOPATH/src/
```
Install dependencies :
```bash
go get -v github.com/spf13/cobra/cobra
```
Download and add logrus to your $GOPATH
```bash
cd $GOPATH/src/github.com
git clone https://github.com/Sirupsen/logrus.git
```
Build the program :
```bash
cd $GOPATH/src/github.com/rodesousa/lantern
go install .
```
## Run

```
lantern run `yaml__file`
```

Mode server (port 8080)
```
lantern run `yaml__file` -s &
```

Debug
```
lantern run `yaml__file` -d
```

### luz with vim

- Check install https://github.com/rodesousa/vim_conf.git
- In vim, :GoInstallBinaries
