package api

type CertificateIssue struct {
	ObjectMeta
	TypeMeta

	Spec   CertificateIssueSpec   `json:"spec,omitempty"`
	Status CertificateIssueStatus `json:"status,omitempty"`
}

type CertificateIssueList struct {
	ListMeta
	TypeMeta

	Items []CertificateIssue `json:"items,omitempty"`
}

type CertificateIssueSpec struct {
	CommonName          string `json:"commonName,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	KeySize             int    `json:"keySize,omitempty"`
	SignatureAlgorithm  string `json:"signatureAlgorithm,omitempty"`
}

type CertificateIssueStatus struct {
	Phase CertificateIssuePhase `json:"phase,omitempty"`
}

type CertificateIssuePhase string

const (
	CertificateIssueActive CertificateIssuePhase = "Active"
)
