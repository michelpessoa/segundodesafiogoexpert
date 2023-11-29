package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/michelpessoa/segundodesafiogoexpert/internal/dto"
)

func FetchViaCep(cep string, ch chan<- string) {
	start := time.Now()
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	var viaCepResponse dto.ViaCepResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("URL: %s, Tempo decorrido: %d ms, Resposta: %s", url, int(secs*1000), viaCepResponse)
}
