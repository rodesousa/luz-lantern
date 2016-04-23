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
Build the program :
```bash
cd $GOPATH/src/github.com/rodesousa/lantern
go install .
```
Enjoy !
```bash
$GOPATH/bin/lantern run `yaml__file`
```

### luz with vim

- Check install https://github.com/rodesousa/vim_conf.git
- In vim, :GoInstallBinaries
