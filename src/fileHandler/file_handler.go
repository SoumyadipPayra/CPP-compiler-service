package filehandler

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const ()

func Handle(ctx context.Context, userName, filePath string, fileContent []byte) ([]byte, error) {
	savedFilePath, err := saveCPPFile(ctx, userName, filePath, fileContent)
	if err != nil {
		return nil, err
	}
	executableFilePath, err := compileFile(ctx, savedFilePath)
	if err != nil {
		return nil, err
	}

	logsInBytes, err := executeFile(ctx, executableFilePath)

	if err != nil {
		return nil, err
	}

	return logsInBytes, nil
}

func saveCPPFile(_ context.Context, userName, fileSavingPath string, fileContent []byte) (string, error) {
	//prepend username to given filePath
	savePath := fmt.Sprintf("%s/%s", userName, fileSavingPath)

	//create the directories if needed
	if err := os.MkdirAll(filepath.Dir(savePath), 0777); err != nil {
		return "", fmt.Errorf("error while creating directories")
	}

	//write bytes in the file
	err := os.WriteFile(savePath, fileContent, 0644)
	if err != nil {
		return "", err
	}

	fmt.Printf("Data written to file successfully : %s\n", savePath)
	return savePath, nil
}

func compileFile(_ context.Context, filePath string) (string, error) {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	// Extract the directory and the file name without extension
	dir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)
	extension := filepath.Ext(fileName)

	// Check if the file has a .cpp extension
	if extension != ".cpp" {
		return "", fmt.Errorf("invalid file type: expected .cpp file")
	}

	// Remove the extension to get the base name of the file
	baseName := strings.TrimSuffix(fileName, extension)

	// The output binary will be saved in the same directory with the same base name
	outputPath := filepath.Join(dir, baseName)

	// Compilation flags for optimization
	compileFlags := []string{
		"-O2",            // Optimization level 2 for general speed optimization
		"-march=native",  // Optimize code for the host architecture
		"-flto",          // Link Time Optimization (LTO)
		"-Wall",          // Enable all common warnings
		"-Wextra",        // Enable extra warnings
		"-Werror",        // Treat warnings as errors
		"-o", outputPath, // Output binary file
		filePath, // Input C++ file
	}

	// Run the g++ command to compile the file
	cmd := exec.Command("g++", compileFlags...)

	// Run the command and capture any errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to compile: %s\n%s", err, string(output))
	}

	fmt.Printf("Compilation successful! Binary saved at: %s\n", outputPath)
	return outputPath, nil
}

func executeFile(_ context.Context, executableFilePath string) ([]byte, error) {
	// Check if the binary exists
	if _, err := os.Stat(executableFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("binary does not exist: %s", executableFilePath)
	}

	// Run the binary
	cmd := exec.Command(executableFilePath)

	// Capture the output (stdout and stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run binary: %s\n%s", err, string(output))
	}

	return output, nil
}
