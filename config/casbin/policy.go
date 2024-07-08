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
		{"member", "/totp/generate", "GET"},
		{"member", "/totp/verify/*", "POST"},
		{"member", "/totp/disable", "PATCH"},
		{"member", "/user/me", "GET"},
		{"member", "/user/update", "PUT"},
		{"member", "/user/del", "DELETE"},
		{"ADMIN", "*", "*"},
		{"READ", "/user/*", "GET"},
	}

	maxAdmins := config.ReadSecretInt(config.GetEnvString("authz", "max_admin"))
	maxReads := config.ReadSecretInt(config.GetEnvString("authz", "max_read"))

	for i := 0; i < maxAdmins; i++ {
		p.Admin = append(p.Admin, config.ReadSecretString(config.GetEnvString("authz", "admin_"+strconv.Itoa(i+1))))
	}

	for i := 0; i < maxReads; i++ {
		p.Read = append(p.Read, config.ReadSecretString(config.GetEnvString("authz", "read_"+strconv.Itoa(i+1))))
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
