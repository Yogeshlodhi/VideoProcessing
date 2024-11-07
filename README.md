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
  git clone https://github.com/Yogeshlodhi/VideoProcessingCLI.git
  cd VideoProcessingCLI

- **Install Go dependencies: (Make sure you have Go installed on your system)**
  ```bash
   go mod tidy

## Usage
After building the project, you can use the toolkit directly from the command line.
- **Trim Video**
  ```bash
  ./video-tool trim --input <input_file> --start <start_time> --end <end_time> --output <output_file>
  ```
    - input : The path to the input video file.
    - start : The start time (in seconds) for the trim.
    - end : The end time (in seconds) for the trim.
    - output : The path for the output trimmed video file.

- **Convert Video Format**
  ```bash
  ./video-tool convert --input <input_file> --output <output_file> --format <desired_format>
  ```
    - input : The path to the input video file.
    - output : The path for the output trimmed video file.
    - format : The format you want to convert the video to (e.g., mp4, avi).

- **Compress Video**
  ```bash
  ./video-tool compress --input <input_file> --output <output_file> --quality <quality_percentage>
  ```
    - input : The path to the input video file.
    - output : The path for the output trimmed video file.
    - quality : The desired quality percentage (e.g., 80 for 80% quality).

## Contributing
Contributions are always welcome!
1. Fork the repository.
2. Create a new branch for your feature/bugfix.
3. Implement your changes.
4. Run tests (if applicable).
5. Create a pull request with a description of your changes.
