<template>
  <v-container class="mt-16">
    <v-layout></v-layout>
    <v-layout class="mt-5">
      <v-flex>
        <billitem @EachNewMed="NewMed" @changerough="roughemit1" :roughdata="rough" :RespStock="Resp"  :itemArr="data"/>
        <billtable :userArr="Medarr" :tp="totalprice" :Previewarr="Medarr"
          @emptypreview="emptyit" @EmptyuserArr="EmptyArr" :billNum="bill_No" @billno="AssignBillNo" />
      </v-flex>
    </v-layout>


  </v-container>
</template>
  
<script>

import EventService from '../Services/EventService';
import billitem from '../components/billentry1/billitem.vue';
import billtable from '../components/billentry1/billtable.vue';



export default {
  name: "BillEntry",
  components: {
    billitem, billtable,
  },

  data() {
    return {
      bill_No: 0,
      Medarr: [],
      totalprice: 0,
      rough: "",      
      Resp: [],
      data:[],
      
      
  }
  },
  methods: {
    EmptyArr(){   

      this.Medarr = [];
      this.totalprice = 0;
      this.data=[]   
    },
    NewMed(arr,totalprice) {
      this.Medarr = arr;  
      this.totalprice=totalprice   

    },
    emptyit() {
      this.Medarr = [];
      this.totalprice = 0;
      this.rough = "true";
    },
    roughemit1() {
      this.rough = "";
    },
    AssignBillNo(new_bill_no){
      this.bill_No=new_bill_no
    },
    GetBillStock() {
      EventService.BillStock()
        .then((response) => {
          console.log(response);

          if (response.data.status == "S") {
            this.Resp = response.data.billstockarr
            this.bill_No=response.data.bill_no
            // this.$store.commit('setBillDetails',this.Resp)
            // console.log("BillEntry");
            // console.log(this.Resp);
          } else if (response.data.status == "E") {
            console.log(response.data.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        })
    }
  },
  mounted(){
    this.GetBillStock();
  }


};
</script>
  