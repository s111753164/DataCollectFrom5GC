package datacollectfrom5gc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

	// 通用方法：收集並保存 POST 請求的參數
func CollectAndSavePostParams(c *gin.Context, filePath string) {
	// 創建一個 map 來保存傳入參數
	params := make(map[string]interface{})

	// 獲取 URL 參數
	for _, param := range c.Params {
		params[param.Key] = param.Value
	}

	// 獲取查詢參數和表單參數
	for key, values := range c.Request.URL.Query() {
		if len(values) == 1 {
			params[key] = values[0]
		} else {
			params[key] = values
		}
	}

	// 獲取請求體中的 JSON 參數
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err == nil {
		for key, value := range body {
			params[key] = value
		}
	}

	// 將參數序列化為 JSON 格式
	jsonData, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	// 將 JSON 數據寫入文件
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write JSON to file"})
		return
	}
}
