# pdf-tools

可以，下面是一个用 Go 实现的 **将 PDF 按页转换为 JPG 图片** 的完整示例。

---

### ✅ 使用的工具和库：

1. `github.com/gen2brain/go-fitz`：Go 的 MuPDF 封装，用于读取 PDF 页面并转成图像。
2. `image/jpeg`：Go 标准库中的 JPEG 编码器。

---

### 🛠️ 安装依赖：

```bash
go get github.com/gen2brain/go-fitz
```

你需要本地安装 **MuPDF** 库（这是 `go-fitz` 的底层依赖）：

#### macOS 安装 MuPDF：

```bash
brew install mupdf-tools
```

---

### 📄 示例代码：将 PDF 每页保存为一张 JPG 图片

```go
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
	pdfPath := "example.pdf"     // 替换成你的 PDF 文件路径
	outputDir := "output_images" // 输出文件夹

	// 打开 PDF
	doc, err := fitz.New(pdfPath)
	if err != nil {
		log.Fatalf("无法打开 PDF: %v", err)
	}
	defer doc.Close()

	// 创建输出目录
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("无法创建输出目录: %v", err)
	}

	// 遍历每一页
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			log.Printf("读取第 %d 页失败: %v", n, err)
			continue
		}

		outputPath := filepath.Join(outputDir, fmt.Sprintf("page_%03d.jpg", n+1))
		outFile, err := os.Create(outputPath)
		if err != nil {
			log.Printf("无法创建图片文件: %v", err)
			continue
		}

		// 保存为 JPG
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
		if err != nil {
			log.Printf("保存图片失败: %v", err)
		}

		outFile.Close()
		fmt.Println("生成图片:", outputPath)
	}
}
```

---

### 📦 输出结果：

假设你有一个名为 `example.pdf` 的文件，上述代码会生成：

```
output_images/
├── page_001.jpg
├── page_002.jpg
├── ...
```

---

### 📌 注意事项：

* 转换速度取决于页面复杂度和分辨率（默认 72 DPI 左右）。
* 如果你需要更高分辨率，可以用其他底层工具如 `pdftoppm` 或 `pdfium`。

---

需要我帮你加一个命令行参数支持（例如传 PDF 路径）或者把它打包成可执行 CLI 工具吗？
