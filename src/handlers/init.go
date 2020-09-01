package handlers

/*
func init() {
	// 判断Redis有缓存的Buyers数据
	match := src.Store.Rds(common.FishingHandOverRecordStore)
	if exists := match.Exists("Buyers").Val(); exists {
		if err := json.Unmarshal([]byte(match.Get("Buyers").Val()), &buyers); err != nil {
			logger.FatalF("json unmarshal buyers error: %v", err)
		} // 重新统计入库
		for area, dates := range buyers {
			for date, users := range dates {
				match.HSet(area, date, len(users))
			}
		}

	} else { // 没有缓存数据则初始化Buyers
		buyers = make(map[string]map[string]map[int]int)
		buyers["Buyers11"] = make(map[string]map[int]int)
		buyers["Buyers12"] = make(map[string]map[int]int)
		buyers["Buyers13"] = make(map[string]map[int]int)
		buyers["Buyers14"] = make(map[string]map[int]int)
		buyers["Buyers21"] = make(map[string]map[int]int)
		buyers["Buyers22"] = make(map[string]map[int]int)
		buyers["Buyers23"] = make(map[string]map[int]int)
		buyers["Buyers31"] = make(map[string]map[int]int)
		buyers["Buyers32"] = make(map[string]map[int]int)
		buyers["Buyers33"] = make(map[string]map[int]int)
		buyers["Buyers34"] = make(map[string]map[int]int)
		buyers["Buyers35"] = make(map[string]map[int]int)
		buyers["TotalBuyers"] = make(map[string]map[int]int)
	}
	logger.Info("fishing buyers load successful")
}
*/
