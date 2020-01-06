# Muti-Magick

## Processing images with concurrency in goroutines with ImageMagick

**Problem:** [ImageMagick](https://imagemagick.org/index.php) is a single-threaded application used for image manipulation. This makes batch processing images time-consuming.

**Solution:** Using Golang's [Concurrency](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3) to make ImageMagick a "muti-threaded" application.

| Test Run        | Images | Time Elaped   | Test Command                           |
| --------------- | ------ | ------------- | -------------------------------------- |
| Concurrency     | 100    | 2.845795397s  | `cd concurrency && go run main.go`     |
| Non-Concurrency | 100    | 14.862155041s | `cd non-concurrency && go run main.go` |

![gophereartrumpet](https://user-images.githubusercontent.com/18271248/71789087-27857100-2fee-11ea-8e10-3c9ad839be3c.jpg)
