package blocklist_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/jonathanfoster/freedom/model/blocklist"
)

func TestBlocklist(t *testing.T) {
	Convey("Blocklist", t, func() {
		Convey("All", func() {
			Convey("Should return blocklists", func() {
				list, err := blocklist.All()
				if err != nil {
					assert.Error(t, err)
				}

				assert.NotEmpty(t, list)
			})
		})
	})
}
