package main

type DataTypeStruct struct {
	Quarters     float64 `csv:"quarters"`
	Dimes        float64 `csv:"dimes"`
	Nickels      float64 `csv:"nickels"`
	Pennies      float64 `csv:"pennies"`
	ChangeNeeded float64 `csv:"changeneeded"`
	AssertIs     bool    `csv:"assertisenoughchange"`
}

type DataMap struct {
	TotalMoney     float64
	ChangeNeeded   float64
	IsEnoughChange bool
}

func calcData(data DataTypeStruct) DataMap {
	mapData := DataMap{}
	mapData.ChangeNeeded = data.ChangeNeeded
	mapData.TotalMoney += 0.25 * data.Quarters
	mapData.TotalMoney += 0.10 * data.Dimes
	mapData.TotalMoney += 0.05 * data.Nickels
	mapData.TotalMoney += 0.01 * data.Pennies
	mapData.IsEnoughChange = mapData.TotalMoney >= mapData.ChangeNeeded

	return mapData
}

func main() {
	dataToTest, log := SetupTestData()

	for i, v := range dataToTest {
		returned := calcData(v.Data)
		if v.ToAssert != returned.IsEnoughChange {
			log.Error("", "index", i, "changeGiven", returned.TotalMoney, "changeNeeded", returned.ChangeNeeded, "isEnoughChange", returned.IsEnoughChange, "assertIsEnoughChange", v.ToAssert)
		} else {
			log.Info("", "index", i, "changeGiven", returned.TotalMoney, "changeNeeded", returned.ChangeNeeded, "isEnoughChange", returned.IsEnoughChange, "assertIsEnoughChange", v.ToAssert)
		}
	}
}
