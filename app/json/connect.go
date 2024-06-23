package json

import (
	"errors"
	"ssh+/app/file"
	"ssh+/app/output"
)

type Connections struct {
	Connects []Connect `json:"connects"`
}

type Connect struct {
	Alias    string `json:"alias"`
	Login    string `json:"login"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (c *Connections) GetConnectionsAlias() []string {
	filePath := GetPathToConnectFile()

	data, err := file.ReadFile(filePath)
	if err != nil {
		output.GetOutError("Ошибка открытия файла")
	}

	c.SerializationJson(data)
	c.SetDecryptData()

	var result []string

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	if len(result) == 0 {
		output.GetOutError("Подключений не найдено")
	}

	return result
}

func (c *Connections) ExistConnectJsonByIndex(alias string) (int, error) {
	var noFound = -1

	filePath := GetPathToConnectFile()

	data, err := file.ReadFile(filePath)
	if err != nil {
		output.GetOutError("Ошибка открытия файла")
	}

	c.SerializationJson(data)
	c.SetDecryptData()

	defer c.SetCryptAllData()

	for i, v := range c.Connects {
		if v.Alias == alias {
			return i, nil
		}
	}

	return noFound, errors.New("not found")
}

func (c *Connections) WriteConnectToJson(connect Connect) {
	_, err := c.ExistConnectJsonByIndex(connect.Alias)
	if err == nil {
		output.GetOutError("Алиас должен быть уникальным")
	}

	filePath := GetPathToConnectFile()

	data, err := file.ReadFile(filePath)
	if err != nil {
		output.GetOutError("Ошибка открытия файла")
	}

	c.SerializationJson(data)

	encodedConnect := SetCryptData(connect)
	c.Connects = append(c.Connects, encodedConnect)

	file.WriteFile(filePath, c.deserializationJson())
}

func (c *Connections) updateJsonDataByIndex(index int, connect Connect) error {
	if index >= 0 && index < len(c.Connects) {
		c.Connects[index].Alias = connect.Alias
		c.Connects[index].Address = connect.Address
		c.Connects[index].Login = connect.Login
		c.Connects[index].Password = connect.Password
		return nil
	}

	return errors.New("Ошибка обновления подключения")
}

func (c *Connections) UpdateConnectJson(alias string, connect Connect) {
	index, err := c.ExistConnectJsonByIndex(alias)
	if err != nil {
		output.GetOutError("Не найдено подключение")
	}

	err = c.updateJsonDataByIndex(index, SetCryptData(connect))
	if err != nil {
		output.GetOutError("Ошибка обновления подключения")
	}

	file.WriteFile(GetPathToConnectFile(), c.deserializationJson())
}

func (c *Connections) deleteJsonDataByIndex(index int) {
	copy(c.Connects[index:], c.Connects[index+1:])

	c.Connects[len(c.Connects)-1] = Connect{}
	c.Connects = c.Connects[:len(c.Connects)-1]
}

func (c *Connections) DeleteConnectToJson(alias string) {
	index, err := c.ExistConnectJsonByIndex(alias)
	if err != nil {
		output.GetOutError("Не найдено подключение")
	}

	c.deleteJsonDataByIndex(index)

	file.WriteFile(GetPathToConnectFile(), c.deserializationJson())
}
