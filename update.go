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
var OKDESK_ADD_COMMENT = "%s/api/v1/issues/%v/comments"

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

	res = model.Issue{}
	return res, json.NewDecoder(body).Decode(&res)
}

/**

	https://okdesk.ru/apidoc#!kommentarii-dobavlenie-kommentariya

content				string	обязательный	Текст комментария. Текст комментария является html-текстом, а значит требуемое форматирование необходимо задавать с использованием html-тегов. Например, для переноса строк необходимо использовать тег </br>
author_id			integer	обязательный	ID пользователя, являющегося автором комментария
author_type			string	опционально	Тип пользователя, являющегося автором комменатрия
	Допустимые типы: 	employee (по умолчанию) - сотрудник, contact - контактное лицо
public				boolean	опционально	Флаг публичности комментария. Принимает значение true или false
attachments			array	опционально	Список приложенных файлов

*/

func AddCommentRaw(ctx context.Context, id string, req io.ReadCloser, header map[string]string) (res io.ReadCloser, err error) {
	Url := fmt.Sprintf(OKDESK_ADD_COMMENT, os.Getenv("OKDESK_URL"), id)
	return Request(ctx, "POST", Url, req, header)
}

type CommentRequest struct {
	Content    string  `json:"content"`
	AuthorID   *int    `json:"author_id,omitempty"`
	AuthorType *string `json:"author_type,omitempty"`
	Public     bool    `json:"public,omitempty"`
	//todo: files
	//attachments *

}

type CommentResponse struct {
	ID      int          `json:"id"`
	Content string       `json:"content"`
	Public  bool         `json:"public"`
	Author  model.User   `json:"author"`
	Errors  *model.Error `json:"errors,omitempty"`
}

func AddComment(ctx context.Context, id string, req CommentRequest, header map[string]string) (res CommentResponse, err error) {

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

	res = CommentResponse{}
	return res, json.NewDecoder(body).Decode(&res)

}
