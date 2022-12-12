package persist

import (
	"crawler/model"
	"crawler/persist"
)

type ItemSaveServer struct {
	Client *persist.ElasticStruct
	Index  string
}

func (s *ItemSaveServer) Save(item model.User, result *string) error {
	*result = s.Client.Save(s.Index, item.Id, item)
	if *result != "" {
		*result = "ok"
	}
	return nil
}
