# verscmp

A lightweight Go library for comparing version strings (e.g., `10.0.1`), supporting natural number comparison (like `1.10 > 1.2`). Suitable for version control, dependency management, and upgrade detection.

## 📦 Installation

```bash
go get github.com/phantom-cube/verscmp
```

🧪 Example
```bash
package main

import (
    "fmt"
    "github.com/phantom-cube/verscmp"
)

func main() {
    res, err := verscmp.VersionCompare("10.0.1", "2.0.1")
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("Result:", res) // RESULT_EQ=1, RESULT_LT=2, RESULT_GT=3
}
```

✅ Features
- Natural number comparison (e.g., 1.10 > 1.2)
- Input validation (e.g., reject version starting/ending with .)
- Comparison results: RESULT_EQ (equal), RESULT_LT (less than), RESULT_GT (greater than)
- Supports mixed numeric and alpha versions (e.g., 1.0a < 1.0b)

📚 Future Roadmap
- Support version ranges (e.g., >=1.0.0 <2.0.0)
- Support version prefixes (e.g., v1.0.0)
🤝 Contributing
Feel free to open issues or submit PRs to improve this lightweight version comparison utility.

📄 License  
MIT License