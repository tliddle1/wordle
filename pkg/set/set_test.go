package set

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSetFixture(t *testing.T) {
	gunit.Run(new(SetFixture), t)
}

type SetFixture struct {
	*gunit.Fixture
	set Set[int]
}

func (this *SetFixture) Setup() {
	this.set = Set[int]{}
}

func (this *SetFixture) TestNoOp() {
	this.So(this.set, should.HaveLength, 0)
}

func (this *SetFixture) TestAdd() {
	this.set.Add(1)
	this.So(this.set, should.HaveLength, 1)
	this.So(this.set.Contains(1), should.BeTrue)
}

func (this *SetFixture) TestRemove() {
	this.set.Add(1)
	this.set.Remove(1)
	this.So(this.set, should.HaveLength, 0)
	this.So(this.set.Contains(1), should.BeFalse)
}

func (this *SetFixture) TestContains() {
	this.set.Add(1)
	this.So(this.set.Contains(1), should.BeTrue)
	this.So(this.set.Contains(0), should.BeFalse)
}
