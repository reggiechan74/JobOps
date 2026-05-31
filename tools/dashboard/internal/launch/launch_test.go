package launch

import (
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func TestCompose(t *testing.T) {
	rec := model.Record{Slug: "Acme_Product_Manager_20260531", Company: "Acme"}
	cases := []struct {
		spec model.SkillSpec
		want string
	}{
		{model.SkillSpec{Name: "assessjob", Arg: model.ArgJD},
			"/jobops:assessjob Acme_Product_Manager_20260531.md --app=Acme_Product_Manager_20260531"},
		{model.SkillSpec{Name: "osint", Arg: model.ArgCompany}, "/jobops:osint Acme"},
		{model.SkillSpec{Name: "convert-to-pdf", Arg: model.ArgNone}, "/jobops:convert-to-pdf"},
	}
	for _, c := range cases {
		if got := Compose(c.spec, rec); got != c.want {
			t.Errorf("Compose(%s) = %q, want %q", c.spec.Name, got, c.want)
		}
	}
}

func TestComposeCompanyUnknown(t *testing.T) {
	rec := model.Record{Slug: "X_Y_20260101"} // no Company
	got := Compose(model.SkillSpec{Name: "osint", Arg: model.ArgCompany}, rec)
	if got != "/jobops:osint" {
		t.Errorf("Compose osint w/o company = %q, want /jobops:osint", got)
	}
}

func TestComposePlugin(t *testing.T) {
	rec := model.Record{Slug: "X_Y_20260101"}
	got := Compose(model.SkillSpec{Name: "ratecard", Arg: model.ArgNone, Plugin: "jobops-ic"}, rec)
	if got != "/jobops-ic:ratecard" {
		t.Errorf("Compose with plugin = %q, want /jobops-ic:ratecard", got)
	}
	got2 := Compose(model.SkillSpec{Name: "idealjob", Arg: model.ArgNone}, rec)
	if got2 != "/jobops:idealjob" {
		t.Errorf("Compose default plugin = %q, want /jobops:idealjob", got2)
	}
}
