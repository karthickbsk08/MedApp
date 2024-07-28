<template>
    <div class="text-center">
        <v-dialog v-model="dialog" width="500px">
            <template v-slot:activator="{ on, attrs }">
                <v-btn class="grey darken-1 white--text" v-bind="attrs" v-on="on" block>
                    <v-icon left>mdi-plus</v-icon>
                    <span text color="black" dense>ADD</span>
                </v-btn>
            </template>

            <v-card v-show="dialog">
                <v-card-title class="text-h6 grey lighten-2 text-center">
                    <span>ADD STOCK</span>
                    <v-spacer></v-spacer>
                    <v-icon @click="dialog = false">mdi-close</v-icon>

                </v-card-title>

                <v-card-text>
                    <v-layout class="mt-4">
                        <v-flex md-12>
                            <v-text-field v-model="NewMedicine.medicine_name" solo clearable label="Medicine Name"
                                class="black--text darken-6"></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout>
                        <v-flex md-12>
                            <v-text-field v-model="NewMedicine.brand" solo clearable label="brand"
                                class="black--text darken-6"></v-text-field>
                        </v-flex>
                    </v-layout>
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn dark color="blue darken-4" block @click="add">
                        ADD
                    </v-btn>
                </v-card-actions>
                <!-- {{ laststock }} -->
                <!-- {{ this.$store.state.UserDetails }} -->
            </v-card>
        </v-dialog>
        <snackbar :obj="obj" />

    </div>
</template>
<script>
import EventService from '../Services/EventService';
import snackbar from '../components/snackbar.vue';
export default {
    name: '',
    components: {
        snackbar,
    },
    data() {
        return {
            dialog: false,
            NewMedicine: {
                medicine_name: "",
                brand: "",
                user_id: "",
            },
            value1: true,

            laststock: '',
            obj: {
                snackbar: false,
                color: 'green',
                topvar: false,
                bottomvar: false,
                text: ' ',
            },


        }
    },
    methods: {
        add() {
            if (this.NewMedicine.medicine_name != '' && this.NewMedicine.brand != '') {
                this.NewMedicine.user_id = this.$store.state.UserDetails.User_id
                EventService.AddMedicine(this.NewMedicine)
                    .then((response) => {
                        console.log(response);

                        if (response.data.status == "S") {
                            this.laststock=response.data.last_medicine
                            this.$emit('lastitem', this.laststock);
                            this.obj = {
                                snackbar: true,
                                color: 'green',
                                topvar: true,
                                bottomvar: false,
                                text: `${response.data.msg}`,
                            };
                            this.NewMedicine.medicine_name = '';
                            this.NewMedicine.brand = '';


                        } else if (response.data.status == "E") {
                            this.obj = {
                                snackbar: true,
                                color: 'red',
                                topvar: true,
                                bottomvar: false,
                                text: `${response.data.msg}`,
                            };
                            this.NewMedicine.medicine_name = '';
                            this.NewMedicine.brand = '';
                        }
                    })
                    .catch((error) => {
                        console.log(error);
                    })
                 
                 this.NewMedicine.medicine_name = '';
                 this.NewMedicine.brand = '';
                 

            } else {

                this.obj = {
                    snackbar: true,
                    color: 'red',
                    topvar: true,
                    bottomvar: false,
                    text: 'Enter the valid Inputs',
                };
            }
        }
    }
}
</script>