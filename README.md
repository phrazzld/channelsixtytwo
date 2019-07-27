# Channel Sixty Two

Turn your command line into a VLC remote control.

## Setup
*channelsixtytwo* makes some strong assumptions about what file types you want to play and where your VLC installation is located. More thorough user customization options are on the way, but for now make sure your media player is at the absolute path `/Applications/VLC.app/Contents/MacOS/VLC`, and that all the media you want to play is in one of the following file formats: MKV, AVI, MP4, or M4V.

Then `go get github.com/phrazzld/channelsixtytwo`

## Usage
`channelsixtytwo [p...] x`
Where `[p...]` is any number of paths and `x` is the number of files to randomly load from those paths into a VLC playlist.

### Example
`channelsixtytwo /Volumes/External_HDD/TV/30_Rock /Volumes/External_HDD/TV/South_Park 10`
The above command will randomly select ten MKV, AVI, MP4, or M4V files from the `.../30_Rock` and `.../South_Park` directories, launch VLC, and play those ten files as a playlist.

## Testing
```
go test
go test -bench .
```

## LICENSE
[MIT](https://opensource.org/licenses/MIT)
