package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	router := newRouter()
	fmt.Println("servidor a la espera de peticiones...")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func scraperData(info map[string]dataStruct) {
	// urlProgramas := "http://webserver.siiaa.siu.buap.mx/serv_social/ss_oferta_ss?se=ABE63459925DAD50"
	// urlProgramasPracticasProfesionalesFCC := "http://webserver.siiaa.siu.buap.mx/serv_social/Ss_OFERTA_prog?se=CF593C89C3F382C4&un=695238F6612D0665"
	urlProgramasPracticasProfesionalesFCC := "http://webserver.siiaa.siu.buap.mx/serv_social/SS_ALUM_AUTSV_VAL?PIDM=8689CF68DE8A0DEC3B43BE3456A93921"
	// urlProgramasPracticasProfesionalesFCC := "http://localhost/x/temp/practicasProfesionales/temp/ppp/temp.html"

	c := colly.NewCollector(
		// colly.AllowedDomains("https://www.clotheswebsite.com/"),
		colly.CacheDir(".programasPracticasProfesionalesFCC_cache"),
		// colly.MaxDepth(5), // keeping crawling limited for our initial experiments
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 50})
	space := regexp.MustCompile(`\s+`)

	c.OnHTML(`table`, func(e *colly.HTMLElement) {

		e.ForEach("tr", func(_ int, tr *colly.HTMLElement) {

			folio := tr.ChildText("td:nth-child(1)")
			folio = space.ReplaceAllString(folio, " ")

			if _, err := strconv.Atoi(folio); err == nil {

				// quitar espacios innecesarios en el nombre del programa
				nombrePrograma := tr.ChildText("td:nth-child(2)")
				nombrePrograma = space.ReplaceAllString(nombrePrograma, " ")

				area := tr.ChildText("td:nth-child(3)")
				area = space.ReplaceAllString(area, " ")

				link := "http://webserver.siiaa.siu.buap.mx/serv_social/" + tr.ChildAttr("td>a", "href")

				coordinadorInscribe := tr.ChildText("td:nth-child(4)")
				coordinadorInscribe = space.ReplaceAllString(coordinadorInscribe, " ")

				cursoInduccion := tr.ChildText("td:nth-child(5)")
				cursoInduccion = space.ReplaceAllString(cursoInduccion, " ")

				empresa := tr.ChildText("td:nth-child(6)")
				empresa = space.ReplaceAllString(empresa, " ")

				domicilio := tr.ChildText("td:nth-child(7)")
				domicilio = space.ReplaceAllString(domicilio, " ")

				infoTemp, ok := info[folio]
				if !ok {

					info[folio] = dataStruct{
						nombrePrograma,
						area,
						link,
						coordinadorInscribe,
						cursoInduccion,
						empresa,
						domicilio,
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
						"",
					}
				} else {
					info[folio] = dataStruct{
						nombrePrograma,
						area,
						link,
						coordinadorInscribe,
						cursoInduccion,
						empresa,
						domicilio,
						infoTemp.FechaRegistro,
						infoTemp.SectorEmpresa,
						infoTemp.FechaInicio,
						infoTemp.FechaTermino,
						infoTemp.Responsable,
						infoTemp.ResponsableCargo,
						infoTemp.ResponsableCorreo,
						infoTemp.ResponsableTelefono,
						infoTemp.DireccionEntrevista,
						infoTemp.DireccionFisica,
						infoTemp.NombramientoNombre,
						infoTemp.NombramientoCargo,
						infoTemp.ApoyoEconomico,
						infoTemp.OtrosApoyos,
						infoTemp.ActividadesRealizar,
					}
				}
				c.Visit(link)
			}
		})
	})

	c.OnHTML("div#Layer0>table[class=tableBUAPGrisAlternaRows]", func(e *colly.HTMLElement) {

		folio := e.ChildText("tr:nth-child(2)>td:nth-child(2)")
		folio = space.ReplaceAllString(folio, " ")

		fechaRegistro := e.ChildText("tr:nth-child(3)>td:nth-child(2)")
		fechaRegistro = space.ReplaceAllString(fechaRegistro, " ")

		sectorEmpresa := e.ChildText("tr:nth-child(5)>td:nth-child(2)")
		sectorEmpresa = space.ReplaceAllString(sectorEmpresa, " ")

		fechaInicio := e.ChildText("tr:nth-child(15)>td:nth-child(2)")
		fechaInicio = space.ReplaceAllString(fechaInicio, " ")

		fechaTermino := e.ChildText("tr:nth-child(16)>td:nth-child(2)")
		fechaTermino = space.ReplaceAllString(fechaTermino, " ")

		responsable := e.ChildText("tr:nth-child(21)>td:nth-child(2)")
		responsable = space.ReplaceAllString(responsable, " ")

		responsableCargo := e.ChildText("tr:nth-child(22)>td:nth-child(2)")
		responsableCargo = space.ReplaceAllString(responsableCargo, " ")

		responsableCorreo := e.ChildText("tr:nth-child(23)>td:nth-child(2)")
		responsableCorreo = space.ReplaceAllString(responsableCorreo, " ")

		responsableTelefono := e.ChildText("tr:nth-child(24)>td:nth-child(2)")
		responsableTelefono = space.ReplaceAllString(responsableTelefono, " ")

		direccionEntrevista := e.ChildText("tr:nth-child(25)>td:nth-child(2)")
		direccionEntrevista = space.ReplaceAllString(direccionEntrevista, " ")

		direccionFisica := e.ChildText("tr:nth-child(26)>td:nth-child(2)")
		direccionFisica = space.ReplaceAllString(direccionFisica, " ")

		nombramientoNombre := e.ChildText("tr:nth-child(30)>td:nth-child(2)")
		nombramientoNombre = space.ReplaceAllString(nombramientoNombre, " ")

		nombramientoCargo := e.ChildText("tr:nth-child(31)>td:nth-child(2)")
		nombramientoCargo = space.ReplaceAllString(nombramientoCargo, " ")

		apoyoEconomico := e.ChildText("tr:nth-child(33)>td:nth-child(2)")
		apoyoEconomico = space.ReplaceAllString(apoyoEconomico, " ")

		otrosApoyos := e.ChildText("tr:nth-child(34)>td:nth-child(2)")
		otrosApoyos = space.ReplaceAllString(otrosApoyos, " ")

		actividadesRealizar := e.ChildText("tr:nth-child(39)>td:nth-child(2)")
		actividadesRealizar = space.ReplaceAllString(actividadesRealizar, " ")

		infoTemp, ok := info[folio]
		if !ok {
			info[folio] = dataStruct{
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				fechaRegistro,
				sectorEmpresa,
				fechaInicio,
				fechaTermino,
				responsable,
				responsableCargo,
				responsableCorreo,
				responsableTelefono,
				direccionEntrevista,
				direccionFisica,
				nombramientoNombre,
				nombramientoCargo,
				apoyoEconomico,
				otrosApoyos,
				actividadesRealizar,
			}
		} else {
			info[folio] = dataStruct{
				infoTemp.NombrePrograma,
				infoTemp.Area,
				infoTemp.LinkMoreInfo,
				infoTemp.CoordinadorQueInscribe,
				infoTemp.CursoInduccion,
				infoTemp.Empresa,
				infoTemp.Domicilio,
				fechaRegistro,
				sectorEmpresa,
				fechaInicio,
				fechaTermino,
				responsable,
				responsableCargo,
				responsableCorreo,
				responsableTelefono,
				direccionEntrevista,
				direccionFisica,
				nombramientoNombre,
				nombramientoCargo,
				apoyoEconomico,
				otrosApoyos,
				actividadesRealizar,
			}
		}
	})

	// printing visiting message for debug purposes
	c.OnRequest(func(r *colly.Request) {
		log.Println("--VisitingProgramInfo:", r.URL.String())
	})

	c.Visit(urlProgramasPracticasProfesionalesFCC)
	c.Wait()
	// fmt.Println(info)
}

/*
wg.Add(1)
				go func(wg *sync.WaitGroup, infoPrograma map[string]dataStruct, infoTemp map[string]string) {

					defer wg.Done()

					moreInfoProgramHTML := colly.NewCollector(
						colly.CacheDir(".programasPracticasProfesionalesFCC_cache/" + folio),
					)
					moreInfoProgramHTML.OnHTML(`table[class=tableBUAPGrisAlternaRows]`, func(element *colly.HTMLElement) {

						fmt.Println(element.Text)
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fmt.Println("--------------------------------------------------------")
						fechaRegistro := element.ChildText("tr:nth-child(2)")
						fechaRegistro = space.ReplaceAllString(fechaRegistro, " ")

						sectorEmpresa := element.ChildText("tr:nth-child(5)")
						sectorEmpresa = space.ReplaceAllString(sectorEmpresa, " ")

						fechaInicio := element.ChildText("tr:nth-child(13)")
						fechaInicio = space.ReplaceAllString(fechaInicio, " ")

						fechaTermino := element.ChildText("tr:nth-child(14)")
						fechaTermino = space.ReplaceAllString(fechaTermino, " ")

						responsable := element.ChildText("tr:nth-child(17)")
						responsable = space.ReplaceAllString(responsable, " ")

						responsableCargo := element.ChildText("tr:nth-child(18)")
						responsableCargo = space.ReplaceAllString(responsableCargo, " ")

						responsableCorreo := element.ChildText("tr:nth-child(19)")
						responsableCorreo = space.ReplaceAllString(responsableCorreo, " ")

						responsableTelefono := element.ChildText("tr:nth-child(20)")
						responsableTelefono = space.ReplaceAllString(responsableTelefono, " ")

						direccionEntrevista := element.ChildText("tr:nth-child(21)")
						direccionEntrevista = space.ReplaceAllString(direccionEntrevista, " ")

						direccionFisica := element.ChildText("tr:nth-child(22)")
						direccionFisica = space.ReplaceAllString(direccionFisica, " ")

						nombramientoNombre := element.ChildText("tr:nth-child(24)")
						nombramientoNombre = space.ReplaceAllString(nombramientoNombre, " ")

						nombramientoCargo := element.ChildText("tr:nth-child(25)")
						nombramientoCargo = space.ReplaceAllString(nombramientoCargo, " ")

						apoyoEconomico := element.ChildText("tr:nth-child(33)")
						apoyoEconomico = space.ReplaceAllString(apoyoEconomico, " ")

						otrosApoyos := element.ChildText("tr:nth-child(28)")
						otrosApoyos = space.ReplaceAllString(otrosApoyos, " ")

						actividadesRealizar := element.ChildText("tr:nth-child(39)")
						actividadesRealizar = space.ReplaceAllString(actividadesRealizar, " ")

						// fmt.Println("\n\nfolio:" + infoTemp["folio"])
						// fmt.Println("actividades a Realizar:" + actividadesRealizar)
						mutex.Lock()
						infoPrograma[infoTemp["folio"]] = dataStruct{
							infoTemp["nombrePrograma"],
							infoTemp["tipo"],
							infoTemp["area"],
							infoTemp["carrera"],
							infoTemp["link"],
							infoTemp["coordinadorInscribe"],
							infoTemp["cursoInduccion"],
							infoTemp["empresa"],
							infoTemp["domicilio"],
							fechaRegistro,
							sectorEmpresa,
							fechaInicio,
							fechaTermino,
							responsable,
							responsableCargo,
							responsableCorreo,
							responsableTelefono,
							direccionEntrevista,
							direccionFisica,
							nombramientoNombre,
							nombramientoCargo,
							apoyoEconomico,
							otrosApoyos,
							actividadesRealizar,
						}
						mutex.Unlock()
					})

					moreInfoProgramHTML.OnRequest(func(r *colly.Request) {
						log.Println("----VisitingProgramMoreInfo:", r.URL.String())
					})

					moreInfoProgramHTML.Visit(link)

				}(&wg, info, infoTemp)
*/
