package futils

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path"
)

func GetOutputDir(filePath *string) (*string, error) {
	fileInfo, err := os.Stat(*filePath)
	if err == nil {
		if fileInfo.IsDir() {
			return filePath, nil
		} else {
			// 获取选择文件父目录信息
			dir, _ := path.Split(*filePath)
			return &dir, nil
		}
	}
	if os.IsNotExist(err) {
		mkerr := os.Mkdir(*filePath, 0700)
		if mkerr != nil {
			return nil, mkerr
		}
	}
	return nil, err
}

func CheckH265Format(filePath *string) (bool, error) {
	if _, err := os.Stat(*filePath); err == nil {
		// 获取ffprobe的路径位置
		cmdPath, err := exec.LookPath("ffprobe")
		if err != nil {
			return false, err
		}
		// 使用命令检查视频格式
		videoCmd := exec.Command(cmdPath, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", *filePath)
		var outputString bytes.Buffer
		videoCmd.Stdout = &outputString
		err = videoCmd.Run()
		if err != nil {
			return false, err
		}
		var data map[string]interface{}
		json.Unmarshal(outputString.Bytes(), &data)
		var videoType string
		streamData := data["streams"]
		mapArray := streamData.([]interface{})
		for _, v := range mapArray {
			entry := v.(map[string]interface{})
			if entry["codec_type"] != nil && entry["codec_type"] == "video" {
				videoType = entry["codec_name"].(string)
				break
			}
		}
		if videoType == "hevc" || videoType == "h265" {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, err
	}
}

func ConvertH265toH264(srcFilePath string, destDirPath string) (string, error) {
	srcFile, err := os.Stat(srcFilePath)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(destDirPath); err != nil {
		return "", err
	}
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return "", err
	}
	fileName := srcFile.Name()
	fileName = "convert_" + fileName
	outFilePath := path.Join(destDirPath, fileName)
	convertCmd := exec.Command(cmdPath, "-i", srcFilePath, "-x265-params", "crf=25", outFilePath)
	var out bytes.Buffer
	convertCmd.Stdout = &out
	execErr := convertCmd.Run()
	if execErr != nil {
		return "", execErr
	}
	return out.String(), nil
}
