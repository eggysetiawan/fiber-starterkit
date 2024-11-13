package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/eggysetiawan/fiber-starterkit/config"
	"github.com/eggysetiawan/fiber-starterkit/errs"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
)

type BigDataRepositoryAPI struct {
	client *http.Client
}

func NewBigDataRepositoryAPI() *BigDataRepositoryAPI {
	return &BigDataRepositoryAPI{&http.Client{}}
}

func (api *BigDataRepositoryAPI) GetPostalCode(c context.Context, pc int, cpc chan domain.PostalCode, wg *sync.WaitGroup, errCh chan *errs.AppError) {
	defer wg.Done()

	raw := map[string]int{
		"kodePos": pc,
	}

	bodyReq, err := json.Marshal(raw)
	if err != nil {
		errCh <- errs.NewUnexpectedError("failed to marshall body request in GetPostalCode: " + err.Error())
		return
	}

	url := fmt.Sprintf("%s%s", config.AppConfig.GetString("BRIGATE_BASE_URL"), "/gateway/apiDataHub/1.0/kodepos")
	fmt.Println("36: ", url, string(bodyReq))
	apiReq, err := http.NewRequestWithContext(c, "POST", url, bytes.NewBuffer(bodyReq))
	if err != nil {
		errCh <- errs.NewUnexpectedError("failed to making request api GetPostalCode: " + err.Error())
		return
	}

	apiReq.Header.Set("Content-Type", "application/json")
	// apiReq.Header.Set("X-DATAHUB-USER", config.AppConfig.GetString("X_DATAHUB_USER"))
	// apiReq.Header.Set("X-DATAHUB-CHANNEL", config.AppConfig.GetString("X_DATAHUB_CHANNEL"))
	apiReq.SetBasicAuth(config.AppConfig.GetString("BRIGATE_BASIC_USER"), config.AppConfig.GetString("BRIGATE_BASIC_PASS"))

	resp, err := api.client.Do(apiReq)
	if err != nil {
		errCh <- errs.NewUnexpectedError("failed to do request GetPostalCode: " + err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errCh <- errs.NewUnexpectedError("failed to ReadAll GetPostalCode: " + err.Error())
		return
	}

	var data domain.BrigateResponse[domain.PostalCode]

	if err := json.Unmarshal(body, &data); err != nil {
		errCh <- errs.NewUnexpectedError("failed to unmarshall response GetPostalCode: " + err.Error())
		return
	}

	var postalCode domain.PostalCode
	pcStr := strconv.Itoa(pc)

	for _, resp := range data.Data {
		if pcStr == resp.PostalCode {
			postalCode.PostalCode = resp.PostalCode
			postalCode.AreaCode = resp.AreaCode
			postalCode.Province = resp.Province
			postalCode.City = resp.City
			postalCode.SubDistrict = resp.SubDistrict
			postalCode.Village = resp.Village
		}
	}

	cpc <- postalCode
}
