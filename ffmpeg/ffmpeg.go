package ffmpeg

import "os/exec"

func ExecuteFFmpegCommand(args ...string) error {
	cmd := exec.Command("ffmpeg", args...)
	return cmd.Run()
}