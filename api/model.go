package api

import "time"

type Domain struct {
	ObjectClassName string `json:"objectClassName"`
	Handle          string `json:"handle"`
	LdhName         string `json:"ldhName"`
	Links           []struct {
		Value string `json:"value"`
		Rel   string `json:"rel"`
		Href  string `json:"href"`
		Type  string `json:"type"`
	} `json:"links"`
	Status   []string `json:"status"`
	Entities []struct {
		ObjectClassName string   `json:"objectClassName"`
		Handle          string   `json:"handle"`
		Roles           []string `json:"roles"`
		PublicIds       []struct {
			Type       string `json:"type"`
			Identifier string `json:"identifier"`
		} `json:"publicIds"`
		VcardArray []any `json:"vcardArray"`
		Entities   []struct {
			ObjectClassName string   `json:"objectClassName"`
			Roles           []string `json:"roles"`
			VcardArray      []any    `json:"vcardArray"`
		} `json:"entities"`
	} `json:"entities"`
	Events []struct {
		EventAction string    `json:"eventAction"`
		EventDate   time.Time `json:"eventDate"`
	} `json:"events"`
	SecureDNS struct {
		DelegationSigned bool `json:"delegationSigned"`
	} `json:"secureDNS"`
	Nameservers []struct {
		ObjectClassName string `json:"objectClassName"`
		LdhName         string `json:"ldhName"`
	} `json:"nameservers"`
	RdapConformance []string `json:"rdapConformance"`
	Notices         []struct {
		Title       string   `json:"title"`
		Description []string `json:"description"`
		Links       []struct {
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"links"`
	} `json:"notices"`
}

type UrlDomain struct {
	DomainParam string `json:"domainParameter"`
}
type Response struct {
	EventDate   time.Time `json:"eventDate"`
	FinalResult int       `json:"difference"`
}
