<template>
  <div>
    <v-layout></v-layout>
    <v-layout class="mt-16">
      <v-container class="mt-16 ">
        <v-card class="pa-6  grey lighten-3 rounded-b-xl" elevation="1">
          <h1 class="font-weight-medium pa-2 blue--text font-weight-bold text-uppercase">SALES REPORT</h1>
          <v-row class="ml-15 mt-5 pb-0">
            <v-col cols="12" sm="4" md="4">
              <v-menu ref="menu" v-model="menu" :close-on-content-click="false" :return-value.sync="DateRange.from_date"
                transition="scale-transition" offset-y min-width="auto">
                <template v-slot:activator="{ on, attrs }">

                  <v-text-field v-model="DateRange.from_date" label="From" append-icon="mdi-calendar" readonly
                    v-bind="attrs" v-on="on" outlined dense class="pb-0"> </v-text-field>
                </template>
                <v-date-picker v-model="DateRange.from_date" no-title scrollable>
                  <v-spacer></v-spacer>
                  <v-btn text color="primary" @click="menu = false">
                    Cancel
                  </v-btn>
                  <v-btn text color="primary" @click="$refs.menu.save(DateRange.from_date)">
                    OK
                  </v-btn>
                </v-date-picker>
              </v-menu>
            </v-col>

            <v-spacer></v-spacer>

            <v-col cols="12" sm="4" md="4" class="pb-0">
              <v-menu ref="modal" v-model="modal" :close-on-content-click="false" :return-value.sync="DateRange.to_date"
                transition="scale-transition" offset-y min-width="auto">
                <template v-slot:activator="{ on, attrs }">

                  <v-text-field v-model="DateRange.to_date" label="To" append-icon="mdi-calendar" readonly v-bind="attrs"
                    v-on="on" outlined dense class="pb-0"></v-text-field>
                </template>
                <v-date-picker v-model="DateRange.to_date" no-title scrollable>
                  <v-spacer></v-spacer>
                  <v-btn text color="primary" @click="modal = false">
                    Cancel
                  </v-btn>
                  <v-btn text color="primary" @click="$refs.modal.save(DateRange.to_date)">
                    OK
                  </v-btn>
                </v-date-picker>
              </v-menu>
            </v-col>

            <v-col md="4" sm="4" class="pb-0">
              <v-btn dark class="  ml-6 grey darken-1 white--text " dense @click="searchbtn">SEARCH</v-btn>
            </v-col>
          </v-row>

          <v-row>
            <v-col class="d-flex justify-end"><v-btn dark color=" grey darken-1 white--text" @click="down"
                dense>DOWNLOAD</v-btn></v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-card elevation="0" class="grey lighten-3 font-weight-medium pa-3">
                <v-card-title>
                  <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                    class="pt-0"></v-text-field>
                </v-card-title>
                <v-data-table :headers="headers" :items="desserts" :search="search" dense
                  class="grey lighten-3  font-weight-medium pa-3"></v-data-table>
              </v-card>
            </v-col>
          </v-row>
        </v-card>
      </v-container>
    </v-layout>

    <snackbar :obj="obj" />
  </div>
</template>
  
<script>
import EventService from '../Services/EventService';
import snackbar from '../components/snackbar.vue';
import Papa from "papaparse";
export default {
  name: "SalesReport",
  components: {
    snackbar,
  },
  data() {
    return {
      DateRange: {
        from_date: "",
        to_date: "",
      },
      menu: false,
      modal: false,

      search: '',
      headers: [
        { text: 'Bill No', value: 'bill_no', class: "blue lighten-1 rounded-l-lg white--text text-body-2 font-weight-medium" },
        { text: 'Bill date', value: 'bill_date', class: "blue lighten-1  white--text text-body-2 font-weight-medium" },
        { text: 'Medicine Name', value: 'medicine_name', class: "blue lighten-1  white--text text-body-2 font-weight-medium" },
        { text: 'Qty', value: 'quantity', class: "blue lighten-1  white--text text-body-2 font-weight-medium" },
        { text: 'Amount', value: 'amount', class: "blue lighten-1 rounded-r-lg white--text text-body-2 font-weight-medium" },

      ],
      desserts: [],

      amount: 0,
      obj: {
        snackbar: false,
        color: 'red',
        topvar: false,
        bottomvar: false,
        text: '',

      },
      data: ""
    }
  },
  methods: {
    down() {
      if (this.desserts.length != 0) {
        const csv = Papa.unparse(this.desserts);
        // console.log("data", csv, typeof (csv), [csv]);
        const blob = new Blob([csv], { type: "text/csv;charset=utf-8;" });
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.setAttribute("href", url);
        link.setAttribute("download", "data.csv");
        link.style.visibility = "visible";
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      } else {
        this.obj = {
          snackbar: true,
          color: 'red',
          topvar: true,
          bottomvar: false,
          text: 'There is No Bill to Download',

        }
      }
    },
    searchbtn() {

      if (this.DateRange.from_date != "" && this.DateRange.to_date != "") {
        if (this.DateRange.from_date > this.DateRange.to_date || this.DateRange.to_date > new Date().toJSON().slice(0, 10)) {

          this.obj = {
            snackbar: true,
            color: 'red',
            topvar: true,
            bottomvar: false,
            text: 'Enter the Valid Date',

          }
          return;

        }

        EventService.SalesReport(this.DateRange)
          .then((response) => {
            console.log(response);
            if (response.data.status == "S") {
              this.desserts = response.data.salesarr

              this.DateRange.from_date = '';
              this.DateRange.to_date = '';
            } else if (response.data.status == "E") {
              this.obj = {
                snackbar: true,
                color: 'red',
                topvar: true,
                bottomvar: false,
                text: `${response.data.msg}`,

              }
              this.DateRange.from_date = '';
            this.DateRange.to_date = '';

            }
           
          })
          .catch((error) => {
            console.log(error);
          })


      } else {
        // alert("Enter the Valid data!")
        this.obj = {
          snackbar: true,
          color: 'red',
          topvar: true,
          bottomvar: false,
          text: 'Enter the Valid Inputs',

        }
        this.DateRange.from_date = '';
        this.DateRange.to_date = '';
      }
    }
  }


};
</script>
  
<!--  -->
<!-- (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)
(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10) -->
 <!-- <v-container>
     
      <v-card class="pa-6" elevation="1">
        <h4 class="pb-6">SALES REPORT</h4>

        <v-layout class="d-flex justify-space-around">
          <v-flex md="4" color="orange lighten-4">

            <v-menu ref="menu1" v-model="menu" :close-on-content-click="false" :return-value.sync="date"
              transition="scale-transition" offset-y min-width="auto">
              <template v-slot:activator="{ on, attrs }">

                <v-text-field v-model="DateRange.from_date" label="From" class="ml-6" outlined append-icon="mdi-calendar" dense
                  v-bind="attrs" v-on="on"></v-text-field>

              </template>
              <v-date-picker v-model="date" no-title scrollable>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="menu = false">
                  Cancel
                </v-btn>
                <v-btn text color="primary" @click="$refs.menu.save(DateRange.from_date)">
                  OK
                </v-btn>
              </v-date-picker>
            </v-menu>

          </v-flex>
          <v-flex md="4">
            <v-menu ref="menu2" v-model="menu" :close-on-content-click="false" :return-value.sync="date"
              transition="scale-transition" offset-y min-width="auto">
              <template v-slot:activator="{ on, attrs }">

                <v-text-field v-model="DateRange.to_date" label="Picker in menu" class="ml-6" outlined append-icon="mdi-calendar" dense
                  v-bind="attrs" v-on="on"></v-text-field>

              </template>
              <v-date-picker v-model="date" no-title scrollable>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="menu = false">
                  Cancel
                </v-btn>
                <v-btn text color="primary" @click="$refs.menu.save(DateRange.to_date)">
                  OK
                </v-btn>
              </v-date-picker>
            </v-menu>
          </v-flex>
          <v-flex md-4>
            <v-btn dark class="grey lighten-1 black--text ml-6" dense>SEARCH</v-btn>
          </v-flex>
        </v-layout>
        <v-card>
          <v-layout >
            <v-flex  class="d-flex justify-end"><v-btn  dark  color="grey darken-6" dense>DOWNLOAD</v-btn></v-flex>
          </v-layout>
        <v-layout>
          
        <v-flex>
          <v-card elevation="0">
            <v-card-title>
              <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                hide-details></v-text-field>
            </v-card-title>
            <v-data-table :headers="headers" :items="desserts" :search="search"></v-data-table>
          </v-card>
        </v-flex>
      </v-layout>

        </v-card>
      </v-card>

     
    </v-container> -->


   