package main

type dataStruct struct {
	NombrePrograma         string `json:"NombrePrograma"`
	Area                   string `json:"Area"`
	LinkMoreInfo           string `json:"LinkMoreInfo"`
	CoordinadorQueInscribe string `json:"CoordinadorQueInscribe"`
	CursoInduccion         string `json:"CursoInduccion"`
	Empresa                string `json:"Empresa"`
	Domicilio              string `json:"Domicilio"`
	FechaRegistro          string `json:"FechaRegistro"`
	SectorEmpresa          string `json:"SectorEmpresa"`
	FechaInicio            string `json:"FechaInicio"`
	FechaTermino           string `json:"FechaTermino"`
	Responsable            string `json:"Responsable"`
	ResponsableCargo       string `json:"ResponsableCargo"`
	ResponsableCorreo      string `json:"ResponsableCorreo"`
	ResponsableTelefono    string `json:"ResponsableTelefono"`
	DireccionEntrevista    string `json:"DireccionEntrevista"`
	DireccionFisica        string `json:"DireccionFisica"`
	NombramientoNombre     string `json:"NombramientoNombre"`
	NombramientoCargo      string `json:"NombramientoCargo"`
	ApoyoEconomico         string `json:"ApoyoEconomico"`
	OtrosApoyos            string `json:"OtrosApoyos"`
	ActividadesRealizar    string `json:"ActividadesRealizar"`
}
