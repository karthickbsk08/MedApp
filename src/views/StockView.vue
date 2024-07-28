<template>
  <v-container class="pa-5 mt-16">
    <v-card class=" grey lighten-3">
      <div class="d-flex justify-space-between pa-4"><h1 class="font-weight-medium  blue--text font-weight-bold text-uppercase pl-5 pb-0 ">STOCK VIEW 
      </h1><v-btn dark color=" grey darken-1 white--text" class="d-inline-block" @click="download"
          dense>DOWNLOAD</v-btn></div>
      <v-col class="d-flex justify-end"></v-col>
      <v-card-title class="pt-0">

        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
          class="pa-3"></v-text-field>
      </v-card-title>
      <v-data-table :headers="headers" :items="desserts" :search="search"
        class=" grey lighten-3 font-weight-medium pa-3"></v-data-table>
    </v-card>
  </v-container>
</template>
  
<script>
import Papa from "papaparse";
import EventService from '../Services/EventService';
export default {
  name: "Stockview",

  beforeMount() {
    EventService.StockView()
      .then((response) => {
        console.log(response);
        this.desserts = response.data.stockviewarr
      })
      .catch((error) => {
        console.log(error);
      })
  },

  data() {
    return {
      search: '',
      headers: [
        { text: 'Medicine Name', value: 'medicine_name', class: "blue lighten-1 rounded-l-lg white--text text-subtitle-1 font-weight-bold " },
        { text: 'Brand', value: 'brand', class: "blue lighten-1  white--text text-subtitle-1 font-weight-bold " },
        { text: 'Quantity', value: 'quantity', class: "blue lighten-1  white--text text-subtitle-1 font-weight-bold " },
        { text: 'Unit price', value: 'unit_price', class: "blue lighten-1 rounded-r-lg white--text text-subtitle-1 font-weight-bold " },
      ],
      desserts: [],
    }
  },
  methods: {
    download() {

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

    }
  },



};
</script>
  