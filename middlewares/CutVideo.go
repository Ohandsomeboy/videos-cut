package middlewares

import (
	"bytes"
	"os/exec"
)

// func ClipVideo(input []byte, startTime, duration string) (*bytes.Buffer, error) {
func ClipVideo(input bytes.Buffer, startTime, duration string) (*bytes.Buffer, error) {
	// 初始化一个bytes.Buffer来保存剪辑后的视频流
	in := input.Bytes()
	output := bytes.NewBuffer([]byte{})

	// 构建 FFmpeg 的命令行参数
	cmdArgs := []string{
		"-f", "mp4",
		"-i", "pipe:0",
		"-ss", startTime,
		"-t", duration,
		"-c:v", "libx264",
		"-preset", "ultrafast",
		"-c:a", "copy",
		"-f", "mp4",
		"pipe:1",
	}

	// 创建 ffmpeg 命令的进程
	cmd := exec.Command("E:/depand/ffmpeg/bin/ffmpeg.exe", cmdArgs...)
	//cmd := exec.Command("cmd", "/C", "E:/depand/ffmpeg/bin/ffmpeg.exe", cmdArgs...)
	cmd.Stdin = bytes.NewReader(in)
	cmd.Stdout = output

	// 执行 FFmpeg 命令行操作
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return output, nil
}

//package middlewares

//
//import (
//	"bytes"
//	"io"
//	"os/exec"
//)
//
//func ClipVideo(videoStream io.Reader, startTime, duration string) (bytes.Buffer, error) {
//	// 初始化一个bytes.Buffer来保存剪辑后的视频流
//	output := bytes.Buffer{}
//
//	// 构建 FFmpeg 的命令行参数
//	cmdArgs := []string{
//		"-i", "pipe:0",
//		"-ss", startTime,
//		"-t", duration,
//		"-c", "copy",
//		"pipe:1",
//	}
//
//	// 创建 ffmpeg 命令的进程
//	cmd := exec.Command("E:/depand/ffmpeg/bin/ffmpeg.exe", cmdArgs...)
//	cmd.Stdin = videoStream
//	cmd.Stdout = &output
//
//	// 执行 FFmpeg 命令行操作
//	if err := cmd.Run(); err != nil {
//		return bytes.Buffer{}, err
//	}
//
//	// 将剪辑后的视频流写入新的缓冲区
//	var clippedVideo bytes.Buffer
//	if _, err := io.Copy(&clippedVideo, &output); err != nil {
//		return bytes.Buffer{}, err
//	}
//
//	return clippedVideo, nil
//}
