# matching-system

## Project Structure
```
matching-system/
|-- api/
|   |-- handlers.go          # API handlers (AddSinglePersonAndMatch, RemoveSinglePerson, QuerySinglePeople)
|-- doc/
|   |-- openapi.json         # API documentation
|-- model/
|   |-- person.go            # Defines the Person model and matching logic
|-- store/
|   |-- memory_store.go      # In-memory data store operations
|-- test/
    |-- api_test.go          # Unit tests for API handlers
    |-- store_test.go        # Unit tests for in-memory store
|-- utils/
|   |-- utils.go             # Utility functions
|-- main.go                  # Main application file
|-- Dockerfile               # Docker configuration
|-- README.md                # Project documentation
```

## System Design
### 系統概述
* 本系統旨在為用戶提供一個類似於 Tinder 的匹配系統，支援新增用戶、移除用戶和查詢匹配對象。系統透過 RESTful API 提供服務，並採用記憶體存儲結構來管理用戶數據。

### 系統架構
* 本系統由以下主要組件構成：
    * HTTP 伺服器：使用 Gin 框架建立，處理外部請求。
    * API 處理層：處理具體的業務邏輯。
    * 記憶體存儲：暫存用戶數據，支援快速讀寫操作。
    * 匹配引擎：實現匹配邏輯。

### API 設計
* 新增單身人士並尋找匹配
    * 路徑：/add_single_person_and_match
    * 方法：POST
    * 輸入：包含 name, height, gender, numberOfDates 的 JSON
    * 處理：新增單身人士至系統，並進行匹配。
    * 回應：新增的單身人士資訊和匹配對象。
* 移除單身人士
    * 路徑：/remove_single_person/{name}
    * 方法：DELETE
    * 處理：根據名稱移除指定的單身人士。
    * 回應：操作結果。
* 查詢單身人士
    * 路徑：/query_single_people
    * 方法：GET
    * 參數：n (返回的最多人數)
    * 處理：返回最多 n 個可能的匹配人選。
    * 回應：單身人士列表。

### 時間複雜度分析
* 新增單身人士並尋找匹配：O(N)，N 為當前系統中用戶的數量。
* 移除單身人士：O(1)，直接根據名稱在 map 中查找並移除。
* 查詢單身人士：O(N)，需要對所有用戶進行篩選。

### 數據結構
* 使用 Go 語言內建的數據結構來存儲用戶數據，並進行有效的匹配操作。

## TBD Tasks
### 用戶驗證和安全性
* 用戶認證機制，如 JWT（JSON Web Tokens）或 OAuth。
* 確保數據傳輸的安全性，例如通過 HTTPS。
* 考慮加入防止常見網絡攻擊（如 SQL 注入、跨站腳本（XSS）等）的安全措施。
### 數據持久化
* 目前系統使用記憶體存儲，需要考慮數據持久化方案，例如引入 MySQL 或 MongoDB。
* 實現數據庫連接和管理邏輯。
### 系統性能和優化
* 針對大規模數據進行性能測試和優化。
* 評估並優化匹配算法的效率。
### 容錯和可擴展性
* 設計和實現錯誤處理和日誌記錄機制。
* 考慮系統的橫向擴展能力。
### 用戶界面和體驗
* 如果計劃開發前端界面，則需要設計和實現用戶界面。
### API 文檔和版本控制
* 持續更新和維護 API 文檔。
* 實施 API 版本控制策略，以支持向後兼容。
### 測試策略
* 擴展測試覆蓋率，包括單元測試、集成測試和壓力測試。
* 實現持續集成（CI）和持續部署（CD）流程。
### 國際化和本地化：
* 如果服務範圍包括多個國家/地區，則需要考慮多國語言支持。