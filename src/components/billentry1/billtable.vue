<template>
  <v-container class="pt-0 mt-0">

    <v-layout>
      <v-flex lg12 class="d-flex justify-end">
        <snackbar :obj="obj" :rounded-pill="true" />
      </v-flex>
    </v-layout>
    <v-card class="grey lighten-3 pa-2">
      <v-layout>
        <v-flex md-2 lg-3>
          <preview :Previewarr="userArr" :Total="tp" :GST="gst" :NetPrice="netpayable" class="mt-2" />
        </v-flex>
        <v-flex md-2 lg-3>
          <v-btn class="mr-6 mb-5 grey darken-1 white--text mt-2" @click="savebtn">SAVE</v-btn>
        </v-flex>
        <v-flex md-4 class="mt-3" lg-3>
          <span class="mt-2 mr-6 "><b>BILL NO : </b>{{ billNum }}</span>
          <span class="mt-2 mr-6"><b>DATE : </b>{{ date }}</span>
          <span class="mt-2 mr-6"><b>TOTAL :</b>{{ tp }}</span>
         </v-flex>
         <v-flex md-4 class="mt-3" lg-3>
          <span class="mt-2 mr-6"><b>GST (18%) :</b>{{ ' ' + (Math.trunc((tp * (18 / 100)) * 100)) / 100 + '' }}</span>
          <span class="mt-2 mr-6"><b>NetPayable : </b>{{ this.tp + ((Math.trunc((tp * (18 / 100)) * 100)) / 100) }}</span>
         </v-flex>
      </v-layout>
    </v-card>
    <v-card elevation="2" class="grey lighten-3">
      <div class="d-flex justify-end pa-2"> <v-btn dark color=" grey darken-1 white--text" @click="download"
          dense> <v-icon></v-icon></v-btn></div>
      <v-card-title>
        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
          dense></v-text-field>
      </v-card-title>
      <v-data-table :headers="headers" :items="userArr" :search="search" dense rounded
        class="grey lighten-3"></v-data-table>
    </v-card>


  </v-container>
</template>
  
  
<script>
import Papa from "papaparse";
import EventService from '../../Services/EventService';
import preview from '../billentry1/preview.vue';
import snackbar from '../snackbar.vue';


export default {
  name: "billtable",
  data() {
    return {
      search: '',
      headers: [
        { text: 'Medicine Name', value: 'medicine_name', class: "blue darken-1 white--text rounded-l-lg" },
        { text: 'Brand', value: 'brand', class: "blue  darken-1 white--text" },
        { text: 'Quantity', value: 'quantity', class: "blue darken-1 white--text" },
        { text: 'Amount', value: 'amount', class: "blue darken-1 white--text rounded-r-lg" },
      ],
      SaveBillDetails: {
        user_id: "",
        bill_no: this.billNum,
        bill_date: "",
        login_id: 0,
        billsarr: [],
      },
      SaveBillResp: "",
      date: new Date().toJSON().slice(0, 10),
      total: 0,
      gst: 0,
      netpayable: 0,
      UserId: '',
      obj: {
        snackbar: false,
        color: 'red',
        topvar: false,
        bottomvar: false,
        text: '',
      },


    }
  },
  components: {
    preview,
    snackbar,
  },
  props: {
    billNum: Number,
    userArr: [],
    tp: Number,

  },

  methods: {

    savebtn() {


      if (this.userArr.length == 0) {
        this.obj = {
          snackbar: true,
          color: 'red',
          topvar: true,
          bottomvar: false,
          text: 'Bill is Empty '.toUpperCase(),
        }

      } else {
        this.SaveBillDetails.bill_no = this.billNum
        this.SaveBillDetails.bill_date = this.date
        this.SaveBillDetails.billsarr = this.userArr,
          this.SaveBillDetails.login_id = this.$store.state.UserDetails.Login_id
        this.SaveBillDetails.user_id = this.$store.state.UserDetails.User_id

        EventService.SaveBill(this.SaveBillDetails)
          .then((response) => {
            console.log(response);
            this.SaveBillResp = response.data
            this.$emit('billno', response.data.bill_no)
          })
          .catch((error) => {
            console.log(error);
          })
        this.$emit('EmptyuserArr')
      }
    },
    download() {
      if (this.userArr.length != 0) {

        const csv = Papa.unparse(this.userArr);
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
    }


    // getGst() {
    //   // this.gst = Math.floor(this.tp * (18 / 100));
    //   this.gst = Math.trunc((this.tp * (18 / 100)) * 100) / 100
    // },
    // getNetPay() {
    //   this.netpayable = this.tp + Math.floor(this.tp * (18 / 100));
    // },
    // billdataEmpty() {
    //     this.gst = 0,
    //     this.netpayable = 0,
    //     this.$emit('emptypreview');
    // }
  },
  watch: {
    tp() {
      // this.gst = Math.floor(this.tp * (18 / 100));
      this.gst = (Math.trunc((this.tp * (18 / 100)) * 100)) / 100

      // this.netpayable = this.tp + Math.floor(this.tp * (18 / 100));

      this.netpayable = this.tp + ((Math.trunc((this.tp * (18 / 100)) * 100)) / 100)
    },


  }

}

</script>
  