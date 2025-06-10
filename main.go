// main.go 示例：将 PDF 每一页导出为 JPG 图片
package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func main() {
	pdfPath := "example.pdf"     // PDF 文件路径（请根据实际情况修改）
	outputDir := "output_images" // 输出图片的文件夹

	// 打开 PDF 文档
	doc, err := fitz.New(pdfPath)
	if err != nil {
		log.Fatalf("无法打开 PDF: %v", err) // 打开失败则终止程序
	}
	defer doc.Close() // 程序结束时关闭文档

	// 创建输出目录（如果不存在则自动创建）
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("无法创建输出目录: %v", err)
	}

	// 遍历 PDF 的每一页
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n) // 获取第 n 页的图片
		if err != nil {
			log.Printf("读取第 %d 页失败: %v", n, err) // 读取失败则跳过
			continue
		}

		// 构造输出图片的路径，格式如 page_001.jpg
		outputPath := filepath.Join(outputDir, fmt.Sprintf("page_%03d.jpg", n+1))
		outFile, err := os.Create(outputPath) // 创建图片文件
		if err != nil {
			log.Printf("无法创建图片文件: %v", err)
			continue
		}

		// 将图片编码为 JPG 格式并保存到文件
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
		if err != nil {
			log.Printf("保存图片失败: %v", err)
		}

		outFile.Close()                  // 关闭文件
		fmt.Println("生成图片:", outputPath) // 输出生成结果
	}
}
