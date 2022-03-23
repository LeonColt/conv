package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/leoncolt/conv/env"
)

type ConvService struct {
	Service
}

func (service *ConvService) DocxToPdf(file multipart.File) ([]byte, error) {
	name := uuid.New().String()
	sourcePath := filepath.Join(env.GetTempDir(), name)
	dst, err := os.Create(sourcePath)
	if err != nil {
		return []byte{}, err
	}
	defer os.Remove(sourcePath)
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return []byte{}, err
	}

	_, err = exec.Command(env.GetLibreOfficePath(), "--headless", "--convert-to", "pdf:writer_pdf_Export", sourcePath, "--outdir", env.GetTempDir()).Output()
	if err != nil {
		return []byte{}, err
	}

	resultPath := filepath.Join(env.GetTempDir(), fmt.Sprintf("%s.%s", name, "pdf"))

	defer os.Remove(resultPath)

	res, err := os.ReadFile(resultPath)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}

func (service *ConvService) DocxToHtml(file multipart.File) ([]byte, error) {
	name := uuid.New().String()
	sourcePath := filepath.Join(env.GetTempDir(), name)
	dst, err := os.Create(sourcePath)
	if err != nil {
		return []byte{}, err
	}
	defer os.Remove(sourcePath)
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return []byte{}, err
	}

	_, err = exec.Command(env.GetLibreOfficePath(), "--headless", "--convert-to", "html:HTML:EmbedImages", sourcePath, "--outdir", env.GetTempDir()).Output()
	if err != nil {
		return []byte{}, err
	}

	resultPath := filepath.Join(env.GetTempDir(), fmt.Sprintf("%s.%s", name, "html"))

	defer os.Remove(resultPath)

	res, err := os.ReadFile(resultPath)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}

func (service *ConvService) HtmlToDocx(file multipart.File) ([]byte, error) {
	name := uuid.New().String()
	sourcePath := filepath.Join(env.GetTempDir(), name)
	dst, err := os.Create(sourcePath)
	if err != nil {
		return []byte{}, err
	}
	defer os.Remove(sourcePath)
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return []byte{}, err
	}

	_, err = exec.Command(env.GetLibreOfficePath(), "--headless", "--infilter=writerglobal8_HTML", "--convert-to", `docx:MS Word 2007 XML`, sourcePath, "--outdir", env.GetTempDir()).Output()
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	resultPath := filepath.Join(env.GetTempDir(), fmt.Sprintf("%s.%s", name, "docx"))

	defer os.Remove(resultPath)

	res, err := os.ReadFile(resultPath)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	return res, nil
}

func (service *ConvService) HtmlToPdf(file multipart.File) ([]byte, error) {
	name := uuid.New().String()
	sourcePath := filepath.Join(env.GetTempDir(), name)
	dst, err := os.Create(sourcePath)
	if err != nil {
		return []byte{}, err
	}
	defer os.Remove(sourcePath)
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return []byte{}, err
	}

	_, err = exec.Command(env.GetLibreOfficePath(), "--headless", "--infilter=writerglobal8_HTML", "--convert-to", "pdf:writer_web_pdf_Export", sourcePath, "--outdir", env.GetTempDir()).Output()
	if err != nil {
		return []byte{}, err
	}

	resultPath := filepath.Join(env.GetTempDir(), fmt.Sprintf("%s.%s", name, "pdf"))

	defer os.Remove(resultPath)

	res, err := os.ReadFile(resultPath)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}
