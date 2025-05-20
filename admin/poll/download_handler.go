package poll

/*
func HandleDownload(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Запрос GET на /admin/download")

	LogJsonRecieved(req.Body)

	var r model.DownloadRequest
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	c, e := req.Cookie("sessionId")
	if e != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	a, err := db.GetAnswers(r.Link, c.Value)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}

	res, err := db.GetPoll(r.Link)
	if err != nil {
		fmt.Println("Опрос не найден")
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.FileType{
	case model.EXCEL{

	}
	case model.CSV{

	}

	}
}

func CreateCSV(poll *model.Poll, answers *[]model.Answer) (*os.File, error) {
	writer, file, err := createCSVWriter("temp")
	if err != nil {
		fmt.Println("Error creating CSV writer:", err)
		return nil, err
	}
	defer file.Close()

	var records  [][]string
	total := float32(len(*answers))

	for i,ans := range *answers {
		if(poll.Questions[i].Type == model.Text){
			continue
		}
		if(ans.Select!= nil){

			p := float32(*ans.Select)
			p *= 100
			p /= total

			var rec []string
			rec = append(rec, poll.Questions[i].Text)
			rec = append(rec, string(p))
		}
	}

	for _, record := range records {
		writeCSVRecord(writer, record)
	}
	// Flush the writer and check for any errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
	}
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	writer := csv.NewWriter(f)
	return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing record to CSV:", err)
	}
}
*/
