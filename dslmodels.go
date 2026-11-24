package main

type DSLData struct {
    DSLModulationType         string  `json:"DSL Modulation Type"`
    AnnexType                 string  `json:"Annex Type"`
    CurrentRateUpStream       uint16  `json:"Current Rate(kbps) Upstream"`
    CurrentRateDownStream     uint16  `json:"Current Rate(kbps) Downstream"`
    MaxRateUpStream           uint32  `json:"Max Rate(kbps) Upstream"`
    MaxRateDownStream         uint32  `json:"Max Rate(kbps) Downstream"`
    SNRMarginUpStream         float32 `json:"SNR Margin(dB) Upstream"`
    SNRMarginDownStream       float32 `json:"SNR Margin(dB) Downstream"`
    LineAttenuationUpStream   float32 `json:"Line Attenuation(dB) Upstream"`
    LineAttenuationDownStream float32 `json:"Line Attenuation(dB) Downstream"`
    ATUCCRCErrors             int32   `json:"Errors(ATUCCRCE) Upstream"`
    CRCErrors                 int32   `json:"Errors(CRC) Downstream"`
    Status                    string  `json:"Status"`
}

type ClientData struct {
    IPAddress          string `json:"IPAddress"`
    AddressSource      string `json:"AddressSource"`
    LeaseTimeRemaining int32  `json:"LeaseTimeRemaining"`
    MACAddress         string `json:"MACAddress"`
    Hostname           string `json:"Hostname"`
    ConnectionType     string `json:"ConnectionType"`
    Active             bool   `json:"Active"`
}
