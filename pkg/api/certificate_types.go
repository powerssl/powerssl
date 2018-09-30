package api

type Certificate struct {
	ObjectMeta
	TypeMeta

	Spec   CertificateSpec   `json:"spec,omitempty"`
	Status CertificateStatus `json:"status,omitempty"`
}

type CertificateList struct {
	ListMeta
	TypeMeta

	Items []Certificate `json:"items,omitempty"`
}

type CertificateSpec struct {
	CommonName          string `json:"commonName,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	KeySize             int    `json:"keySize,omitempty"`
	SignatureAlgorithm  string `json:"signatureAlgorithm,omitempty"`
	AutoRenew           bool   `json:"autoRenew,omitempty"`
}

type CertificateStatus struct {
	Phase CertificatePhase `json:"phase,omitempty"`
}

type CertificatePhase string

const (
	CertificateActive CertificatePhase = "Active"
)
