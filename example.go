package main

//func (s *Schedule) ExportExcel(ctx *context.Context, seances []*pbCatalogSchedule.Schedule) ([]byte, error) {
//	f := excelize.NewFile()
//
//	for i, seance := range seances {
//		row := fmt.Sprintf("%d", i)
//
//		f.SetCellValue("Sheet1", "A"+row, row)
//		f.SetCellValue("Sheet1", "B"+row, seance.Seance.Timeframe.Start)
//		f.SetCellValue("Sheet1", "C"+row, seance.Seance.Timeframe.End)
//		f.SetCellValue("Sheet1", "D"+row, seance.Seance.Timeframe.Interval)
//		f.SetCellValue("Sheet1", "E"+row, seance.Movie.Properties.Duration)
//		f.SetCellValue("Sheet1", "F"+row, seance.Movie.Name)
//	}
//
//	if b, err := f.WriteToBuffer(); err != nil {
//		return nil, err
//	} else {
//		return b.Bytes(), nil
//	}
//}
