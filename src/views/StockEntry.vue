<template>
  <div class="StockEntry">
    <v-layout class="mt-16"></v-layout>
    <v-layout class="mt-16">
      <v-container class="pa-6 mt-16 grey lighten-3">
        <v-layout>
          <v-flex xs12 md12>
            <h3 class="font-weight-medium  blue--text font-weight-bold text-uppercase">REFILL STOCK</h3>
          </v-flex>
          <v-flex>
            <popup @lastitem="lastmedreceive" />
          </v-flex>
        </v-layout>
        <v-layout class="mt-6">
          <v-flex class="mr-3">
            <v-autocomplete :items="item" label="Medicine Name" dense rounded outlined color="grey"
              v-model="StockDetails.medicine_name" @change="med"></v-autocomplete>

          </v-flex>
          <v-flex class="mr-3">
            <v-text-field label="Brand" disabled dense rounded outlined color="grey"
              v-model="StockDetails.brand"></v-text-field>

          </v-flex>
          <v-flex class="mr-3">
            <v-text-field v-model.number="StockDetails.quantity" label="quantity" type="Number" rounded dense outlined
              color="grey"></v-text-field>
          </v-flex>
          <v-flex class="mr-3">
            <v-text-field v-model.number="StockDetails.unit_price" label="Unit Price" type="Number" rounded dense outlined
              color="grey"></v-text-field>
          </v-flex>
        </v-layout>
        <v-layout>
          <v-flex>
            <v-btn block dark color="blue lighten-6" @click="update">UPDATE</v-btn>
          </v-flex>
        </v-layout>
        <snackbar :obj="obj" />
      </v-container>

    </v-layout>
    <br><br>
    <!-- {{ Resp }} -->
    <!-- {{ this.$store.state.UserDetails.User_id }} -->

  </div>
</template>
  
<script>
import EventService from '../Services/EventService';
import snackbar from '../components/snackbar.vue';
import popup from '../components/popup.vue';

export default {
  name: "StockEntry",
  beforeMount() {
    EventService.StockEntryGetStocks()
      .then((response) => {
        console.log(response);
        if (response.data.status == "S") {

          this.Resp = response.data.stockarr;
          for (let i = 0; i < this.Resp.length; i++) {
            this.item.push(this.Resp[i].medicine_name);
          }

        } else if (response.data.status == "E") {
          console.log(response.data.msg);
        }
      })
      .catch((error) => {
        console.log(error);
      })



  },
  data() {
    return {
      Resp: [],
      StockDetails: {
        medicine_name: "",
        brand: "",
        quantity: 0,
        unit_price: 0,
        user_id: ""
      },
      Medicine: [],
      snackval: false,
      value1: true,
      item: [],
      lastmed: "",

      obj: {
        snackbar: false,
        color: 'green',

        topvar: false,
        bottomvar: false,
        text: ' ',
      },
    }
  },

  components: {
    popup,
    snackbar,
  },
  methods: {
    med() {




      for (let j = 0; j < this.Resp.length; j++) {
        if (this.Resp[j].medicine_name == this.StockDetails.medicine_name) {
          this.StockDetails.brand = this.Resp[j].brand;
        }
      }
      // for (let i = 0; i < this.MedicineMaster.length; i++) {
      //   if (this.StockDetails.medicine_name == this.MedicineMaster[i].StockDetails.medicine_name) {
      //     this.StockDetails.brand = this.MedicineMaster[i].Brand;
      //   }
      // }
    },
    update() {
      if (this.StockDetails.medicine_name != '' && this.StockDetails.brand != '' && this.StockDetails.quantity > 0 && this.StockDetails.unit_price > 0 && this.StockDetails.quantity != '' && this.StockDetails.unit_price != '') {
        // for (let i = 0; i < this.stockarr.length; i++) {
        //   if (this.StockDetails.medicine_name == this.stockarr[i].StockDetails.medicine_name) {
        //     // this.snackval=true;
        //     this.stockarr[i].quantity += this.StockDetails.quantity;
        //     this.stockarr[i].StockDetails.unit_price = this.StockDetails.unit_price;
        //     this.value1 = false;
        //   }
        // }

        this.StockDetails.user_id=this.$store.state.UserDetails.User_id
        console.log(this.StockDetails);
        EventService.UpdateStock(this.StockDetails)
          .then((response) => {
            console.log(response);

            if (response.data.status == "S") {
              this.obj = {
                snackbar: true,
                color: 'green',
                topvar: true,
                bottomvar: false,
                text: 'Updated Successfully',
              };

            } else if (response.data.status == "E") {
              console.log(response.data.msg);

            }
          })
          .catch((error) => {
            console.log(error);
          })


        if (this.value1 == true) {

          this.StockDetails.medicine_name = '';
          this.StockDetails.quantity = 0;
          this.StockDetails.unit_price = 0;
          this.StockDetails.brand = '';

        } else {
          this.obj = {
            snackbar: true,
            color: 'green',
            topvar: true,
            bottomvar: false,
            text: 'Updated Successfully',
          };
          // alert(`Updated Successfully` );
          this.StockDetails.medicine_name = '';
          this.StockDetails.quantity = 0;
          this.StockDetails.unit_price = 0;
          this.StockDetails.brand = '';
        }
      } else {
        // alert('Enter the proper data');
        this.obj = {
          snackbar: true,
          color: 'red',
          topvar: true,
          bottomvar: false,
          text: 'Enter the valid Inputs',
        };
      }
    },
    lastmedreceive(med) {
      this.lastmed = med.medicine_name
      this.Resp.push(med);
      this.AddLastMed()


    },
    AddLastMed() {
      this.item.push(this.lastmed)
    }
  }


}

</script>




         