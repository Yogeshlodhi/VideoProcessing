# Video Processing CLI Toolkit

This project provides a set of command-line tools for video processing tasks, including video format conversion, compression, and trimming. Built using Go, this toolkit is optimized for high performance and simplicity.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- **Trim Video**: Trim videos by specifying start and end times.
- **Convert Video Format**: Convert videos to different formats such as MP4, AVI, etc.
- **Compress Video**: Compress videos to reduce file size without significant quality loss.
- **Batch Processing**: Process multiple video files at once.
- **Simple Command-Line Interface**: Easy-to-use commands for video manipulation.
- **Cross-Platform**: Build and use the toolkit on various platforms (Windows, macOS, Linux).
- **Docker Support**: Dockerize the tool for easy deployment and portability.

## Installation

Follow these steps to set up the project locally:

- **Clone the repository:**

  ```bash
  git clone https://github.com/Yogeshlodhi/VideoProcessing.git
  cd VideoProcessing

- **Install Go dependencies: (Make sure you have Go installed on your system)**
  ```bash
   go mod tidy

## Usage
After building the project, you can use the toolkit directly from the command line.
- **Know the commands**
  ```bash
    - go run main.go --help

    A CLI tool built in golang to alter/edit, change video options

    Usage:
      videoprocess [flags]
      videoprocess [command]

    Available Commands:
      completion  Generate the autocompletion script for the specified shell
      compress    Compress video to reduce the files size
      convert     Convert video to another format like mp4, mkv
      download    Download files faster than before
      help        Help about any command
      trim        Trim video from start time to end time

- **Download File**
  ```bash
  ./VideoProcessing download --input <input_file> --output <output_file> 
  ```
    - input : The path to the input file to be downloaded.
    - output : The path for the output downloaded file.

- **Convert Video Format**
  ```bash
  ./VideoProcessing convert --input <input_file> --output <output_file> --format <desired_format>
  ```
    - input : The path to the input video file.
    - output : The path for the output trimmed video file.
    - format : The format you want to convert the video to (e.g., mp4, avi).

- **Compress Video**
  ```bash
  ./VideoProcessing compress --input <input_file> --output <output_file> --quality <quality_percentage>
  ```
    - input : The path to the input video file.
    - output : The path for the output trimmed video file.
    - quality : The desired quality percentage (e.g., 80 for 80% quality).

- **Some Test Commands**
  ```bash
    - go run main.go convert "C:/Users/Yogesh Kumar/Desktop/testmp4" mkv
    - go run main.go compress "C:/Users/Yogesh Kumar/Desktop/test.mp4" output.mp4

    - go run main.go download "file_link" "C:\Users\Yogesh Kumar\Desktop\Projects\VideoProcessing\test.mkv"
   

## Contributing
Contributions are always welcome!
1. Fork the repository.
2. Create a new branch for your feature/bugfix.
3. Implement your changes.
4. Run tests (if applicable).
5. Create a pull request with a description of your changes.
