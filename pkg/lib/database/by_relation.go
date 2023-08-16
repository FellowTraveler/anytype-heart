package schema

import (
	"fmt"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
	"github.com/samber/lo"
)

type schemaByRelations struct {
	CommonRelations []*model.RelationLink
}

func NewByRelations(commonRelations []*model.RelationLink) Schema {
	return &schemaByRelations{CommonRelations: commonRelations}
}

func (sch *schemaByRelations) RequiredRelations() []*model.RelationLink {
	alwaysRequired := []*model.RelationLink{
		bundle.MustGetRelationLink(bundle.RelationKeyName),
		bundle.MustGetRelationLink(bundle.RelationKeyType),
	}
	required := append(alwaysRequired, sch.CommonRelations...)
	return lo.UniqBy(required, func(rel *model.RelationLink) string {
		return rel.Key
	})
}

func (sch *schemaByRelations) ListRelations() []*model.RelationLink {
	return sch.CommonRelations
}

func (sch *schemaByRelations) String() string {
	return fmt.Sprintf("relations: %v", pbtypes.GetRelationListKeys(sch.CommonRelations))
}
