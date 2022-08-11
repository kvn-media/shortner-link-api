package database

import (
	"time"

	"github.com/kvn-media/shortner-link-api/types"
	"github.com/kvn-media/shortner-link-api/utils"
)

func CreateNewShortenedLink(url string) *types.ShortenedLink {
	randId := utils.RandStringBytesMaskImprSrcUnsafe(10)
	obj := &types.ShortenedLink{ProvidedLink: url, LinkId: randId, ShortenedAt: time.Now().UnixNano()}
	Session.Create(obj)
	// Session.Commit()
	return obj
}

func GetShortenedLink(id string) *types.ShortenedLink {
	var res *types.ShortenedLink
	obj := &types.ShortenedLink{LinkId: id}
	Session.Take(&res, obj)
	return res
}
