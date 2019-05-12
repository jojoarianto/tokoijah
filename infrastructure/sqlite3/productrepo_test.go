package sqlite3

import (
	"testing"
	"time"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/stretchr/testify/assert"
)

var (
	URIDbConn = "../../tokoijah.sqlite3"
	Dialeg    = "sqlite3"
)

func TestAdd(t *testing.T) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	assert.NoError(t, err)

	nowTimestamp := time.Now()
	sku := nowTimestamp.String()

	t.Run("success", func(t *testing.T) {
		product := model.Product{
			Sku:  sku,
			Name: "Zalekia Plain Casual Blouse",
		}

		repo := NewProductRepo(db)
		err = repo.Add(product)
		if err != nil {
			t.Error("error")
		}

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		product := model.Product{
			Sku:  sku,
			Name: "Zalekia Plain Casual Blouse (L,Broken White)",
		}

		repo := NewProductRepo(db)
		err = repo.Add(product)

		assert.Error(t, err)
	})
}
