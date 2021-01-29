# Mars Rover Coding Challenge

Implementation of the mars rover challenge.

## Testing

Tests can be run from the root of the project with
```bash
go test ./...
```

## Running

First build the executable
```bash
go install github.com/robotlovesyou/mars/mars
```

Assuming your go /bin folder is in your path he executable can then be run by passing it the path to a config file. 
Two config files are included in the root of the project

```bash
mars example.txt
6, 4, NORTH
```

```bash
mars example_stops.txt
5, 4, NORTH STOPPED
```