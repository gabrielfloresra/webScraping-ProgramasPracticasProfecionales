<template>
    <div>
        <v-card>
            <v-card-title>
                <v-text-field v-model="search" label="Search"></v-text-field>
            </v-card-title>
            <v-data-table :headers="headers" :items="items" loading="true" :search="search">
                <template v-slot:body="{ items }">
                    <tbody>
                        <tr
                            v-for="item in items"
                            :key="item.nombrePrograma"
                            @click="filaSeleccionada(item)"
                        >
                            <td>{{ item.NombrePrograma }}</td>
                            <td>{{ item.Area }}</td>
                            <td>{{ item.Empresa }}</td>
                            <td v-if="item.ApoyoEconomico != ''">{{ item.ApoyoEconomico }}</td>
                            <td v-else>{{item.OtrosApoyos}} (OTROS)</td>
                            <td>{{ item.ActividadesRealizar }}</td>
                        </tr>
                    </tbody>
                </template>
            </v-data-table>
        </v-card>
        <v-row justify="center">
            <v-dialog v-model="dialog">
                <v-card>
                    <v-card-title>
                        <span>Mas Informacion...</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col>
                                    <v-card>
                                        <v-card-title></v-card-title>
                                        <v-card-text>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Apoyo Economico"
                                                        v-if="dataDialog.ApoyoEconomico != ''"
                                                        v-model="dataDialog.ApoyoEconomico"
                                                    ></v-text-field>
                                                    <v-text-field
                                                        label="Apoyo Economico"
                                                        v-else
                                                        value="Sin Apoyo Económico"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Otro Apoyos"
                                                        v-if="dataDialog.OtrosApoyos"
                                                        v-model="dataDialog.OtrosApoyos"
                                                    ></v-text-field>
                                                    <v-text-field
                                                        label="Otro Apoyos"
                                                        v-else
                                                        value="ninguno"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="12">
                                                    <v-textarea
                                                        label="Actividades a Realizar"
                                                        v-model="dataDialog.ActividadesRealizar"
                                                    ></v-textarea>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col>
                                    <v-card>
                                        <v-card-title></v-card-title>
                                        <v-card-text>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Coordinador que inscribe"
                                                        v-model="dataDialog.CoordinadorQueInscribe"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Curso de Induccion"
                                                        v-model="dataDialog.CursoInduccion"
                                                    ></v-text-field>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col>
                                    <v-card>
                                        <v-card-title>Info Nombramiento</v-card-title>
                                        <v-card-text>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Dirigido a:"
                                                        v-model="dataDialog.NombramientoNombre"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Cargo"
                                                        v-model="dataDialog.NombramientoCargo"
                                                    ></v-text-field>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col>
                                    <v-card>
                                        <v-card-title>Info Empresa</v-card-title>
                                        <v-card-text>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Nombre"
                                                        v-model="dataDialog.Empresa"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Domicilio"
                                                        v-model="dataDialog.Domicilio"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Sector"
                                                        v-model="dataDialog.SectorEmpresa"
                                                    ></v-text-field>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col>
                                    <v-card>
                                        <v-card-title>Fechas Importantes</v-card-title>
                                        <v-card-text>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Registro del Programa"
                                                        v-model="dataDialog.FechaRegistro"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Inicio Programa"
                                                        v-model="dataDialog.FechaInicio"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Termino Programa"
                                                        v-model="dataDialog.FechaTermino"
                                                    ></v-text-field>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col>
                                    <v-card>
                                        <v-card-title>Info Asesor Servicio</v-card-title>
                                        <v-card-text>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Nombre"
                                                        v-model="dataDialog.Responsable"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Cargo"
                                                        v-model="dataDialog.ResponsableCargo"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Correo"
                                                        v-model="dataDialog.ResponsableCorreo"
                                                    ></v-text-field>
                                                </v-col>
                                                <v-col cols="12" sm="12" md="6">
                                                    <v-text-field
                                                        label="Telefono"
                                                        v-model="dataDialog.ResponsableTelefono"
                                                    ></v-text-field>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-col>
                            </v-row>
                        </v-container>
                        <small>
                            informacion oficial:
                            <a
                                @click="infoProgramaOficial"
                            >{{dataDialog.NombrePrograma}}</a>
                        </small>
                    </v-card-text>
                    <v-card-actions>
                        <div class="flex-grow-1"></div>
                        <v-btn color="blue darken-1" text @click="dialog = false">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </v-row>
    </div>
</template>
<script>
export default {
    data: function() {
        return {
            dialog: false,
            search: "",
            dataDialog: {},
            headers: [
                {
                    text: "Nombre Programa",
                    align: "left",
                    value: "NombrePrograma"
                },
                {
                    text: "Área",
                    value: "Area"
                },
                {
                    text: "Empresa",
                    value: "Empresa"
                },
                {
                    text: "Apoyo Economico",
                    value: "ApoyoEconomico"
                },
                {
                    text: "Actividades a Realizar",
                    value: "ActividadesRealizar"
                }
            ],
            items: []
        };
    },
    methods: {
        filaSeleccionada: function(item) {
            this.dialog = true;
            this.dataDialog = item;
        },
        infoProgramaOficial: function() {
            window.open(this.dataDialog.LinkMoreInfo, "_blank");
        },
        strip_html_tags: function(str) {
            if (str === null || str === "" || str === undefined) return "";
            else str = str.innerText.toString();
            return str.replace(/<[^>]*>/g, "");
        },
        init: function() {
            if (!window.WebSocket) {
                alert(
                    "no es posible iniciar el programa desde tu navegador actual!!!"
                );
            } else {
                window.WebSocket = window.WebSocket || window.MozWebSocket;

                var connection = new WebSocket(

                    // cambiar por la ip del servidpr actual
                    "ws://172.31.184.192:8888/scrapData"
                );

                connection.onopen = function() {
                    connection.send("datos requeridos");
                };

                connection.onerror = function(error) {
                    // an error occurred when sending/receiving data
                    console.log(error);
                };

                connection.onmessage = message => {
                    let data = JSON.parse(message["data"]);
                    if (Object.keys(data).length > 0) {
                        connection.close();
                        Object.keys(data).forEach(key => {
                            let newObject = data[key];
                            newObject["Folio"] = key;
                            this.items.push(newObject);
                        });
                        console.log(this.items);
                    } else {
                        connection.send("datos requeridos");
                    }
                };
            }
        }
    },
    mounted: function() {
        this.init();
    }
};
</script>