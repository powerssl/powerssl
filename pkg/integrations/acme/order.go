package acme

/*
import (
	integrationacme "powerssl.io/pkg/integration/acme"
)

func (i *ACME) XCreateOrder(account integrationacme.Account, identifiers []integrationacme.Identifier, notBefore, notAfter string) (*integrationacme.Order, error) {
	newOrderReq := struct {
		Identifiers []Identifier `json:"identifiers"`
		NotBefore   string       `json:"notBefore,omitempty"`
		NotAfter    string       `json:"notAfter,omitempty"`
	}{
		Identifiers: identifiers,
		NotBefore:   notBefore,
		NotAfter:    notAfter,
	}
	newOrderResp := integrationacme.Order{}
	resp, err := c.post(c.dir.NewOrder, account.URL, account.PrivateKey, newOrderReq, &newOrderResp, http.StatusCreated)
	if err != nil {
		return newOrderResp, err
	}

	newOrderResp.URL = resp.Header.Get("Location")

	return newOrderResp, nil
}
*/
