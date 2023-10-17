package builtintemplate

import (
	"context"
	"testing"

	"github.com/anyproto/any-sync/app"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore/mock_objectstore"
	"github.com/anyproto/anytype-heart/tests/testutil"
	"github.com/anyproto/anytype-heart/util/testMock/mockSource"
)

func Test_registerBuiltin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sourceService := mockSource.NewMockService(ctrl)
	sourceService.EXPECT().NewStaticSource(gomock.Any(), gomock.Any(), gomock.Any(), nil).AnyTimes()
	sourceService.EXPECT().RegisterStaticSource(gomock.Any()).AnyTimes()

	objectStore := mock_objectstore.NewMockObjectStore(t)

	builtInTemplates := New()

	ctx := context.Background()
	a := new(app.App)
	a.Register(testutil.PrepareMock(ctx, a, sourceService))
	a.Register(builtInTemplates)
	a.Register(config.New())
	a.Register(testutil.PrepareMock(ctx, a, objectStore))

	err := builtInTemplates.Init(a)
	assert.NoError(t, err)
	err = builtInTemplates.Run(context.Background())
	assert.NoError(t, err)
}
