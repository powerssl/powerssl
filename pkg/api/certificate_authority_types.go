package api

type CertificateAuthority struct {
	ObjectMeta
	TypeMeta

	Spec CertificateAuthoritySpec `json:"spec,omitempty"`
}

type CertificateAuthorityList struct {
	ListMeta
	TypeMeta

	Items []CertificateAuthority `json:"items,omitempty"`
}

type CertificateAuthoritySpec struct {
	Vendor string `json:"vendor,omitempty"`
}
