package transform

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/cuducos/go-cnpj"
)

var companyNameClenupRegex = regexp.MustCompile(`(\D)(\d{3})(\d{5})(\d{3})$`) // masks CPF in MEI names

func companyNameClenup(n string) string {
	return strings.TrimSpace(companyNameClenupRegex.ReplaceAllString(n, "$1***$3***"))
}

type Company struct {
	CNPJ                             string        `json:"cnpj" bson:"cnpj"`
	IdentificadorMatrizFilial        *int          `json:"identificador_matriz_filial" bson:"identificador_matriz_filial"`
	DescricaoMatrizFilial            *string       `json:"descricao_identificador_matriz_filial" bson:"descricao_identificador_matriz_filial"`
	NomeFantasia                     string        `json:"nome_fantasia" bson:"nome_fantasia"`
	SituacaoCadastral                *int          `json:"situacao_cadastral" bson:"situacao_cadastral"`
	DescricaoSituacaoCadastral       *string       `json:"descricao_situacao_cadastral" bson:"descricao_situacao_cadastral"`
	DataSituacaoCadastral            *date         `json:"data_situacao_cadastral" bson:"data_situacao_cadastral"`
	MotivoSituacaoCadastral          *int          `json:"motivo_situacao_cadastral" bson:"motivo_situacao_cadastral"`
	DescricaoMotivoSituacaoCadastral *string       `json:"descricao_motivo_situacao_cadastral" bson:"descricao_motivo_situacao_cadastral"`
	NomeCidadeNoExterior             string        `json:"nome_cidade_no_exterior" bson:"nome_cidade_no_exterior"`
	CodigoPais                       *int          `json:"codigo_pais" bson:"codigo_pais"`
	Pais                             *string       `json:"pais" bson:"pais"`
	DataInicioAtividade              *date         `json:"data_inicio_atividade" bson:"data_inicio_atividade"`
	CNAEFiscal                       *int          `json:"cnae_fiscal" bson:"cnae_fiscal"`
	CNAEFiscalDescricao              *string       `json:"cnae_fiscal_descricao" bson:"cnae_fiscal_descricao"`
	DescricaoTipoDeLogradouro        string        `json:"descricao_tipo_de_logradouro" bson:"descricao_tipo_de_logradouro"`
	Logradouro                       string        `json:"logradouro" bson:"logradouro"`
	Numero                           string        `json:"numero" bson:"numero"`
	Complemento                      string        `json:"complemento" bson:"complemento"`
	Bairro                           string        `json:"bairro" bson:"bairro"`
	CEP                              string        `json:"cep" bson:"cep"`
	UF                               string        `json:"uf" bson:"uf"`
	CodigoMunicipio                  *int          `json:"codigo_municipio" bson:"codigo_municipio"`
	CodigoMunicipioIBGE              *int          `json:"codigo_municipio_ibge" bson:"codigo_municipio_ibge"`
	Municipio                        *string       `json:"municipio" bson:"municipio"`
	Telefone1                        string        `json:"ddd_telefone_1" bson:"ddd_telefone_1"`
	Telefone2                        string        `json:"ddd_telefone_2" bson:"ddd_telefone_2"`
	Fax                              string        `json:"ddd_fax" bson:"ddd_fax"`
	Email                            *string       `json:"email" bson:"email"`
	SituacaoEspecial                 string        `json:"situacao_especial" bson:"situacao_especial"`
	DataSituacaoEspecial             *date         `json:"data_situacao_especial" bson:"data_situacao_especial"`
	OpcaoPeloSimples                 *bool         `json:"opcao_pelo_simples" bson:"opcao_pelo_simples"`
	DataOpcaoPeloSimples             *date         `json:"data_opcao_pelo_simples" bson:"data_opcao_pelo_simples"`
	DataExclusaoDoSimples            *date         `json:"data_exclusao_do_simples" bson:"data_exclusao_do_simples"`
	OpcaoPeloMEI                     *bool         `json:"opcao_pelo_mei" bson:"opcao_pelo_mei"`
	DataOpcaoPeloMEI                 *date         `json:"data_opcao_pelo_mei" bson:"data_opcao_pelo_mei"`
	DataExclusaoDoMEI                *date         `json:"data_exclusao_do_mei" bson:"data_exclusao_do_mei"`
	RazaoSocial                      string        `json:"razao_social" bson:"razao_social"`
	CodigoNaturezaJuridica           *int          `json:"codigo_natureza_juridica" bson:"codigo_natureza_juridica"`
	NaturezaJuridica                 *string       `json:"natureza_juridica" bson:"natureza_juridica"`
	QualificacaoDoResponsavel        *int          `json:"qualificacao_do_responsavel" bson:"qualificacao_do_responsavel"`
	CapitalSocial                    *float32      `json:"capital_social" bson:"capital_social"`
	CodigoPorte                      *int          `json:"codigo_porte" bson:"codigo_porte"`
	Porte                            *string       `json:"porte" bson:"porte"`
	EnteFederativoResponsavel        string        `json:"ente_federativo_responsavel" bson:"ente_federativo_responsavel"`
	QuadroSocietario                 []PartnerData `json:"qsa" bson:"qsa"`
	CNAESecundarios                  []CNAE        `json:"cnaes_secundarios" bson:"cnaes_secundarios"`
	RegimeTributario                 TaxRegimes    `json:"regime_tributario" bson:"regime_tributario"`
}

func (c *Company) situacaoCadastral(v string) error {
	i, err := toInt(v)
	if err != nil {
		return fmt.Errorf("error trying to parse SituacaoCadastral %s: %w", v, err)
	}

	var s string
	switch *i {
	case 1:
		s = "NULA"
	case 2:
		s = "ATIVA"
	case 3:
		s = "SUSPENSA"
	case 4:
		s = "INAPTA"
	case 8:
		s = "BAIXADA"
	}

	c.SituacaoCadastral = i
	if s != "" {
		c.DescricaoSituacaoCadastral = &s
	}
	return nil
}

func (c *Company) identificadorMatrizFilial(v string) error {
	i, err := toInt(v)
	if err != nil {
		return fmt.Errorf("error trying to parse IdentificadorMatrizFilial %s: %w", v, err)
	}

	var s string
	switch *i {
	case 1:
		s = "MATRIZ"
	case 2:
		s = "FILIAL"
	}

	c.IdentificadorMatrizFilial = i
	if s != "" {
		c.DescricaoMatrizFilial = &s
	}
	return nil
}

func newCompany(row []string, l *lookups, kv kvStorage, privacy bool) (Company, error) {
	var c Company
	if len(row) != 30 {
		return c, fmt.Errorf("invalid row with %d columns (expected 30): %v", len(row), row)
	}
	c.CNPJ = row[0] + row[1] + row[2]
	c.NomeFantasia = row[4]
	c.NomeCidadeNoExterior = row[8]
	c.DescricaoTipoDeLogradouro = row[13]
	c.Logradouro = row[14]
	c.Numero = row[15]
	c.Complemento = row[16]
	c.Bairro = row[17]
	c.CEP = row[18]
	c.UF = row[19]
	c.Telefone1 = row[21] + row[22]
	c.Telefone2 = row[23] + row[24]
	c.Fax = row[25] + row[26]
	c.Email = &row[27]
	c.SituacaoEspecial = row[28]

	if privacy {
		c.NomeFantasia = companyNameClenup(row[4])
		c.Email = nil
		if c.CodigoNaturezaJuridica != nil && strings.Contains(strings.ToLower(*c.NaturezaJuridica), "individual") {
			c.DescricaoTipoDeLogradouro = ""
			c.Logradouro = ""
			c.Numero = ""
			c.Complemento = ""
			c.Telefone1 = ""
			c.Telefone2 = ""
			c.Fax = ""
		}
	}

	if err := c.identificadorMatrizFilial(row[3]); err != nil {
		return c, fmt.Errorf("error trying to parse IdentificadorMatrizFilial: %w", err)
	}

	if err := c.situacaoCadastral(row[5]); err != nil {
		return c, fmt.Errorf("error trying to parse SituacaoCadastral: %w", err)
	}

	dataSituacaoCadastral, err := toDate(row[6])
	if err != nil {
		return c, fmt.Errorf("error trying to parse DataSituacaoCadastral %s: %w", row[3], err)
	}
	c.DataSituacaoCadastral = dataSituacaoCadastral

	if err := c.motivoSituacaoCadastral(l, row[7]); err != nil {
		return c, fmt.Errorf("error trying to parse MotivoSituacaoCadastral: %w", err)
	}

	if err := c.pais(l, row[9]); err != nil {
		return c, fmt.Errorf("error trying to parse CodigoPais: %w", err)
	}

	dataInicioAtividade, err := toDate(row[10])
	if err != nil {
		return c, fmt.Errorf("error trying to parse DataInicioAtividade %s: %w", row[10], err)
	}
	c.DataInicioAtividade = dataInicioAtividade

	if err := c.cnaes(l, row[11], row[12]); err != nil {
		return c, fmt.Errorf("error trying to parse cnae: %w", err)
	}

	if err := c.municipio(l, row[20]); err != nil {
		return c, fmt.Errorf("error trying to parse CodigoMunicipio %s: %w", row[20], err)
	}

	dataSituacaoEspecial, err := toDate(row[29])
	if err != nil {
		return c, fmt.Errorf("error trying to parse DataSituacaoEspecial %s: %w", row[20], err)
	}
	c.DataSituacaoEspecial = dataSituacaoEspecial

	if err := kv.enrichCompany(&c); err != nil {
		return c, fmt.Errorf("error enriching company %s: %w", cnpj.Mask(c.CNPJ), err)
	}
	return c, nil
}

func (c *Company) JSON() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("error while mashaling company JSON: %w", err)
	}
	return string(b), nil
}

func jsonFields(i any) []string {
	var fs []string
	t := reflect.TypeOf(i)
	for i := range t.NumField() {
		f := t.Field(i)
		fs = append(fs, f.Tag.Get("json"))
	}
	return fs
}

// JSONFields lists the field names/paths for the JSON of a company.
func CompanyJSONFields() []string {
	c := jsonFields(Company{})
	t := reflect.TypeOf(Company{})
	var fs []string
	for i := range c {
		f := t.Field(i)
		t := f.Tag.Get("json")
		switch t {
		case "qsa":
			for _, n := range jsonFields(PartnerData{}) {
				fs = append(fs, fmt.Sprintf("%s.%s", t, n))
			}
		case "cnaes_secundarios":
			for _, n := range jsonFields(CNAE{}) {
				fs = append(fs, fmt.Sprintf("%s.%s", t, n))
			}
		case "regime_tributario":
			for _, n := range jsonFields(TaxRegime{}) {
				fs = append(fs, fmt.Sprintf("%s.%s", t, n))
			}
		default:
			fs = append(fs, t)
		}
	}
	return fs
}
