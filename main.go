package main

type DataTypeStruct struct {
	Quarters       float64 `csv:"quarters"`
	Dimes          float64 `csv:"dimes"`
	Nickels        float64 `csv:"nickels"`
	Pennies        float64 `csv:"pennies"`
	ChangeNeeded   float64 `csv:"changeneeded"`
	AssertIs       bool    `csv:"assertisenoughchange"`
	TotalMoney     float64
	IsEnoughChange bool
}

func (d *DataTypeStruct) calcData() DataTypeStruct {
	d.TotalMoney += 0.25 * d.Quarters
	d.TotalMoney += 0.10 * d.Dimes
	d.TotalMoney += 0.05 * d.Nickels
	d.TotalMoney += 0.01 * d.Pennies
	d.IsEnoughChange = d.TotalMoney >= d.ChangeNeeded

	return *d
}

func main() {
	dataToTest, log := SetupTestData()

	for i, v := range dataToTest {
		if v.ToAssert != v.Data.calcData().IsEnoughChange {
			log.Error("", "index", i, "changeGiven", v.Data.TotalMoney, "changeNeeded", v.Data.ChangeNeeded, "isEnoughChange", v.Data.IsEnoughChange, "assertIsEnoughChange", v.ToAssert)
		} else {
			log.Info("", "index", i, "changeGiven", v.Data.TotalMoney, "changeNeeded", v.Data.ChangeNeeded, "isEnoughChange", v.Data.IsEnoughChange, "assertIsEnoughChange", v.ToAssert)
		}
	}
}
