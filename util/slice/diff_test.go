package slice

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_Diff(t *testing.T) {
	origin := []string{"000", "001", "002", "003", "004", "005", "006", "007", "008", "009"}
	changed := []string{"000", "008", "001", "002", "003", "005", "006", "007", "009", "004"}

	chs := Diff(origin, changed)

	assert.Equal(t, chs, []Change{
		{Op: OperationRemove, Ids: []string{"004", "008"}},
		{Op: OperationAdd, Ids: []string{"008"}, AfterId: "000"},
		{Op: OperationAdd, Ids: []string{"004"}, AfterId: "009"}},
	)
}

func Test_ChangesApply(t *testing.T) {
	origin := []string{"000", "001", "002", "003", "004", "005", "006", "007", "008", "009"}
	changed := []string{"000", "008", "001", "002", "003", "005", "006", "007", "009", "004", "new"}

	chs := Diff(origin, changed)

	res := ApplyChanges(origin, chs)

	assert.Equal(t, changed, res)
}

func Test_RandomSameLength(t *testing.T) {
	for i := 0; i< 1000; i++ {
		l := randNum(5, 100)
		origin := getRandArray(l)
		changed := make([]string, len(origin))
		copy(changed, origin)
		rand.Shuffle(len(changed),
			func(i, j int) { changed[i], changed[j] = changed[j], changed[i] })

		chs := Diff(origin, changed)
		res := ApplyChanges(origin, chs)

		assert.Equal(t, res, changed)
	}
}

func randNum(min, max int) int{
	rand.Seed(time.Now().UnixNano())
	return  rand.Intn(max - min) + min
}

func getRandArray(len int) []string {
	res := make([]string, len)
	for i := 0; i < len; i++ {
		res[i] = bson.NewObjectId().Hex()
	}
	return res
}