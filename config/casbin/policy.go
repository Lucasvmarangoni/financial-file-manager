package casbin

import (
	"strconv"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
)

type Policy struct {
	Rules  [][]string
	Groups [][]string
	Admin  []string
	Read   []string
}

func NewPolice() *Policy {
	return &Policy{}
}

func (p *Policy) SetPolicy() {
	p.Rules = [][]string{
		{"anonymous", "/authn/create", "POST"},
		{"anonymous", "/authn/jwt", "POST"},
		{"member", "/user/me", "GET"},
		{"member", "/user/update", "PUT"},
		{"member", "/user/del", "DELETE"},
		{"ADMIN", "*", "*"},
		{"READ", "*", "GET"},
	}

	maxAdminsStr := config.GetEnv("authz_max_admin").(string)
	maxAdmins, err := strconv.Atoi(maxAdminsStr)
	if err != nil {
	}
	maxReadsStr := config.GetEnv("authz_max_read").(string)
	maxReads, err := strconv.Atoi(maxReadsStr)
	if err != nil {
	}

	for i := 0; i < maxAdmins; i++ {
		p.Admin = append(p.Admin, config.GetEnv("admin_"+strconv.Itoa(i +1)).(string))
	}

	for i := 0; i < maxReads; i++ {
		p.Read = append(p.Read, config.GetEnv("read_"+strconv.Itoa(i +1)).(string))
	}

	p.Groups = [][]string{}

	for _, admin := range p.Admin {
		groupAdmin := []string{admin, "ADMIN"}
		p.Groups = append(p.Groups, groupAdmin)
	}

	for _, reader := range p.Read {
		groupRead := []string{reader, "READ"}
		p.Groups = append(p.Groups, groupRead)
	}
}
