# 测试指南

## 环境准备

确保已安装 Go 1.22+ 环境。

## 运行所有测试

```bash
cd server_go
go mod tidy
go test -v ./...
```

## 单独运行 Service 层测试

```bash
go test -v ./service/...
```

## 单独运行 DAO 层测试

```bash
go test -v ./dao/...
```

## 测试覆盖范围

### Service 层测试 (product_service_test.go)
- ✅ 四种排序策略测试：价格升序、价格降序、销量优先、好评率优先
- ✅ 价格统计计算测试（最低价、均价）
- ✅ 空边界情况处理测试

### DAO 层测试 (product_dao_test.go)
- ✅ 商品创建与查询测试
- ✅ 按类目筛选测试
- ✅ 多条件组合筛选测试（价格区间、品牌、销量）
- ✅ 批量插入测试
- ✅ 使用 SQLite 内存数据库，无需外部数据库依赖

## 测试输出示例

```
=== RUN   TestSortProducts
=== RUN   TestSortProducts/sort_by_price_asc
=== RUN   TestSortProducts/sort_by_price_desc
=== RUN   TestSortProducts/sort_by_sales
=== RUN   TestSortProducts/sort_by_rating
--- PASS: TestSortProducts (0.00s)
=== RUN   TestCalculatePriceStats
--- PASS: TestCalculatePriceStats (0.00s)
=== RUN   TestCalculatePriceStatsEmpty
--- PASS: TestCalculatePriceStatsEmpty (0.00s)
PASS
ok      compare_prices/server_go/service    0.001s
```
