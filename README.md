# iTunes CLI
Command line interface for control iTunes

![demo](./demo.gif)  

---  

## Description  
You can control iTunes operations from command line.  

## Equipments
- macOS Sierra or later
- iTunes 12.5.4 or later
- Go

## Installation
``` sh
$ go get github.com/ktr0731/itunes-cli/itunes
```

Or GitHub Releases are also available.

## Usage
Each sub-commands can use shorter name.  
Please see `$ itunes help`.  

### Play
Play current selected music.  
``` sh
$ itunes play
```

Play music that name passed by a argument.
``` sh
$ itunes play reunion
```

### Pause
Pause current music
``` sh
$ itunes pause
```

Replay current music or play previous music.  
If playing music's current played time is a few second, play previous music.  
Other than that, replay it music from beginning.  
``` sh
$ itunes back
```

### Next/Previous
Play next music.  
``` sh
$ itunes next
```

Play previous music.  
``` sh
$ itunes prev
```

### Volume
Change volume in iTunes.  
``` sh
$ itunes vol 20
```

### Find music
You can find musics by a fuzzy-finder.
``` sh
$ itunes find
$ itunes find plist
```

## License
Please see LICENSE.
