package types

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

// Profile contains all data of a single profile
type Profile struct {
	Name   string
	Status Status
	IPList []string
	Routes map[string]*Route
}

// String returns a string representation of the profile
func (p *Profile) String() string {
	return fmt.Sprintf("[%s]%s", p.Status, p.Name)
}

// GetStatus returns a string value of ProfileStatus
func (p *Profile) GetStatus() string {
	return string(p.Status)
}

func (p *Profile) AppendIP(n string) {
	for _, c := range p.IPList {
		if c == n {
			return
		}
	}

	p.IPList = append(p.IPList, n)
}

// AddRoute adds a single route to the profile
func (p *Profile) AddRoute(ip, hostname string) {
	if p.Routes == nil {
		p.Routes = map[string]*Route{}
	}

	if p.Routes[ip] == nil {
		p.AppendIP(ip)
		p.Routes[ip] = &Route{
			IP:        net.ParseIP(ip),
			HostNames: []string{hostname},
		}
	} else {
		p.Routes[ip].HostNames = append(p.Routes[ip].HostNames, hostname)
	}
}

// AddRoutes adds multiple routes to the profile
func (p *Profile) AddRoutes(ip string, hostnames []string) {
	if p.Routes == nil {
		p.Routes = map[string]*Route{}
	}

	if p.Routes[ip] == nil {
		p.AppendIP(ip)
		p.Routes[ip] = &Route{
			IP:        net.ParseIP(ip),
			HostNames: hostnames,
		}
	} else {
		p.Routes[ip].HostNames = append(p.Routes[ip].HostNames, hostnames...)
	}
}

// RemoveRoutes removes multiple hostnames of a profile
func (p *Profile) RemoveRoutes(hostnames []string) {
	for _, h := range hostnames {
		for ip, r := range p.Routes {
			r.HostNames = remove(r.HostNames, h)
			if len(r.HostNames) == 0 {
				delete(p.Routes, ip)
			}
		}
	}
}

// GetHostNames returns a list of all hostnames of the given ip.
func (p *Profile) GetHostNames(ip string) ([]string, error) {
	key := net.ParseIP(ip)
	if key == nil {
		return nil, fmt.Errorf("%w '%s'", ErrInvalidIP, ip)
	}

	hosts, ok := p.Routes[key.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %s[%s] ", ErrNotPresentIP, key, p.Name)
	}

	return hosts.HostNames, nil
}

// GetAllHostNames returns all hostnames of the profile.
func (p *Profile) GetAllHostNames() ([]string, error) {
	list := []string{}

	for _, r := range p.Routes {
		list = append(list, r.HostNames...)
	}

	return list, nil
}

// Render writes the profile content to the given StringWriter
func (p *Profile) Render(w io.StringWriter) error {
	tmp := bytes.NewBufferString("")

	_, err := tmp.WriteString(fmt.Sprintf("\n# profile.%s %s\n", p.Status, p.Name))
	if err != nil {
		return err
	}

	for _, ip := range p.IPList {
		route := p.Routes[ip]
		for _, host := range route.HostNames {
			prefix := ""
			if p.Status == Disabled {
				prefix = "# "
			}

			_, err = tmp.WriteString(fmt.Sprintf("%s%s %s\n", prefix, ip, host))
			if err != nil {
				return err
			}
		}
	}

	_, err = tmp.WriteString("# end\n")
	if err != nil {
		return err
	}

	// Write to input writer after knowing the profile is well formed
	_, err = w.WriteString(tmp.String())

	return err
}

func remove(s []string, n string) []string {
	list := []string{}

	for _, x := range s {
		if x != n {
			list = append(list, x)
		}
	}

	return list
}
