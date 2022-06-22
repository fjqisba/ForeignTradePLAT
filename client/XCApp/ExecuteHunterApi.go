package XCApp

func (this *XCApp)executeHunterApi()  {
	rowCount := this.list_Csv.GetCount_AD()
	for i:=0;i<rowCount;i++{
		var records []string
		for j:=0;j<24;j++{
			records = append(records,this.list_Csv.GetItemText(i,j))
		}
		//domain := records[8]
		hunterData := records[7]
		if hunterData != ""{
			continue
		}
		//emailList := api.SearchDomain(domain)
		//emailStr, _ := json.Marshal(emailList)
		//if this.list_Csv.SetItemText(i,7,string(emailStr)) == false{
		//	log.Println("写入Hunter数据失败")
		//}
	}
}
