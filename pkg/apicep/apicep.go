package apicep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/michelpessoa/segundodesafiogoexpert/internal/dto"
)

func FetchApiCep(cep string, ch chan<- string) {
	cep = cep[:5] + "-" + cep[5:]
	start := time.Now()
	url := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	var apiCepResponse dto.ApiCepResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	err = json.Unmarshal(body, &apiCepResponse)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("URL: %s, Tempo decorrido: %d ms, Resposta: %+v", url, int(secs*1000), apiCepResponse)
}
