<template>
  <v-container class="mt-16">

    <h1 class="blue--text darken-6--text font-weight-black text-center mt-15 display-1 text-capitalize">Your today sales
    </h1>

    <h1 class="black--text lighten-3--text font-weight-black text-center mt-4 display-3">{{ todaytotal }}
      <span v-show="showdata">
        <v-icon size="50" color="green">mdi-arrow-up</v-icon>
        <span class="title green--text">{{ '(' + (Math.floor(tdypercentage)) + '%)' }}</span>
      </span>
      <span v-show="showdatas">
        <v-icon size="50" color="red">mdi-arrow-down</v-icon>
        <span class="title red--text">{{ '(' + (Math.floor(yespercentage)) + '%)' }}</span>
      </span>


    </h1>
  </v-container>
</template>
  
<script>
import EventService from '../../Services/EventService';

export default {
  name: "Billerdash",
  data() {
    return {
      showdata: false,
      showdatas: false,
      tdypercentage: 0,
      yespercentage: 0,
      currentdate: '',
      previousdate: '',
      BillMaster: this.$store.state.BillMaster,
      LoginHistory: this.$store.state.LoginHistory,
      todaytotal: 0,
      yesterdaytotal: 0,
      userId: '',
      // tdypercentage:0,
      // yespercentage:0,

    }
  },

  methods: {
    todaydate() {
      this.currentdate = new Date().toJSON().slice(0, 10);
    },
    yesterdayDate() {
      this.previousdate = new Date(new Date().setDate(new Date().getDate() - 1)).toJSON().slice(0, 10);
    },
    todaysales() {

      for (let j = 0; j < this.LoginHistory.length; j++) {
        if (this.LoginHistory[j].Log_out == 'In_Working') {
          this.userId = this.LoginHistory[j].UserId;
          // console.log(this.userId);
        }

      }
      for (let i = 0; i < this.BillMaster.length; i++) {
        if (this.currentdate == this.BillMaster[i].BillDate) {
          if (this.userId == this.BillMaster[i].UserId) {
            this.todaytotal += this.BillMaster[i].NetPrice;
            // console.log(this.todaytotal);
          }


        }
      }
    },
    previoussales() {
      for (let i = 0; i < this.BillMaster.length; i++) {
        if (this.previousdate == this.BillMaster[i].BillDate) {
          if (this.userId == this.BillMaster[i].UserId) {
            this.yesterdaytotal += this.BillMaster[i].NetPrice;
          }
        }
      }
    }
  },
  beforeMount() {
    var Login_id = this.$store.state.UserDetails.Login_id

    EventService.BillerDash(Login_id)
      .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          this.todaytotal = response.data.todaysales
          this.yesterdaytotal = response.data.yesterdaysales

          if (this.todaytotal > this.yesterdaytotal) {
            this.tdypercentage = (((this.todaytotal - this.yesterdaytotal) / this.todaytotal) * 100);
            console.log("PRofit");
            this.showdata = true;
            

          } else {
            this.yespercentage = (((this.yesterdaytotal - this.todaytotal) / this.yesterdaytotal) * 100);
            console.log("Loss");
            this.showdatas = true;
          }
        } else if (response.data.status == "E") {
          console.log(response.data.msg);
        }

      })
      .catch((error) => {
        console.log(error);
      })


    // this.todaydate();
    // this.yesterdayDate();
    // this.todaysales();
    // this.previoussales();
    // if (this.todaytotal > this.yesterdaytotal) {
    //   this.tdypercentage = (((this.todaytotal - this.yesterdaytotal) / this.todaytotal) * 100);
    //   this.showdata = true;

    // } else {

    //   this.yespercentage = (((this.yesterdaytotal - this.todaytotal) / this.yesterdaytotal) * 100);
    //   this.showdatas = true;
    // }


  },
  computed: {
    // CompSales() {
    //   if ((isNaN(this.tdypercentage) || this.tdypercentage == Infinity)||( isNaN(this.yespercentage) || this.yespercentage == Infinity)) {
    //     return ;
    //   } else if(this.tdypercentage>0) {
    //     return this.tdypercentage
    //   }else{
    //     return this.yespercentage
    //   }

    // }
  }

};
// 
</script>
  