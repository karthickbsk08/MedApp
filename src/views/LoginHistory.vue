<template>
  <v-container class="mt-10">

    <v-card class="pa-2 blue lighten-5 mt-16">
      <div class="d-flex justify-space-between pa-2">
        <h1 class="font-weight-medium blue--text font-weight-bold text-uppercase">Login History</h1>
        <v-btn dark color=" grey darken-1 white--text" class="d-inline-block" @click="download" dense>DOWNLOAD</v-btn>
      </div>
      <v-card-title>
        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
      </v-card-title>
      <v-data-table :headers="headers" :items="desserts" :search="search"
        class="blue lighten-5 font-weight-medium"></v-data-table>
    </v-card>
    <!-- {{ this.$store.state.LoginHistory }} -->
    <!-- {{ desserts }} -->
  </v-container>
</template>
  
<script>
import Papa from "papaparse";
import EventService from '../Services/EventService.js';

export default {
  name: "loginhistory",
  data() {
    return {
      search: '',
      headers: [

        { text: 'UserId', value: 'createdBy', class: "blue lighten-1 rounded-l-lg white--text text-subtitle-1 font-weight-bold ", align: "left" },
        { text: 'Login (Date)', value: 'loginDate', class: "blue lighten-1 white--text text-subtitle-1 font-weight-bold" },
        { text: 'Login (time)', value: 'loginTime', class: "blue lighten-1 white--text text-subtitle-1 font-weight-bold" },
        { text: 'Log_out (Date)', value: 'logoutDate', class: "blue lighten-1 white--text text-subtitle-1 font-weight-bold" },
        { text: 'Log_out (time)', value: 'logoutTime', class: "blue lighten-1 rounded-r-lg white--text text-subtitle-1 font-weight-bold" },
      ],

      desserts: [],
    }
  },
  beforeMount() {
    EventService.LoginHistory()
      .then((response) => {
        console.log(response);
        if (response.data.status=="S"){
          this.desserts = response.data.login_historyarr

        }else if (response.data.status =="E"){
          console.log(response.data.msg);
        }
       
      }).catch((error) => {
        console.log(error);
      })
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
  }
};
</script>
  