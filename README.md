# Muti-Magick

Processing image using ImageMagick with gorouties to run image processsing in parallel
<br>

<img src="https://i.ibb.co/hLJJxLC/gophercises-lifting.gif" alt="drawing" width="300"/>

## Running ImageMagick without goroutine

`Processsing took 24.7486193s`

## Running ImageMagick with goroutine

`Processsing took 3.9396034s!`

<br>

<img src="https://i.ibb.co/0YbWR7H/mindblown.gif" alt="drawing" width="300"/>

## Running the Tests

### The slow test without gorouties

- `cd slow/`
- `go run slow.go`
- This command is compressing `doge.jpg` in half 100 times
- You will now see a directory a halfed doge's

#### You should see a printout of `Processsing took 24.7486193s`

<img src="https://i.ibb.co/8sZJvZf/giphy.gif" alt="drawing" width="300"/>

<br>

### The fast test

- `cd ..`
- `go run main.go`
- This command is also compressing `doge.jpg` in half 100 times
- You will now see another directory a halfed doge's

#### You should see a printout of `Processsing took 3.9396034s!`

<img src="https://i.ibb.co/jRzWfrn/giphy-1.gif" alt="drawing" width="400"/>
