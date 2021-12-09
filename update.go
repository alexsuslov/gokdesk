package gokdesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/gokdesk/model"
	"io"
	"io/ioutil"
	"os"
)

var OKDESK_SET_STATUS = "%s/api/v1/issues/%s/statuses"

/**
 * https://okdesk.ru/apidoc#!status-zayavki-smena-statusa-zayavki

code 				string	обязательный	Код статуса
delay_to 			string	опционально	Время, до которого откладывается заявка. Обязательное поле только для статуса delayed. Для остальных статусов данный параметр будет игнорироваться.
comment				string	опционально	Комментарий к статусу. Например, причина перехода в новый статус. Обязательность данного параметра настраивается в разделе “Настройки/Заявки/Статусы заявок”.
comment_public		boolean	опционально	Флаг публичности комментария. Принимает значение true или false
custom_parameters	associative array	опционально	Дополнительные атрибуты заявки, которые доступны пользователю для заполнения при входе в новый или при выходе из текущего статуса.
time_entry	array	опционально	Массив трудозатрат по заявке, которые будут добавлены при смене статуса

*/

func SetStatusRAW(ctx context.Context, id string, req io.ReadCloser, header map[string]string) (res io.ReadCloser, err error) {

	Url := fmt.Sprintf(OKDESK_SET_STATUS, os.Getenv("OKDESK_URL"), id)

	return Request(ctx, "POST", Url, req, header)
}

type StatusRequest struct {
	Code             string            `json:"code"`
	DelayTo          *string           `json:"delay_to,omitempty"`
	Comment          *string           `json:"comment,omitempty"`
	CommentPublic    *bool             `json:"comment_public,omitempty"`
	CustomParameters []model.Parameter `json:"custom_parameters,omitempty"`
	TimeEntry        []model.TimeEntry `json:"time_entry,omitempty"`
}

// SetStatus Set Status
func SetStatus(ctx context.Context, id string, req StatusRequest, header map[string]string) (res model.Issue, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	r := ioutil.NopCloser(bytes.NewReader(data))

	body, err := SetStatusRAW(ctx, id, r, header)
	if err != nil {
		return
	}

	defer body.Close()

	issue := model.Issue{}
	return issue, json.NewDecoder(body).Decode(&issue)
}
