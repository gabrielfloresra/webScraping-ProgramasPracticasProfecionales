<template>
    <v-app>
        <v-app-bar app>
            <v-toolbar-title class="headline text-uppercase">
                <span>Programas</span>
                <v-divider class="mx-5" vertical></v-divider>
                <span class="font-weight-light">Practicas Profesionales</span>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn text href="#" target="_blank">
                <span class="mr-2">Sitio Oficial</span>
            </v-btn>
        </v-app-bar>
        <v-content>
            <muestraProgramas />
        </v-content>
    </v-app>
</template>
<script>
    import muestraProgramas from './components/muestraProgramas';

    export default {
        name: 'App',
        components: {
            muestraProgramas
        },
        data: () => ({
            //
        })
    };
</script>